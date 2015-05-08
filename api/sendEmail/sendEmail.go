package sendEmail

import (
	"encoding/base64"
	"fmt"
	"github.com/go-martini/martini"
	"net/mail"
	"net/smtp"
	"strings"
)

func SendEmailApiRoute(m martini.Router) {
	m.Get("/", SendEmailTwo)
}

func SendEmailTest() {
	email := "jigenlong@163.com"
	password := "805717"
	host := "smtp.163.com:25"
	toEmail := "317699326@qq.com"
	subject := "Test send email by golang"
	body := `<html>
	         <body>
	         <h3>
	         "Test send email by golang"
	         </h3>
	         </body>
	         </html>
	      `
	emailType := "html"
	fmt.Println("send email")
	err := SendEmail(email, password, host, toEmail, subject, body, emailType)
	if err != nil {
		fmt.Println("send email error!")
		fmt.Println(err)
	} else {
		fmt.Println("send email success!")
	}
}

func SendEmail(email, password, host, toEmail, subject, body, emailType string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", email, password, hp[0])
	var contentType string
	if emailType == "html" {
		contentType = "Content-Type:text/" + emailType + ";charset=UTF-8"
	} else {
		contentType = "Content-Type:text/plain" + ";charset=UTF-8"
	}
	msg := []byte("To:" + toEmail + "\r\n From:" + email + "<" + email + ">\r\n Subject:" + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	sendTo := strings.Split(toEmail, ";")
	err := smtp.SendMail(host, auth, email, sendTo, msg)
	return err
}

func SendEmailTwo() {
	b64 := base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")
	host := "smtp.163.com"
	email := "jigenlong@163.com"
	password := "805717"
	toEmail := "317699326@qq.com"
	from := mail.Address{"发送人", email}
	to := mail.Address{"接收人", toEmail}
	header := make(map[string]string)
	header["From"] = from.String()
	header["To"] = to.String()
	header["Subject"] = fmt.Sprintf("=?UTF-8?B?%s?=", b64.EncodeToString([]byte("修改密码")))
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/html; charset=UTF-8"
	header["Content-Transfer-Encoding"] = "base64"
	body := `<!DOCTYPE html>
				<html>
				<head lang="en">
				    <meta charset="UTF-8">
				    <title></title>
				</head>
				<body>
				<div style="color: aqua">golang send email </div>
				<div><a href="http://192.168.1.6/">wallwa.com</a></div>
				</body>
				</html>
	       `
	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + b64.EncodeToString([]byte(body))
	auth := smtp.PlainAuth(
		"",
		email,
		password,
		host,
	)
	fmt.Println(auth)
	fmt.Println("85")
	err := smtp.SendMail(
		host+":25",
		auth,
		email,
		[]string{to.Address},
		[]byte(message),
	)
	if err != nil {
		panic(err)
	}
}
