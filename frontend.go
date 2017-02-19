package febe

import (
	"fmt"
	"log"
	"os"

	"github.com/barrettbsi/frontend-backend/controllers"
	"github.com/gin-gonic/gin"
)

func StartFrontend() {
	log.Printf("Starting frontend...")

	r := gin.Default()

	r.LoadHTMLGlob("templates/frontend/*")

	r.GET("/", controllers.FrontendHome)

	r.Run(fmt.Sprintf(":%v", os.Getenv("PORT")))
}

func StartBackend() {
	log.Printf("Starting backend...")

	r := gin.Default()

	r.LoadHTMLGlob("templates/backend/*")

	r.GET("/", controllers.BackendHome)

	r.Run(fmt.Sprintf(":%v", os.Getenv("PORT")))
}
