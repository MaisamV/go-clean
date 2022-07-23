package GinDelivery

import (
	"github.com/gin-gonic/gin"
	"os"
)

type Server struct {
	Engine      *gin.Engine
	Interactors *InteractorPackage
}

func (m *Server) Init() error {
	err := os.Setenv("PORT", "3020")
	if err != nil {
		return err
	}
	m.setInteractors()
	m.Engine = gin.Default()
	m.addRoutes()
	return err
}

func (m *Server) Serv() error {
	return m.Engine.Run()
}
