package route

import (
	"GoCleanMicroservice/abstract/domain/interactor"
	"github.com/gin-gonic/gin"
	"net/http"
)

var healthInteractor *interactor.HealthInteractor

func InitHealthInteractor(p *interactor.HealthInteractor) {
	if healthInteractor == nil {
		healthInteractor = p
	}
}

func Health(c *gin.Context) {
	var response = (*healthInteractor).Health()
	c.JSON(http.StatusOK, response)
}
