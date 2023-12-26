// router/router.go
package router

import (
	"github.com/gorilla/mux"
	"my_app/config"
	"my_app/controller"
	"my_app/repository"
	"net/http"
	"html/template"
	"log"
	

)

// Router struct represents a router configuration
type Router struct {
	mainRouter *mux.Router
}

// NewRouter creates a new instance of Router
func NewRouter() *Router {
	return &Router{
		mainRouter: mux.NewRouter(),
	}
}

// SetupRoutes configures the routes for the application
func (r *Router) SetupRoutes() {
	// Initialize MongoDB connection
	dbConnector := config.NewDBConnector()

	// Connect to MongoDB
	err := dbConnector.Connect()
	if err != nil {
		log.Fatal(err)
	}


	var productController *controller.ProductController = controller.NewProductController(repository.NewProductRepository(dbConnector.GetDatabase()))
	// Initialize controllers
	r.mainRouter.HandleFunc("/", IndexHandler)
	// Register routes from different router files
	RegisterProductRoutes(r.mainRouter, productController)
}

// GetHandler returns the main router as an ----http.Handler----  *mux.Router
func (r *Router) GetHandler()  *mux.Router {
	return r.mainRouter
}


func IndexHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("static/index.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    tmpl.Execute(w, nil)
}