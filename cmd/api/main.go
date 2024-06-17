package main

import (
	"fmt"
	"log"
	"net/http"
	"product-warehouse/cmd/api/controller"
	"product-warehouse/cmd/api/routes"
	"product-warehouse/config"
	port "product-warehouse/internal/port/repository"
	"product-warehouse/internal/repository/inMemory"
	productUsecase "product-warehouse/internal/usecase/product"
	stockUsecase "product-warehouse/internal/usecase/stock"
)

func main() {
	PORT := config.GetEnv("PORT")
	ENV := config.GetEnv("ENV")

	var productRepository port.ProductRepository
	var stockRepository port.StockRepository

	switch ENV {
	case "prod":
		fmt.Println("Production environment")
	default:
		productRepository = inMemory.NewInMemoryProductRepository()
		stockRepository = inMemory.NewInMemoryStockRepository()
	}
	
	createProductUsecase := productUsecase.NewCreateProductUsecase(productRepository)
	findProductByIdUsecase := productUsecase.NewFindProductByIdUsecase(productRepository)
	productController := controller.NewProductController(createProductUsecase, findProductByIdUsecase)

	createStockUsecase := stockUsecase.NewCreateStockUsecase(stockRepository, productRepository)
	findStockByProductId := stockUsecase.NewFindStockByProductIdUsecase(stockRepository)
	stockController := controller.NewStockController(createStockUsecase, findStockByProductId)

	controllers := routes.Controllers{
		ProductController: productController,
		StockController: stockController,
	}

	router := routes.RouterInitializer(controllers)

	fmt.Printf("running on port %s - environment '%s'\n", PORT, ENV)
	log.Fatal(http.ListenAndServe(":"+PORT, router))
}