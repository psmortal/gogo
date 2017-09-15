package main

import (
	"github.com/gin-gonic/gin"
	userauth "./handler/auth"
	"os"
	"log"
	"./config"
)


func main()  {

	gin.SetMode(config.ENV())

	router := gin.Default()

	router.Use(gin.Logger())


	router.POST("/",userauth.Login)



	err := router.Run(":8080")
	if err != nil {
		log.Printf("启动失败......",err.Error())
		os.Exit(10000)
	}

}
