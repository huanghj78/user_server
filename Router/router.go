package Router

import (
	"fmt"
	"log"
	"user_server/Apps/common"
	"user_server/Config"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()
	router.Use()

	router.Static("/assets", fmt.Sprintf("%s/assets", Config.Server.Frontend))
	router.StaticFile("/", Config.Server.Frontend)
	//router.Static("/upload", Config.Server.FilePath)

	//router.Use(Middlewares.)

	v1 := router.Group("v1/")
	{
		v1.GET("echo/hello", common.HandleEchoHello)
	}

	err := router.Run(fmt.Sprintf("%s:%d", Config.Server.Address, Config.Server.Port))
	if err != nil {
		log.Fatal(err)
		return
	}

}
