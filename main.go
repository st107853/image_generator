package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/golang-queue/queue"
	"github.com/golang-queue/queue/core"
	"golang.org/x/exp/maps"
)

type jobData struct {
	Id        string
	Generator string
}

type ImageAnswer struct {
	Id  string `json:"id"`
	URL string `json:"url"`
}

var sm sync.Map

func (j *jobData) Bytes() []byte {
	b, _ := json.Marshal(j)
	return b
}

func router() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("index.html")

	q := queue.NewPool(30, queue.WithFn(func(ctx context.Context, m core.QueuedMessage) error {
		j, _ := m.(*jobData)
		json.Unmarshal(m.Bytes(), &j)

		path := DrawOne(j.Generator)
		sm.Store(j.Id, path)
		fmt.Printf("Stored: %s:%s [%s]\n", j.Id, j.Generator, path)

		return nil
	}))

	imageRoute := r.Group("/image")
	{
		imageRoute.GET("/:generator", func(c *gin.Context) {
			generator := c.Param("generator")
			file := DrawOne(generator)
			c.Header("Content-Type", "image/png")
			c.File(file)
		})
	}

	r.Static("/img", "./img")

	r.GET("/list", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"keys": maps.Keys(DRAWINGS),
		})
	})

	newRoute := r.Group("/new")
	{
		newRoute.GET("/load/:id", func(c *gin.Context) {
			id := c.Param("id")
			path, ok := sm.Load(id)

			if ok {
				fmt.Printf("Found %s for id: %s\n", path, id)
				c.Writer.Header().Set("Content-Type", "image/png")
				c.File(fmt.Sprintf("%s", path))
				c.Writer.WriteHeader(200)
				//	fmt.Println(c.Header().Get("Content-Type"))
			} else {
				fmt.Printf("Path not found for id: %s\n", id)
				c.Writer.Header().Set("Content-Type", "image/jpg")
				c.Header("Cache-Control", "no-cache")
				c.File("img/loading.jpg")
				c.Writer.WriteHeader(200)

			}
		})
	}
	newRoute.GET("/:generator", func(c *gin.Context) {
		generator := c.Param("generator")
		newJob := jobData{
			Id:        strconv.Itoa(rand.Int()),
			Generator: generator,
		}
		q.Queue(&newJob)
		res := ImageAnswer{
			Id:  newJob.Id,
			URL: "http://" + c.Request.Host + "/new/load/" + newJob.Id,
		}
		c.JSON(200, res)
	})

	return r
}

func main() {
	router().Run()
}
