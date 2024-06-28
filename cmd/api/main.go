package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"product-warehouse/cmd/api/controller"
	"product-warehouse/cmd/api/routes"
	"product-warehouse/config"
	port "product-warehouse/internal/port/repository"
	"product-warehouse/internal/repository/inMemory"
	"product-warehouse/internal/repository/postgre"
	productUsecase "product-warehouse/internal/usecase/product"
	stockUsecase "product-warehouse/internal/usecase/stock"
)

func main() {
	PORT := config.GetEnv("PORT")
	ENV := config.GetEnv("ENV")

	var (
		host = config.GetEnv("DB_HOST")
		db_port = config.GetEnv("DB_PORT")
		user = config.GetEnv("DB_USERNAME")
		password = config.GetEnv("DB_PASSWORD")
		dbname = config.GetEnv("DB_NAME")
	)

	var (
		productRepository port.ProductRepository
		stockRepository port.StockRepository
	)
	

	switch ENV {
	case "prod":
		psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, db_port, user, password, dbname)
		
		dbConnection := initDbConnection(psqlInfo)

		productRepository = postgre.NewPostgreProductRepository(dbConnection)
		stockRepository = postgre.NewPostgreStockRepository(dbConnection)
	default:
		productRepository = inMemory.NewInMemoryProductRepository()
		stockRepository = inMemory.NewInMemoryStockRepository()
	}
	
	startServer(PORT, ENV, productRepository, stockRepository)
}

func initDbConnection(psqlInfo string) *sql.DB {
	return config.InitConfig(psqlInfo)
}

func startServer(PORT, ENV string, productRepository port.ProductRepository, stockRepository port.StockRepository) {
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

	log.Printf("running on port %s - environment '%s'\n", PORT, ENV)
	log.Fatal(http.ListenAndServe(":"+PORT, router))
}

