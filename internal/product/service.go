package product

import model "go/adv-example/model"

type ProductService struct {
	Repo *ProductRepo
}

func NewProductService(repo *ProductRepo) *ProductService {
	return &ProductService{
		Repo: repo,
	}
}

func (s *ProductService) GetAll() ([]model.Product, error) {
	products, err := s.Repo.GetAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}
