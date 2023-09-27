package interfaces

import "jayalsa/project_golang/entities"

type IProduct interface {
	Insert(product *entities.Product) (string, error)
	GetProducts() ([]*entities.Product, error)
	GetProductByID(id string) (*entities.Product, error)
}
