package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/maps"
)

func router() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("template.tmpl")

	imageRoute := r.Group("/image")
	{
		imageRoute.GET("/:generator", func(c *gin.Context) {
			generator := c.Param("generator")
			file := DrawOne(generator)
			c.Header("Content-Type", "image/png")
			c.File(file)
		})
	}

	listRoute := r.Group("/list")
	{
		listRoute.GET("/simple", func(c *gin.Context) {
			c.HTML(http.StatusOK, "template.tmpl", gin.H{
				"keys": maps.Keys(DRAWINGS),
			})
		})
	}
	return r
}

func main() {
	router().Run()
}
