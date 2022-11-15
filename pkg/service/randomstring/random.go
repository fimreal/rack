package randomstring

import (
	"net/http"

	"github.com/fimreal/goutils/crypto/random"
	"github.com/fimreal/goutils/ezap"
	"github.com/gin-gonic/gin"
)

func GenAlpha(length int) string {
	return random.Alpha(length)
}

func GenRandomString(ctx *gin.Context) {
	var randreq RandomCode
	if err := ctx.ShouldBindJSON(&randreq); err != nil {
		ezap.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var res string
	// if len(randreq.charpool) != 0 {
	// 	do sth
	// }
	switch randreq.Chartype {
	case "number":
		res = random.Num(randreq.Length)
	case "alpha":
		res = random.Alpha(randreq.Length)
	case "lowalpha":
		res = random.LowAlpha(randreq.Length)
	case "upalpha":
		res = random.UpAlpha(randreq.Length)
	default:
		res = random.GenPassword(randreq.Length)
	}
	ctx.JSON(http.StatusOK, gin.H{"result": res})
}
