package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func helloWorld(c *gin.Context) {
	c.String(http.StatusNotFound, "Hey, world!\n")
}
