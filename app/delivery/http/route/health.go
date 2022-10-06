package route

import (
	"GoCleanMicroservice/abstract/domain/interactor"
	"github.com/gin-gonic/gin"
	"net/http"
)

var healthInteractor *interactor.HealthInteractor

func InitHealthInteractor(p any) {
	healthInteractor = p.(*interactor.HealthInteractor)
}

func Health(c *gin.Context) {
	var response = (*healthInteractor).Health()
	c.JSON(http.StatusOK, response)
}
