package views

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Homepage(c *gin.Context) {
	c.HTML(http.StatusOK, "homepage", gin.H{
		"Title": "Hello World",
	})
}
