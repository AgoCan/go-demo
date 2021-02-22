package main

import (
	"bytes"
	"encoding/base64"
	"flag"
	"io/ioutil"
	"log"
	"net/smtp"
	"strings"
	"time"
)

// 传参信息
var (
	user     = flag.String("u", "hank@hankbook.cn", "email user.")
	passwd   = flag.String("p", "****", "email password.")
	smtpHost = flag.String("h", "smtp.qiye.aliyun.com", "email smtp host.")
	smtpPort = flag.String("port", "25", "email smtp host.")
	to       = flag.String("t", "123456@qq.com,88888@qq.com", "mail to address. use ',' split")
	subject  = flag.String("s", "subject", "mail subject.")
	bodyPath = flag.String("f", "./1.csv", "mail body.")
	message  = flag.String("m", "null", "mail message")
)

// Mail define email interface, and implemented auth and send method
type Mail interface {
	Auth()
	Send(message Message) error
}

// SendMail 定义发送信息
type SendMail struct {
	user     string
	password string
	host     string
	port     string
	auth     smtp.Auth
}

// Attachment 附件
type Attachment struct {
	name        string
	contentType string
	withFile    bool
}

// Message 邮件信息
type Message struct {
	from        string
	to          []string
	cc          []string
	bcc         []string
	subject     string
	body        string
	contentType string
	attachment  Attachment
}

func buildBody() (bodyString string) {
	var temp string
	bodyArray := make([]string, 6)
	bodyArray = append(bodyArray, *message)
	bodyArray = append(bodyArray, "<html><head><meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\" /></head>")
	bodyArray = append(bodyArray, "<table bgcolor='' style=' font-size: 14px;'border='1' cellspacing='0' cellpadding='0' bordercolor='#000000' width='95%' align='center' >")
	body, _ := ioutil.ReadFile(*bodyPath)
	lines := strings.Split(string(body), "\n")

	for _, line := range lines {
		rows := strings.Split(line, ",")
		temp = "<tr>"
		for _, row := range rows {
			temp = temp + "<td>" + row + "</td>"
		}
		temp = temp + "</tr>" + "\n"
		bodyArray = append(bodyArray, temp)
	}
	bodyArray = append(bodyArray, "</table></body></html>")
	bodyString = strings.Join(bodyArray, "")
	return bodyString
}
func main() {
	flag.Parse()
	var mail Mail
	mail = &SendMail{user: *user, password: *passwd, host: *smtpHost, port: *smtpPort}
	message := Message{from: *user,
		to:          strings.Split(*to, ","),
		cc:          []string{},
		bcc:         []string{},
		subject:     *subject,
		body:        buildBody(),
		contentType: "text/html;charset=utf-8",
		attachment: Attachment{
			name:        "test.jpg",
			contentType: "image/jpg",
			withFile:    false,
		},
	}
	mail.Send(message)
}

// Auth 用户认证
func (mail *SendMail) Auth() {
	mail.auth = smtp.PlainAuth("", mail.user, mail.password, mail.host)
}

// Send 发送邮件
func (mail SendMail) Send(message Message) error {
	mail.Auth()
	buffer := bytes.NewBuffer(nil)
	boundary := "GoBoundary"
	Header := make(map[string]string)
	Header["From"] = message.from
	Header["To"] = strings.Join(message.to, ";")
	Header["Cc"] = strings.Join(message.cc, ";")
	Header["Bcc"] = strings.Join(message.bcc, ";")
	Header["Subject"] = message.subject
	Header["Content-Type"] = "multipart/mixed;boundary=" + boundary
	Header["Mime-Version"] = "1.0"
	Header["Date"] = time.Now().String()
	mail.writeHeader(buffer, Header)

	body := "\r\n--" + boundary + "\r\n"
	body += "Content-Type:" + message.contentType + "\r\n"
	body += "\r\n" + message.body + "\r\n"
	buffer.WriteString(body)

	if message.attachment.withFile {
		attachment := "\r\n--" + boundary + "\r\n"
		attachment += "Content-Transfer-Encoding:base64\r\n"
		attachment += "Content-Disposition:attachment\r\n"
		attachment += "Content-Type:" + message.attachment.contentType + ";name=\"" + message.attachment.name + "\"\r\n"
		buffer.WriteString(attachment)
		defer func() {
			if err := recover(); err != nil {
				log.Fatalln(err)
			}
		}()
		mail.writeFile(buffer, message.attachment.name)
	}

	buffer.WriteString("\r\n--" + boundary + "--")
	smtp.SendMail(mail.host+":"+mail.port, mail.auth, message.from, message.to, buffer.Bytes())
	return nil
}

func (mail SendMail) writeHeader(buffer *bytes.Buffer, Header map[string]string) string {
	header := ""
	for key, value := range Header {
		header += key + ":" + value + "\r\n"
	}
	header += "\r\n"
	buffer.WriteString(header)
	return header
}

// read and write the file to buffer
func (mail SendMail) writeFile(buffer *bytes.Buffer, fileName string) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err.Error())
	}
	payload := make([]byte, base64.StdEncoding.EncodedLen(len(file)))
	base64.StdEncoding.Encode(payload, file)
	buffer.WriteString("\r\n")
	for index, line := 0, len(payload); index < line; index++ {
		buffer.WriteByte(payload[index])
		if (index+1)%76 == 0 {
			buffer.WriteString("\r\n")
		}
	}
}
