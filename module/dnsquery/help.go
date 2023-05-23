package dnsquery

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func help(ctx *gin.Context) {
	ctx.String(http.StatusOK, `
/dns/:host
/whois/:domain
`)
}
