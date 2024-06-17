package routes

import (
	"product-warehouse/cmd/api/controller"

	"github.com/go-chi/chi/v5"
)

type Controllers struct {
	ProductController *controller.ProductController
	StockController *controller.StockController
}

func Routes(cont Controllers, c *chi.Mux) {
	c.Route("/api/v1", func(r chi.Router) {
		r.Group(func(router chi.Router) {
			router.Post("/product", cont.ProductController.Create)
			router.Get("/product/{id}", cont.ProductController.FindById)
			router.Post("/stock", cont.StockController.Create)
			router.Get("/stock/{productId}", cont.StockController.FindStockByProductId)
		})
	})
}