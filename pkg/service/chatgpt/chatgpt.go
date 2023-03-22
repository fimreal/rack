package chatgpt

import (
	"context"
	"net/http"
	"net/url"
	"strings"

	"github.com/fimreal/goutils/ezap"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
	"github.com/spf13/viper"
)

func Ask(c *gin.Context) {

	var ask NewASk
	if err := c.ShouldBind(&ask); err != nil {
		ezap.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ezap.Debugf("ask: %s", ask.ASK)
	gptsay, err := hiGPT(ask.ASK)
	if err != nil {
		ezap.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	gptsay = strings.Replace(gptsay, "\n", "", 2)

	if ask.H {
		c.Data(http.StatusOK, "text/html; charset=utf-8", md2html(gptsay))
		return
	}
	c.String(http.StatusOK, gptsay)
}

func hiGPT(askStr string) (string, error) {
	openaiToken := viper.GetString("chatgpt.token")
	customApiUrl := viper.GetString("chatgpt.api")
	customProxyUrl := viper.GetString("chatgpt.proxyurl")

	// 创建默认配置
	config := openai.DefaultConfig(openaiToken)

	// 修改 API 地址
	apiUrl, err := url.Parse(customApiUrl)
	if err != nil {
		return "", err
	}
	baseUrl, _ := url.Parse(config.BaseURL)
	config.BaseURL = strings.Replace(config.BaseURL, baseUrl.Host, apiUrl.Host, 1)
	ezap.Debug("chatgpt API: " + config.BaseURL)

	// 添加代理地址
	if customProxyUrl != "" {
		proxyUrl, err := url.Parse(customProxyUrl)
		if err != nil {
			return "", err
		}
		config.HTTPClient = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyUrl),
			},
		}
		ezap.Debug("chatgpt API with proxy: " + customProxyUrl)
	}

	// 创建连接配置
	c := openai.NewClientWithConfig(config)

	req := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role: openai.ChatMessageRoleUser,
				// 这里输入对话内容
				Content: askStr,
			},
		}}

	resp, err := c.CreateChatCompletion(
		context.Background(),
		req,
	)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
