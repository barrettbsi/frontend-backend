package febe

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func StartFrontend() {
	log.Printf("Starting frontend...")

	r := gin.Default()

	r.LoadHTMLGlob("templates/frontend/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.Run(fmt.Sprintf(":%v", os.Getenv("PORT")))
}

func StartBackend() {
	log.Printf("Starting backend...")

	r := gin.Default()

	r.LoadHTMLGlob("templates/backend/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.Run(fmt.Sprintf(":%v", os.Getenv("PORT")))
}
