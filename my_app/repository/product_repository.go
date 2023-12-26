// repository/product_repository.go
package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"my_app/model"
)

// ProductRepository is a repository for the Product model.
type ProductRepository struct {
	Collection *mongo.Collection
}

const CollectionName = "products"
// NewProductRepository creates a new ProductRepository with a specific collection name.
func NewProductRepository(database *mongo.Database) *ProductRepository {
	return &ProductRepository{Collection: database.Collection(CollectionName)}
}


func (r *ProductRepository) GetAllProducts() ([]models.Product, error) {
	var products []models.Product

	cursor, err := r.Collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var product models.Product
		err := cursor.Decode(&product)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}



// GetProductByID retrieves a product by its ID.
func (r *ProductRepository) GetProductByID(id string) (*models.Product, error) {
	var product models.Product
	err := r.Collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&product)
	return &product, err
}

// AddProduct adds a single product to the database.
func (r *ProductRepository) AddProduct(product *models.Product) error {
	_, err := r.Collection.InsertOne(context.Background(), product)
	return err
}

// AddBatchProducts adds multiple products to the database.
func (r *ProductRepository) AddBatchProducts(products []models.Product) error {
	// Convert products to interface{} to insert multiple documents
	documents := make([]interface{}, len(products))
	for i, p := range products {
		documents[i] = p
	}

	_, err := r.Collection.InsertMany(context.Background(), documents)
	return err
}