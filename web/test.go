package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Test() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})
	router.Run(":8000")
}
