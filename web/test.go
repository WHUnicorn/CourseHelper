package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Test() {
	router := gin.Default()
	router.Static("/static", "./static")
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})
	err := router.Run(":8000")
	if err != nil {
		return
	}
}
