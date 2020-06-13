package misc

import (
	"math/rand"
	"net/smtp"
	"time"

	"github.com/Gogotchuri/GoFast/config"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
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

func SendMail(recipient, msg, subject string) string {
	msg = "To: " + recipient + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" + msg + "\r\n"
	// Set up authentication information.
	auth := smtp.PlainAuth(
		"",
		config.GetInstance().Serv_mail,
		config.GetInstance().Serv_pass,
		"smtp.gmail.com",
	)
	// Send the email.
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		config.GetInstance().Serv_mail,
		[]string{recipient},
		[]byte(msg),
	)
	if err != nil {
		return err.Error()
	}
	return ""
}
