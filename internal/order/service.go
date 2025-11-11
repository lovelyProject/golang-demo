package order

import "go/adv-example/internal/model"

type Service struct {
	repo *OrderRepository
}

func NewService(repo *OrderRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) MakeOrder(userId uint, productsId []uint) (model.Order, error) {
	result, err := s.repo.CreateWithProducts(userId, productsId)
	if err != nil {
		return model.Order{}, err
	}

	return result, nil
}

func (s *Service) GetOrders(userId uint) ([]model.Order, error) {
	result, err := s.repo.GetByUserId(userId)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *Service) GetOrder(userId, orderId uint) (model.Order, error) {
	order, err := s.repo.GetById(orderId, userId)
	if err != nil {
		return model.Order{}, err
	}

	return order, nil
}
