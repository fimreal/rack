package common

import (
	"net/http"
	"strings"

	"log"
	"os/exec"

	caputils "github.com/fimreal/rack/pkg/cap/utils"
	"github.com/gin-gonic/gin"
)

func Hostname(context *gin.Context) {
	c := exec.Command("hostname")
	cmdResult, err := c.Output()
	if err != nil {
		log.Println(err)
	}
	res := caputils.CommandResult{Command: "/hostname", ReturnCode: 0, Result: strings.Trim(string(cmdResult), "\n")}
	context.JSON(http.StatusOK, res)
}

// 记住要限制容器内权限，同时在 nginx 反代增加简单认证保证安全
func ShellExec(context *gin.Context) {
	res := caputils.CommandResult{}
	pass, ok := context.GetPostForm("pass")
	if pass != "xm" || !ok {
		res = caputils.CommandResult{Command: "/shell", ReturnCode: 1, Result: "出错了！缺少参数，或服务临时不可用。"}
		context.JSON(http.StatusOK, res)
		return
	}
	args, ok := context.GetPostForm("cmd")
	if !ok {
		res = caputils.CommandResult{Command: "/shell", ReturnCode: 1, Result: "出错了！请检查传入参数：" + args}
		context.JSON(http.StatusOK, res)
		return
	}
	c := exec.Command("/bin/sh", "-c", args)
	log.Println("开始执行命令：", c)
	cmdResult, err := c.Output()
	if err != nil {
		log.Println(err)
	}
	res = caputils.CommandResult{Command: "/hostname", ReturnCode: 0, Result: strings.Trim(string(cmdResult), "\n")}
	context.JSON(http.StatusOK, res)
}
