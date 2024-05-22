package fileserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func help(ctx *gin.Context) {
	ctx.String(http.StatusOK, `/		browse file directory
/upload	use POST method to upload file
/dev webdev service
`)
}
