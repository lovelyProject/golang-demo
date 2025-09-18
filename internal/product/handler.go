package product

import "go/adv-example/pkg/res"

type ProductHandler struct {
	Service *ProductService
}

func NewProductHandler(service *ProductService) *ProductHandler {
	return &ProductHandler{
		Service: service,
	}
}

func (h *ProductHandler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /", h.)
}

func (h *ProductHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.Service.GetAll()
		res.Json(w, http.StatusOK, h.Service.GetAll())
	}
}