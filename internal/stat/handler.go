package stat

import (
	"go/adv-example/configs"
	"go/adv-example/pkg/middleware"
	"go/adv-example/pkg/res"
	"net/http"
	"time"
)

type Handler struct {
	Service *StatService
	Config  *configs.Config
}

func NewStatHandler(service *StatService, config *configs.Config) *Handler {
	return &Handler{
		Service: service,
		Config:  config,
	}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.Handle("GET /stat", middleware.IsAuthenticated(h.GetAll(), h.Config))
}

func (h *Handler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		from := query.Get("from")
		to := query.Get("to")
		by := query.Get("by")

		parsedFrom, err := time.Parse("2006-01-02", from)
		if err != nil {
			res.Json(w, http.StatusBadRequest, "invalid from")
			return
		}
		parsedTo, err := time.Parse("2006-01-02", to)
		if err != nil {
			res.Json(w, http.StatusBadRequest, "invalid to")
			return
		}

		type Result struct {
			From time.Time `json:from`
			To   time.Time `json:to`
			By   string    `json:by`
		}
		result := Result{
			From: parsedFrom,
			To:   parsedTo,
			By:   by,
		}
		res.Json(w, http.StatusOK, result)
	}
}
