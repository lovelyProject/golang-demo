package link

import (
	requestHandler "go/adv-example/pkg/req"
	"go/adv-example/pkg/res"
	"net/http"
	"strconv"
)

type LinkHandlerDeps struct {
	LinkService *LinkService
}

type LinkHandler struct {
	LinkService *LinkService
}

func NewLinkHandler(deps LinkHandlerDeps) *LinkHandler {
	return &LinkHandler{
		LinkService: deps.LinkService,
	}
}

func (linkHandler *LinkHandler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /link", linkHandler.CreateLink())
	router.HandleFunc("GET /link/{hash}", linkHandler.GoTo())
	router.HandleFunc("PUT /link/{hash}", linkHandler.UpdateLink())
	router.HandleFunc("DELETE /link/{hash}", linkHandler.DeleteLink())
}

func (linkHandler *LinkHandler) CreateLink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
		// hash := r.PathValue("hash")
		// link, err := linkHandler.LinkService.LinkRepo.GetByHash(hash)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusBadRequest)
		// 	return
		// }

		// http.Redirect(w, r, link.Url, http.StatusTemporaryRedirect)
	}
}
