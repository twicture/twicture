package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Static("/assets", "./public")
	router.LoadHTMLGlob("./public/*.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	if err := router.Run(); err != nil {
		fmt.Println(err.Error())
	}
}
