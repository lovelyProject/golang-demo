package order

import (
	"go/adv-example/configs"
	"go/adv-example/pkg/middleware"
	req "go/adv-example/pkg/req"
	"go/adv-example/pkg/res"
	"net/http"
	"strconv"
)

type Handler struct {
	Service *Service
	Config  *configs.Config
}

func NewOrderHandler(service *Service, config *configs.Config) *Handler {
	return &Handler{
		Service: service,
		Config:  config,
	}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.Handle("POST /order", middleware.IsAuthenticated(h.CreateOrder(), h.Config))
	router.Handle("GET /order/{id}", middleware.IsAuthenticated(h.GetById(), h.Config))
	router.Handle("GET /my-orders", middleware.IsAuthenticated(h.GetAllOrders(), h.Config))
}

func (h *Handler) CreateOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := req.HandleBody[CreateOrderRequest](w, r)
		if err != nil {
			res.Json(w, http.StatusBadRequest, err.Error())
			return
		}

		idRaw := r.Context().Value(middleware.ContextIdKey)
		userId, ok := idRaw.(uint)
		if !ok {
			res.Json(w, http.StatusInternalServerError, "id should be uint")
			return
		}
		order, err := h.Service.MakeOrder(userId, result.ProductsID)
		if err != nil {
			res.Json(w, http.StatusInternalServerError, err.Error())
			return
		}

		res.Json(w, http.StatusOK, order)
	}
}

func (h *Handler) GetAllOrders() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.Context().Value(middleware.ContextIdKey)
		userIdUint, ok := userId.(uint)
		if !ok {
			res.Json(w, http.StatusInternalServerError, "user id should be uint")
		}
		orders, err := h.Service.GetOrders(userIdUint)
		if err != nil {
			res.Json(w, http.StatusInternalServerError, err.Error())
		}

		res.Json(w, http.StatusOK, orders)
	}
}

func (h *Handler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orderId := r.PathValue("id")
		userId := r.Context().Value(middleware.ContextIdKey)
		userIdUint, ok := userId.(uint)
		if !ok {
			res.Json(w, http.StatusInternalServerError, "user id should be uint")
		}

		orderIdUint, err := strconv.ParseUint(orderId, 10, 32)
		if err != nil {
			res.Json(w, http.StatusInternalServerError, "order id should be uint")
			return
		}

		order, err := h.Service.GetOrder(userIdUint, uint(orderIdUint))
		if err != nil {
			res.Json(w, http.StatusInternalServerError, err.Error())
			return
		}

		res.Json(w, http.StatusOK, order)
	}
}
