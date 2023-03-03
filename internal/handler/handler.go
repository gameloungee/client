package handler

import "github.com/gin-gonic/gin"

func InitHanlder() *gin.Engine {
	r := gin.New()
	r.LoadHTMLGlob("pages/*")
	r.Static("/static", "/static/")

	r.GET("/", IndexHandler)

	return r
}
