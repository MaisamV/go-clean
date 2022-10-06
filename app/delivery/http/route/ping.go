package route

import (
	"GoCleanMicroservice/abstract/domain/interactor"
	"github.com/gin-gonic/gin"
	"net/http"
)

var pingInteractor *interactor.PingInteractor

func InitPingInteractor(p any) {
	pingInteractor = p.(*interactor.PingInteractor)
}

func Ping(c *gin.Context) {
	res, err := (*pingInteractor).Ping()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "failed to ping"})
	}
	c.JSON(http.StatusOK, res)
}
