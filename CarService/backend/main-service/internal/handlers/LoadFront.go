package handlers

import "github.com/gin-gonic/gin"

func LoadFrontend(ctx *gin.Context) {
	ctx.File("./frontend/index.html")
}
