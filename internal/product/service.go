package product

import model "go/adv-example/internal/model"

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

func (s *ProductService) Create(body *ProductCreateRequest) (model.Product, error) {
	product, err := s.Repo.Create(body)
	if err != nil {
		return model.Product{}, err
	}
	return product, nil
}

func (s *ProductService) Update(product *model.Product) error {
	_, err := s.Repo.Update(product)
	if err != nil {
		return err
	}
	return nil
}

func (s *ProductService) Delete(product *model.Product) error {
	err := s.Repo.Delete(product)
	if err != nil {
		return err
	}

	return nil
}

func (s *ProductService) GetById(product *model.Product) (model.Product, error) {
	result, err := s.Repo.GetById(product)
	if err != nil {
		return model.Product{}, err
	}
	return result, nil
}
