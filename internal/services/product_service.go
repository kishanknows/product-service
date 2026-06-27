package services

import (
	"context"
	"errors"
	"time"

	"github.com/kishanknows/product-service/internal/database"
	apperr "github.com/kishanknows/product-service/internal/errors"
	"github.com/kishanknows/product-service/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type ProductService struct {
}

func NewProductService() *ProductService {
	return &ProductService{}
}

func (s *ProductService) CreateProduct(ctx context.Context, product *models.Product) *apperr.AppError {
	collection := database.DB.Collection("products")

	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	_, err := collection.InsertOne(ctx, product)

	if err != nil {
		return apperr.ErrInternalServer
	}

	return nil
}

func (s *ProductService) GetAllProducts(ctx context.Context) ([]models.Product, *apperr.AppError) {
	collection := database.DB.Collection("products")

	cursor, err := collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, apperr.ErrInternalServer
	}

	defer cursor.Close(ctx)

	var products []models.Product

	if err := cursor.All(ctx, &products); err != nil {
		return nil, apperr.ErrInternalServer
	}

	return products, nil
}

func (s *ProductService) GetProductById(ctx context.Context, id string) (*models.Product, *apperr.AppError) {
	collection := database.DB.Collection("products")

	objectID, err := bson.ObjectIDFromHex(id)

	if err != nil {
		return nil, apperr.ErrInvalidProductID
	}

	var product models.Product

	err = collection.FindOne(ctx, bson.M{
		"_id": objectID,
	}).Decode(&product)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, apperr.ErrProductNotFound
		}
		return nil, apperr.ErrInternalServer
	}

	return &product, nil
}

func (s *ProductService) DeleteProductById(ctx context.Context, id string, sellerID int) *apperr.AppError {
	collection := database.DB.Collection("products")

	objectID, err := bson.ObjectIDFromHex(id)

	if err != nil {
		return apperr.ErrInvalidProductID
	}

	res, err := collection.DeleteOne(ctx, bson.M{
		"_id": objectID,
		"seller_id": sellerID,
	})

	if res.DeletedCount == 0 {
		return apperr.ErrProductNotFound
	}

	if err != nil {
		return apperr.ErrInternalServer
	}

	return nil
}

func (s *ProductService) ReplaceProductById(ctx context.Context, id string, product *models.Product, sellerID int) *apperr.AppError {
	collection := database.DB.Collection("products")

	objectID, err := bson.ObjectIDFromHex(id)

	if err != nil {
		return apperr.ErrInvalidProductID
	}

	updatedProduct := bson.M{
		"$set": bson.M{
			"name": product.Name,
			"price": product.Price,
			"updated_at": time.Now(),
		},
	}

	res, err := collection.UpdateOne(ctx, bson.M{"_id": objectID, "seller_id": sellerID}, updatedProduct)

	if res.ModifiedCount == 0 {
		return apperr.ErrProductNotFound
	}

	if err != nil {
		return apperr.ErrInternalServer
	}
	
	return nil
}

func (s *ProductService) UpdateProductById(ctx context.Context, id string, req *models.UpdateProductRequest, sellerID int) (*models.Product, *apperr.AppError) {
	collection := database.DB.Collection("products")

	objectID, err := bson.ObjectIDFromHex(id)

	if err != nil {
		return nil, apperr.ErrInvalidProductID
	}

	updateFields := bson.M{}

	if req.Name != nil {
		updateFields["name"] = req.Name
	}

	if req.Price != nil {
		updateFields["price"] = req.Price
	}

	updateFields["updated_at"] = time.Now()

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var product models.Product

	err = collection.FindOneAndUpdate(
		ctx,
		bson.M{"_id": objectID, "seller_id": sellerID},
		bson.M{"$set": updateFields},
		opts,
	).Decode(&product)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, apperr.ErrProductNotFound
		}
		return nil, apperr.ErrInternalServer
	}

	return &product, nil
}