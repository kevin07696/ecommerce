package auth

import (
	"net/http"

	"github.com/kevin07696/ecommerce/adapters/driving"
	"github.com/kevin07696/ecommerce/domain/auth/services"
)

type Handler struct {
	router  *http.ServeMux
	userAPI services.API
}

func Handle(router *http.ServeMux, userAPI services.API, middlewares ...driving.Middleware) *Handler {
	h := &Handler{
		router:  router,
		userAPI: userAPI,
	}

	h.initAppRoutes()

	middlewareChain := driving.MiddlewareChain(middlewares...)
	middlewareChain(h.router)

	return h
}
