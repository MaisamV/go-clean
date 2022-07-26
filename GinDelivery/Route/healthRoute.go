package Route

import (
	"GoCleanMicroservice/Domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

var healthInteractor *Domain.HealthInteractor

func InitHealthInteractor(p *Domain.HealthInteractor) {
	if healthInteractor == nil {
		healthInteractor = p
	}
}

func Health(c *gin.Context) {
	var response = (*healthInteractor).Health()
	c.JSON(http.StatusOK, response)
}
