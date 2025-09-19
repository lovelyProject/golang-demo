package product

import (
	req "go/adv-example/pkg/req"
	"go/adv-example/pkg/res"
	"net/http"
)

type ProductHandler struct {
	Service *ProductService
}

func NewProductHandler(service *ProductService) *ProductHandler {
	return &ProductHandler{
		Service: service,
	}
}

func (h *ProductHandler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /products", h.GetAll())
	router.HandleFunc("POST /products", h.Create())
	router.HandleFunc("PUT /products/{id}", h.Update())
	router.HandleFunc("GET /products/{id}", h.GetById())
	router.HandleFunc("DELETE /products/{id}", h.Delete())
}

func (h *ProductHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products, err := h.Service.GetAll()
		if err != nil {
			res.Json(w, http.StatusInternalServerError, err.Error())
		}
		res.Json(w, http.StatusOK, products)
	}
}

func (h *ProductHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[ProductCreateRequest](w, r)
		if err != nil {
			res.Json(w, http.StatusBadRequest, err)
		}
		result, err := h.Service.Create(body)
		if err != nil {
			res.Json(w, http.StatusInternalServerError, err)
		}
		res.Json(w, http.StatusOK, result)
	}
}

func (h *ProductHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.ParseInt(idStr, 10, 64)

		res.Json(w, http.StatusOK, "updated")
	}
}

func (h *ProductHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res.Json(w, http.StatusOK, "GetById")
	}
}

func (h *ProductHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res.Json(w, http.StatusOK, "Delete")
	}
}
