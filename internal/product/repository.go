package product

import (
	"go/adv-example/db"
	model "go/adv-example/model"
)

type ProductRepo struct {
	Database *db.Db
}

func NewProductRepo(database *db.Db) *ProductRepo {
	return &ProductRepo{
		Database: database,
	}
}

func (r *ProductRepo) GetAll() ([]model.Product, error) {
	var products []model.Product
	result := r.Database.DB.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

func (r *ProductRepo) Create() (model.Product, error) {
	var product model.Product
	result := r.Database.DB.Create(&product)
	if result.Error != nil {
		return model.Product{}, result.Error
	}

	return product, nil
}
