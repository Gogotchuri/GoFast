package misc

import (
	"fmt"
	"math/rand"
	"net/smtp"
	"time"

	"github.com/Gogotchuri/GoFast/config"
)

const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
const codeLength = 10

/*randCode Generates a random code of constant length from a set of characters*/
func RandCode() string {
	rand.Seed(time.Now().UTC().UnixNano())
	res := make([]byte, codeLength)
	for i := range res {
		res[i] = letters[rand.Intn(len(letters))]
	}
	return string(res)
}

func SendMail(recipient, body, subject string) string {
	conf := config.GetInstance()

	mime := "MIME-version: 1.0;\r\nContent-Type: text/html;\r\n charset=\"UTF-8\";"
	msg := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n%s\r\n\n%s", conf.ServMail, recipient, subject, mime, body)
	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", conf.ServMail, conf.ServPass, "smtp.gmail.com"),
		conf.ServMail, []string{recipient}, []byte(msg))
	if err != nil {
		return err.Error()
	}

	// conf := config.GetInstance()
	// msg = "To: " + recipient + "\r\n" +
	// 	"Subject: " + subject + "\r\n" +
	// 	"\r\n" + msg + "\r\n"
	// // Set up authentication information.
	// auth := smtp.PlainAuth(
	// 	"",
	// 	conf.ServMail,
	// 	conf.ServPass,
	// 	"smtp.gmail.com",
	// )
	// // Send the email.
	// err := smtp.SendMail(
	// 	"smtp.gmail.com:587",
	// 	auth,
	// 	conf.ServMail,
	// 	[]string{recipient},
	// 	[]byte(msg),
	// )
	// if err != nil {
	// 	return err.Error()
	// }
	return ""
}
