package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleEchoHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg": "OK",
		"data": "hello",
	})
	return
}
