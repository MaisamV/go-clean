package GinDelivery

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type Server struct {
	Engine *gin.Engine
}

func (m *Server) Init() error {
	err := os.Setenv("PORT", "3020")
	if err != nil {
		return err
	}
	m.Engine = gin.Default()
	m.Engine.GET("/ping", ping)
	return err
}

func (m *Server) Serv() error {
	return m.Engine.Run()
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
