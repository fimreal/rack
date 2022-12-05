package chatgpt

import (
	"net/http"
	"strings"
)

func HttpPost(url string, data string, headers map[string]string) (*http.Response, error) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil {
		return nil, err
	}
	for key, header := range headers {
		req.Header.Set(key, header)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, err
}
