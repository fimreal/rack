package ipquery

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func help(ctx *gin.Context) {
	ctx.String(http.StatusOK, `/ip
/ipinfo
/ipinfo/:ip
/ip2dec/:ip
/dec2ip/:ip
/cidr2ip/:ip/:cidr
/ip2cidr/:ipfrom/:ipto
/pipv6/:ip
/pip/:ip
/rip/:ip
`)
}
