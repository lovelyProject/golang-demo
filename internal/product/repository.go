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

func (r *ProductRepo) Create(body *ProductCreateRequest) (model.Product, error) {
	product := model.Product{
		Name:        body.Name,
		Description: body.Description,
		Type:        body.Type,
		Price:       body.Price,
		Images:      body.Images,
	}
	result := r.Database.DB.Create(&product)

	if result.Error != nil {
		return model.Product{}, result.Error
	}

	return product, nil
}
