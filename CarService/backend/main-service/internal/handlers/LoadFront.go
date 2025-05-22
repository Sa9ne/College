package handlers

import "github.com/gin-gonic/gin"

func LoadFrontend(ctx *gin.Context) {
	ctx.File("/Users/user/important/College/CarService/frontend/build/web/index.html")
}
