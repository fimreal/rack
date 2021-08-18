package common

import (
	"net/http"
	"strings"

	"log"
	"os/exec"

	"github.com/fimreal/go-tools/cmd"
	"github.com/gin-gonic/gin"
)

// healthcheck
func Healthcheck(context *gin.Context) {
	context.String(http.StatusOK, "ok")
}

func Hostname(context *gin.Context) {
	c := exec.Command("hostname")
	cmdResult, err := c.Output()
	if err != nil {
		log.Println(err)
	}
	res := cmd.CommandResult{Command: "/hostname", ReturnCode: 0, Result: strings.Trim(string(cmdResult), "\n")}
	context.JSON(http.StatusOK, res)
}

// 记住要限制容器内权限，同时在 nginx 反代增加简单认证保证安全
func ShellExec(context *gin.Context) {
	res := cmd.CommandResult{}
	pass, ok := context.GetPostForm("pass")
	if pass != "xm" || !ok {
		res = cmd.CommandResult{Command: "/shell", ReturnCode: 1, Result: "出错了！缺少参数，或服务临时不可用。"}
		context.JSON(http.StatusOK, res)
		return
	}
	args, ok := context.GetPostForm("cmd")
	if !ok {
		res = cmd.CommandResult{Command: "/shell", ReturnCode: 1, Result: "出错了！请检查传入参数：" + args}
		context.JSON(http.StatusOK, res)
		return
	}
	c := exec.Command("/bin/sh", "-c", args)
	log.Println("开始执行命令：", c)
	cmdResult, err := c.Output()
	if err != nil {
		log.Println(err)
	}
	res = cmd.CommandResult{Command: "/hostname", ReturnCode: 0, Result: strings.Trim(string(cmdResult), "\n")}
	context.JSON(http.StatusOK, res)
}
