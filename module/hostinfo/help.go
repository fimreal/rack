package hostinfo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func help(ctx *gin.Context) {
	ctx.String(http.StatusOK, `/hostip	show host ip list
/hostname	show hostname
`)
}
