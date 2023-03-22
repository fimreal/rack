package chatgpt

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/fimreal/goutils/ezap"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	// openapi = "https://api.openai.com/v1/completions"
	openapi = viper.GetString("chatgpt.api")
	headers = map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "",
	}
)

func Ask(c *gin.Context) {
	headers["Authorization"] = "Bearer " + viper.GetString("chatgpt.token")
	ezap.Debug(headers["Authorization"])

	var ask NewASk
	if err := c.ShouldBind(&ask); err != nil {
		ezap.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ezap.Debugf("ask: %s", ask.ASK)
	gptsay, err := hiGPT(ask.ASK)
	if err != nil {
		ezap.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	gptsayString := gptsay.Choices[0].Text
	gptsayString = strings.Replace(gptsayString, "\n", "", 2)

	if ask.H {
		c.Data(http.StatusOK, "text/html; charset=utf-8", md2html(gptsayString))
		return
	}
	c.String(http.StatusOK, gptsayString)
}

func hiGPT(askStr string) (*ChatGPTSay, error) {

	postDataStruct := &AskChatGPT{
		Model:            "text-davinci-003",
		Prompt:           askStr,
		Temperature:      0.7,
		MaxTokens:        1024,
		TopP:             1,
		FrequencyPenalty: 0,
		PresencePenalty:  0,
	}
	postDataByte, err := json.Marshal(postDataStruct)
	if err != nil {
		return nil, err
	}
	ezap.Debug("req: ", string(postDataByte))
	resp, err := HttpPost(openapi, string(postDataByte), headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	gptsay := &ChatGPTSay{}
	err = json.Unmarshal(body, &gptsay)
	if err != nil {
		return nil, err
	}

	ezap.Debug("GPT say: ", string(body))
	return gptsay, err
}
