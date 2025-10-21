package link

import (
	"fmt"
	"go/adv-example/configs"
	"go/adv-example/pkg/event"
	middleware "go/adv-example/pkg/middleware"
	requestHandler "go/adv-example/pkg/req"
	"go/adv-example/pkg/res"
	"net/http"
	"strconv"
)

type LinkHandlerDeps struct {
	LinkService *LinkService
	Config      *configs.Config
	EventBus    *event.EventBus
}

type LinkHandler struct {
	LinkService *LinkService
	Config      *configs.Config
	EventBus    *event.EventBus
}

func NewLinkHandler(deps LinkHandlerDeps, config *configs.Config) *LinkHandler {
	return &LinkHandler{
		LinkService: deps.LinkService,
		Config:      deps.Config,
		EventBus:    deps.EventBus,
	}
}

func (linkHandler *LinkHandler) RegisterRoutes(router *http.ServeMux) {
	router.Handle("POST /link", middleware.IsAuthenticated(linkHandler.CreateLink(), linkHandler.Config))
	router.Handle("GET /link", middleware.IsAuthenticated(linkHandler.GetAll(), linkHandler.Config))
	router.HandleFunc("GET /link/{hash}", linkHandler.GoTo())
	router.HandleFunc("PUT /link/{hash}", linkHandler.UpdateLink())
	router.HandleFunc("DELETE /link/{hash}", linkHandler.DeleteLink())
}

func (linkHandler *LinkHandler) CreateLink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := r.Context().Value(middleware.ContextEmailKey).(string)
		fmt.Println(email)
		body, err := requestHandler.HandleBody[LinkCreateRequest](w, r)
		if err != nil {
			res.Json(w, http.StatusBadRequest, err.Error())
			return
		}

		createdLink, err := linkHandler.LinkService.CreateUniqueHash(body.Url)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res.Json(w, http.StatusCreated, createdLink)
	}
}

func (linkHandler *LinkHandler) GetLink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res.Json(w, http.StatusOK, "Link fetched")
	}
}

func (linkHandler *LinkHandler) UpdateLink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res.Json(w, http.StatusOK, "Link updated")
	}
}

func (linkHandler *LinkHandler) DeleteLink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.ParseInt(idStr, 10, 32)
		if err != nil {
			res.Json(w, http.StatusBadRequest, err.Error())
			return
		}
		_, err = linkHandler.LinkService.FindById(uint(id))
		if err != nil {
			res.Json(w, http.StatusBadRequest, err.Error())
			return
		}
		err = linkHandler.LinkService.Delete(uint(id))
		if err != nil {
			res.Json(w, http.StatusBadRequest, err.Error())
			return
		}
		res.Json(w, http.StatusOK, "Link deleted")
	}
}

func (linkHandler *LinkHandler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		link, err := linkHandler.LinkService.Repo.GetByHash(hash)
		go linkHandler.EventBus.Publish(event.Event{
			Type: event.EventLinkVisited,
			Data: link.ID,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		http.Redirect(w, r, link.Url, http.StatusTemporaryRedirect)
	}
}

func (linkHandler *LinkHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil {
			limit = 20
		}
		offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
		if err != nil {
			offset = 0
		}
		result, err := linkHandler.LinkService.GetAll(limit, offset)
		if err != nil {
			res.Json(w, http.StatusInternalServerError, err.Error())
			return
		}

		res.Json(w, http.StatusOK, result)
	}
}
