package Route

import (
	"GoCleanMicroservice/Domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

var pingInteractor *Domain.PingInteractor

func InitPingInteractor(p *Domain.PingInteractor) {
	if pingInteractor == nil {
		pingInteractor = p
	}
}

func Ping(c *gin.Context) {
	res, err := (*pingInteractor).Ping()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "failed to ping"})
	}
	c.JSON(http.StatusOK, res)
}
