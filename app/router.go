package app

import "github.com/gin-gonic/gin"

func InitRoute(router *gin.Engine) {
	router.GET("/profile", GetProfile)
	router.POST("/profile", CreateProfile)
	router.DELETE("/profile", DeleteProfile)
}
