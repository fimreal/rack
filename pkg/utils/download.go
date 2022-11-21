package utils

import (
	"crypto/tls"
	"io"
	"net/http"
	"os"
)

// 创建文件并下载
func DownloadNotls(url string, filename string) error {
	// 在容器内挂载主机的证书文件 /etc/ssl/certs，可不跳过证书验证
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	return err
}
