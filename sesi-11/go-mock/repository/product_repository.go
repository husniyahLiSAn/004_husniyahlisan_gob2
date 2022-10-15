package repository

import "go-mock/entity"

type ProductRepository interface {
	FindById(id string) *entity.Product
}
