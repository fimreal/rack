package dnsquery

import (
	"encoding/json"
	"errors"
	"io"
	"net"
	"net/http"
	"strings"
	"syscall"
	"time"

	"github.com/fimreal/goutils/ezap"
	"github.com/gin-gonic/gin"
)

const (
	CLOUDFLAREDOHAPI = "https://cloudflare-dns.com/dns-query"
	APIPROXY         = "http://cfdown.2fw.top/"
)

type CFDOHResp struct {
	Status   int  `json:"Status"`
	Tc       bool `json:"TC"`
	Rd       bool `json:"RD"`
	Ra       bool `json:"RA"`
	Ad       bool `json:"AD"`
	Cd       bool `json:"CD"`
	Question []struct {
		Name string `json:"name"`
		Type int    `json:"type"`
	} `json:"Question"`
	Answer []struct {
		Name string `json:"name"`
		Type int    `json:"type"`
		TTL  int    `json:"TTL"`
		Data string `json:"data"`
	} `json:"Answer"`
}

func DNSQueryCFDOH(c *gin.Context) {
	host := c.Param("host")
	answer, err := CloudFlareDoh(host)
	if err != nil {
		ezap.Error(err.Error())
		c.AbortWithStatusJSON(500, answer)
		return
	}
	c.JSON(200, answer)
}

// ref. https://developers.cloudflare.com/1.1.1.1/encryption/dns-over-https/make-api-requests/dns-json/
func CloudFlareDoh(name string) (answer *CFDOHResp, err error) {
	api := CLOUDFLAREDOHAPI + "?name=" + name
	headers := map[string]string{"accept": "application/dns-json"}

	resp, err := httpDo(api, "GET", nil, headers)
	if err != nil && (errors.Is(err, syscall.ECONNRESET) || err.(net.Error).Timeout()) {
		ezap.Info("Connection to https://cloudflare-dns.com/dns-query timed out or certificate validation failed. Try using built-in proxy to access.")
		api = APIPROXY + api
		resp, err = httpDo(api, "GET", nil, headers)
	}
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &answer)
	return
}

func httpDo(url string, method string, data []byte, headers map[string]string) (*http.Response, error) {
	client := &http.Client{Timeout: 5 * time.Second}

	req, err := http.NewRequest(method, url, strings.NewReader(string(data)))
	if err != nil {
		return nil, err
	}
	for key, header := range headers {
		req.Header.Set(key, header)
	}
	return client.Do(req)
}
