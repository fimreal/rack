package common

import (
	"net/http"
	"strings"

	"os/exec"

	"github.com/fimreal/goutils/ezap"
	caputils "github.com/fimreal/rack/pkg/cap/utils"
	"github.com/gin-gonic/gin"
	"github.com/ory/viper"
)

func Hostname(context *gin.Context) {
	c := exec.Command("hostname")
	cmdResult, err := c.Output()
	if err != nil {
		ezap.Error(err)
	}
	res := caputils.CommandResult{Command: "/hostname", ReturnCode: 0, Result: strings.Trim(string(cmdResult), "\n")}
	context.JSON(http.StatusOK, res)
}

// 记住要限制容器内权限，同时在 nginx 反代增加简单认证保证安全
func ShellExec(context *gin.Context) {
	res := caputils.CommandResult{}
	pass, ok := context.GetPostForm("pass")
	if pass != viper.GetString("tools.password") || !ok {
		res = caputils.CommandResult{Command: "/shell", ReturnCode: 1, Result: "出错了！缺少参数，或服务临时不可用。"}
		ezap.Error("req pass: ", pass)
		context.JSON(http.StatusOK, res)
		return
	}
	args, ok := context.GetPostForm("cmd")
	if !ok {
		res = caputils.CommandResult{Command: "/shell", ReturnCode: 1, Result: "出错了！请检查传入参数：" + args}
		ezap.Error("req pass: ", pass, ", cmd: ", args)
		context.JSON(http.StatusOK, res)
		return
	}
	c := exec.Command("/bin/sh", "-c", args)
	ezap.Debug("开始执行命令：", c)
	cmdResult, err := c.Output()
	if err != nil {
		ezap.Error(err)
	}
	res = caputils.CommandResult{Command: "/hostname", ReturnCode: 0, Result: strings.Trim(string(cmdResult), "\n")}
	context.JSON(http.StatusOK, res)
}
