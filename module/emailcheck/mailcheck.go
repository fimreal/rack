package emailcheck

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

var (
	MEETCHOPRA_API   = "https://verifier.meetchopra.com/verify/"
	MEETCHOPRA_TOKEN = "2b1e810090b21cab8a8753ec6bd1f091f35ce5a28bc6135561a12eb0814ae6de5e5603dc00fabffba179541b31d27dac"
)

// 判断是否为真实 email
func IsRealMail(c *gin.Context) {
	api := MEETCHOPRA_API + c.Param("email") + "?token=" + MEETCHOPRA_TOKEN

	remote, _ := url.Parse(api)
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Director = func(req *http.Request) {
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.Header = c.Request.Header
		req.URL.Host = remote.Host
		req.URL.Path = remote.Path
		req.URL.RawQuery = remote.RawQuery
	}
	proxy.ServeHTTP(c.Writer, c.Request)
}
