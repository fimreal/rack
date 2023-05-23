package crtag

import (
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"syscall"
	"time"

	"github.com/fimreal/goutils/ezap"
)

const (
	dockerhub = "https://hub.docker.com"
	pageSize  = 100 // min 10, max 100, default 1
	page      = 1   // default 1
	APIPROXY  = "http://cfdown.2fw.top/"
)

// https://hub.docker.com/v2/namespaces/{namespace}/repositories/{repository}/tags
func listRepoTags(namespace, repository, rawQuery string) (httpcode int, body []byte, err error) {
	url := fmt.Sprintf("%s/v2/namespaces/%s/repositories/%s/tags?%s", dockerhub, namespace, repository, rawQuery)
	ezap.Debug("request " + url)

	resp, err := httpDo(url, "GET", nil, nil)
	if errors.Is(err, syscall.ECONNRESET) || err.(net.Error).Timeout() {
		ezap.Info("Connection to " + dockerhub + " timed out or certificate validation failed. Try using built-in proxy to access.")
		url = APIPROXY + url
		ezap.Debug("re request " + url)
		resp, err = httpDo(url, "GET", nil, nil)
	}
	if err != nil {
		return
	}
	httpcode = resp.StatusCode
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)
	return
}

func httpDo(url string, method string, data []byte, headers map[string]string) (*http.Response, error) {
	client := &http.Client{Timeout: 3 * time.Second}

	req, err := http.NewRequest(method, url, strings.NewReader(string(data)))
	if err != nil {
		return nil, err
	}
	for key, header := range headers {
		req.Header.Set(key, header)
	}
	return client.Do(req)
}
