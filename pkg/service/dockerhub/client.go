package dockerhub

import (
	"fmt"
	"io"
	"net/http"

	"github.com/fimreal/goutils/ezap"
)

const (
	dockerhub = "https://hub.docker.com"
	pageSize  = 100 // min 10, max 100, default 1
	page      = 1   // default 1
)

// https://hub.docker.com/v2/namespaces/{namespace}/repositories/{repository}/tags
func listRepoTags(namespace, repository, rawQuery string) (httpcode int, body []byte, err error) {
	url := fmt.Sprintf("%s/v2/namespaces/%s/repositories/%s/tags?%s", dockerhub, namespace, repository, rawQuery)
	ezap.Debug("req " + url)

	resp, err := http.Get(url)
	httpcode = resp.StatusCode
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)
	return
}
