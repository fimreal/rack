package common

import (
	"github.com/fimreal/rack/pkg/utils"
	"github.com/gin-gonic/gin"
)

func SixNumber(c *gin.Context) {
	c.String(200, utils.RandomString("1234567890", 6))
}

func GenRandomPassword(c *gin.Context) {
	c.String(200, utils.RandomString("1234567890qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM+=-@#.$%^*", 16))
}
