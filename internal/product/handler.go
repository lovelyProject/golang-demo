package product

import (
	model "go/adv-example/model"
	req "go/adv-example/pkg/req"
	"go/adv-example/pkg/res"
	"gorm.io/gorm"
	"net/http"
	"strconv"
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
	router.HandleFunc("PATCH /products/{id}", h.Update())
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
		body, err := req.HandleBody[ProductUpdateRequest](w, r)
		if err != nil {
			res.Json(w, http.StatusBadRequest, err)
		}
		idStr := r.PathValue("id")
		id, err := strconv.ParseInt(idStr, 10, 32)
		if err != nil {
			res.Json(w, http.StatusBadRequest, err)
		}
		err = h.Service.Update(&model.Product{
			Model: gorm.Model{
				ID: uint(id),
			},
			Name:  body.Name,
			Price: body.Price,
		})

		if err != nil {
			res.Json(w, http.StatusInternalServerError, err)
		}
		res.Json(w, http.StatusOK, "updated")
	}
}

func (h *ProductHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.ParseInt(idStr, 10, 32)
		if err != nil {
			res.Json(w, http.StatusBadRequest, err)
			return
		}

		result, err := h.Service.GetById(&model.Product{
			Model: gorm.Model{
				ID: uint(id),
			},
		})

		if err != nil {
			res.Json(w, http.StatusNotFound, err)
			return
		}

		res.Json(w, http.StatusOK, result)
	}
}

func (h *ProductHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		strId := r.PathValue("id")
		id, err := strconv.ParseInt(strId, 10, 32)
		if err != nil {
			res.Json(w, http.StatusBadRequest, err)
		}
		err = h.Service.Delete(&model.Product{
			Model: gorm.Model{
				ID: uint(id),
			},
		})

		if err != nil {
			res.Json(w, http.StatusInternalServerError, err)
		}

		res.Json(w, http.StatusOK, "deleted")
	}
}
