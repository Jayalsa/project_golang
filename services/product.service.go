package services

import (
	"context"
	"fmt"
	"jayalsa/project_golang/entities"
	"jayalsa/project_golang/interfaces"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductService struct {
	Product *mongo.Collection
}

func InitProductService(collection *mongo.Collection) interfaces.IProduct {

	return &ProductService{Product: collection}
}

func (p *ProductService) Insert(product *entities.Product) (string, error) {
	product.ID = primitive.NewObjectID()
	_, err := p.Product.InsertOne(context.Background(), product)
	if err != nil {
		return "", err
	} else {
		return "Record Inserted Successfully", nil
	}
}
func (p *ProductService) GetProducts() ([]*entities.Product, error) {
	result, err := p.Product.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	} else {

		fmt.Println(result)
		var products []*entities.Product
		for result.Next(context.TODO()) {
			product := &entities.Product{}
			err := result.Decode(product)

			if err != nil {
				return nil, err
			}
			products = append(products, product)
		}
		if err := result.Err(); err != nil {
			return nil, err
		}
		if len(products) == 0 {
			return []*entities.Product{}, nil
		}
		return products, nil
	}

}
func (p *ProductService) GetProductByID(id string) (*entities.Product, error) {
	filter := bson.M{"_id": id}
	var product entities.Product
	err := p.Product.FindOne(context.TODO(), filter).Decode(&product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &product, nil
}
