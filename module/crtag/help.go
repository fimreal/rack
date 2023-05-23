package crtag

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func help(ctx *gin.Context) {
	ctx.String(http.StatusOK, `/dockerhub/<owner>|[library]/<image name>[/tags|/images]
`)
}
