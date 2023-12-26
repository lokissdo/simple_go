// controller/product_controller.go
package controller

import (
	"encoding/json"
	"net/http"
	"log"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"github.com/gorilla/mux"
	"my_app/model"
	"my_app/repository"
)

// ProductController handles HTTP requests related to products
type ProductController struct {
	productRepo *repository.ProductRepository
}

// NewProductController creates a new instance of ProductController
func NewProductController(productRepo *repository.ProductRepository) *ProductController {
	return &ProductController{
		productRepo: productRepo,
	}
}

// GetAllProducts handles requests to get all products
func (pc *ProductController) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := pc.productRepo.GetAllProducts()
	if err != nil {
		http.Error(w, "Error fetching products: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// GetProductByID handles requests to get a product by ID
func (pc *ProductController) GetProductByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID := vars["id"]

	product, err := pc.productRepo.GetProductByID(productID)
	if err != nil {
		http.Error(w, "Error fetching product: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if product == nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// AddProduct handles requests to add a new product
func (pc *ProductController) AddProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct models.Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = pc.productRepo.AddProduct(&newProduct)
	if err != nil {
		http.Error(w, "Error adding product: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// AddBatchProducts handles requests to add multiple products
func (pc *ProductController) AddBatchProducts(w http.ResponseWriter, r *http.Request) {
	var newProducts []models.Product
	log.Println(r.Body)
	jsonData, err1 := bson.MarshalJSON(r.Body)
	if err1 != nil {
		fmt.Println("Error converting BSON to JSON:", err1)
		return
	}

	fmt.Println("data from request"+string(jsonData))
	decode := json.NewDecoder(r.Body)
	log.Println(decode)
	err := decode.Decode(&newProducts)
	if err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = pc.productRepo.AddBatchProducts(newProducts)
	if err != nil {
		http.Error(w, "Error adding products: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
