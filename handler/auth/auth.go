package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(ctx *gin.Context){
	ctx.JSON(http.StatusOK,"login")
}
