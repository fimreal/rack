package email

import (
	"net/http"

	"github.com/fimreal/goutils/ezap"
	"github.com/gin-gonic/gin"
)

type MailerInterface interface {
	SendMail(c *gin.Context)
}

func SendMail(c *gin.Context) {
	var letter Letter
	if err := c.ShouldBindJSON(&letter); err != nil {
		ezap.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ezap.Debugf("请求发送邮件, 传入收件人: %v, 发送标题: %s, 内容: %s", letter.Mailto, letter.Subject, letter.Body)

	if err := Mailto(&letter); err != nil {
		ezap.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ezap.Debug("邮件发送成功")
	c.JSON(http.StatusOK, gin.H{"result": "邮件发送成功"})
}

// // SendMail 用来配置邮件收件人、标题、内容, 返回 http.Response 格式
// func SendMail(w http.ResponseWriter, r *http.Request) {

// 	// 判断如果不是 POST 则报错
// 	if r.Method != "POST" {
// 		fmt.Fprintf(w, "Sorry, only POST methods are supported.")
// 		return
// 	}

// 	// 解析url传递的参数，对于POST则解析响应包的主体（request body）
// 	if err := r.ParseForm(); err != nil {
// 		fmt.Fprintf(w, "ParseForm() err: %v", err)
// 		log.Println("ParseForm() err: ", err)
// 		return
// 	}

// 	// 收件人
// 	mailTo := []string{r.FormValue("mailto")}

// 	// 邮件标题
// 	subject := r.FormValue("subject")

// 	// 邮件内容
// 	body := r.FormValue("body")

// 	err := Mailsrv(mailTo, subject, body)
// 	if err != nil {
// 		// 日志打印失败信息
// 		log.Println(err, "\n    Details: {\"mailTo\": \"", mailTo, "\", \"subject\": \"", subject, "\"}")
// 		// 返回失败信息
// 		fmt.Fprintln(w, "Send fail!")
// 		return
// 	}
// 	// 日志打印成功信息
// 	log.Println("Send successfully!\n    Details: {\"mailTo\": \"", mailTo, "\", \"subject\": \"", subject, "\"}")
// 	// 返回成功信息
// 	fmt.Fprintln(w, "Send successfully!")
// }
