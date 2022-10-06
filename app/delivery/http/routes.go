package http

import (
	"GoCleanMicroservice/app/delivery/http/route"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (m *Server) addRoutes() {
	addRoute("/ping", route.Ping, m, route.InitPingInteractor, m.Interactors.Ping, http.MethodGet)
	addRoute("/health", route.Health, m, route.InitHealthInteractor, m.Interactors.Health, http.MethodGet)
}

func addRoute[T any](relativePath string, handler gin.HandlerFunc, m *Server, setter func(any), interactor T, httpMethod string) {
	panicOnNull(relativePath, interactor)
	m.Engine.Handle(httpMethod, relativePath, handler)
	setter(interactor)
}

func panicOnNull(interactorName string, interactor interface{}) {
	if interactor == nil {
		panic(any(interactorName + " interactor cannot be null."))
	}
}
