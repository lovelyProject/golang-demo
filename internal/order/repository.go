package order

import (
	"errors"
	"go/adv-example/db"
	"go/adv-example/internal/model"
)

type OrderRepository struct {
	db *db.Db
}

func NewOrderRepository(db *db.Db) *OrderRepository {
	return &OrderRepository{db}
}

func (r *OrderRepository) CreateWithProducts(userID uint, productsId []uint) (model.Order, error) {
	var products []model.Product
	if result := r.db.Where("id in ?", productsId).Find(&products); result.Error != nil {
		return model.Order{}, result.Error
	}

	if len(products) != len(productsId) {
		return model.Order{}, errors.New("products count not match")
	}

	order := model.Order{
		UserID:   userID,
		Products: products,
	}

	if result := r.db.Create(&order); result.Error != nil {
		return model.Order{}, result.Error
	}

	return order, nil
}

func (r *OrderRepository) GetByUserId(userID uint) ([]model.Order, error) {
	var orders []model.Order
	err := r.db.Where("user_id = ?", userID).
		Preload("Products").
		Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *OrderRepository) GetById(id, userId uint) (model.Order, error) {
	var order model.Order
	err := r.db.Where("id = ? AND user_id = ?", id, userId).
		Preload("Products").
		First(&order).Error
	if err != nil {
		return order, err
	}
	return order, nil
}
