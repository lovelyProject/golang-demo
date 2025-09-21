package product

import "github.com/lib/pq"

type ProductCreateRequest struct {
	Name        string         `json:"name" validate:"required"`
	Description string         `json:"desc"`
	Type        string         `json:"type"`
	Price       float64        `json:"price" validate:"required"`
	Images      pq.StringArray `json:"images"`
}

type ProductUpdateRequest struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
