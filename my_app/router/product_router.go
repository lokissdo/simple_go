// router/product_router.go
package router

import (
	"github.com/gorilla/mux"
	"my_app/controller"
)

// RegisterProductRoutes configures the routes for product-related endpoints
func RegisterProductRoutes(router *mux.Router, productController *controller.ProductController) {
	router.HandleFunc("/products", productController.GetAllProducts).Methods("GET")
	router.HandleFunc("/products", productController.AddBatchProducts).Methods("POST")
}
