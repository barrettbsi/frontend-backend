package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func FrontendHome(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
