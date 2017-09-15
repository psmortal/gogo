package config

import "github.com/gin-gonic/gin"

func ENV() string{

	return gin.DebugMode
}

const (
	DB_NAME = ""
	DB_USERNAME = ""
	DB_PASSWORD = ""
	DB_HOST = "localhost"
	DB_PORT = "3306"
)