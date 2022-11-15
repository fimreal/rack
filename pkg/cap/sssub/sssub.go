// 可能仅适用于觅云的订阅
package sssub

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SssubToUrl(context *gin.Context) {
	// res := cmd.CommandResult{}
	ssUrls := sssub()
	// res = cmd.CommandResult{Command: "/serverip", ReturnCode: 0, Result: ssUrls}
	// context.JSON(http.StatusOK, res)
	var res string
	for _, v := range ssUrls {
		res += v + "\n"
	}
	context.String(http.StatusOK, res)
}

func sssub() (ssUrls []string) {
	// 发起请求
	// url := os.Args[1]
	url, ok := os.LookupEnv("SSSUBURL")
	if !ok {
		return nil
	}
	resp, err := http.Get(url)

	// 如果重试失败，则加 http scheme 重试一次。
	if err != nil {
		url = "http://" + os.Args[1]
		resp, err = http.Get(url)
	}
	if err != nil {
		log.Fatalf("Err: 获取%s出错, %s", url, err)
	}
	defer resp.Body.Close()

	ssSub, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("解析 http body 出错: %s", err)
	}

	ssUrls = jsonToUrl(ssSub)
	// for _, v := range ss_urls {
	// 	fmt.Println(v)
	// }
	return
}

func jsonToUrl(ssSub []byte) (ss_urls []string) {
	type Servers struct {
		Server     string `json:"server"`
		Password   string `json:"password"`
		Method     string `json:"method"`
		ServerPort int    `json:"server_port"`
		ID         string `json:"id"`
		Remarks    string `json:"remarks"`
	}
	type ssSubJson struct {
		Servers []Servers `json:"servers"`
		Version int       `json:"version"`
	}
	var j ssSubJson

	err := json.Unmarshal(ssSub, &j)
	if err != nil {
		log.Fatalf("解析 json 出错: %s", err)
	}

	for _, v := range j.Servers {
		// 将加密方法和密码 base64 处理
		vbyte := []byte(v.Method + ":" + v.Password)
		str64 := base64.StdEncoding.EncodeToString(vbyte)

		// 处理 ss 格式的 url。其中需要将注释通过 url decoe 。注释掉的标准方法转化空格会变成“+”，参考 hack 办法处理。
		// ss_url += "ss://" + str64 + "@" + v.Server + ":" + strconv.Itoa(v.ServerPort) + "/?#" + url.QueryEscape(v.Remarks)
		remark := &url.URL{Path: v.Remarks}
		remarkStr := remark.String()
		ss_url := "ss://" + str64 + "@" + v.Server + ":" + strconv.Itoa(v.ServerPort) + "/?#" + remarkStr

		// 存入切片。扩展：根据 Server 或者 Remark 可以区分不同节点
		ss_urls = append(ss_urls, ss_url)
	}

	return ss_urls
}
