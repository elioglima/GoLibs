package GoLibs

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/mail"
	"net/smtp"
)

type SendSMTPMailST struct {
	From        mail.Address
	To          mail.Address
	Subj        string
	Body        string
	SMTP_Server string
	SMTP_Mail   string
	SMTP_Senha  string
}

func SendSMTPMail(p SendSMTPMailST) {

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = p.From.String()
	headers["To"] = p.To.String()
	headers["Subject"] = p.Subj

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	message += "\r\n" + p.Body

	message += "Content-Type: text/html; charset=\"utf-8\"\r\n"
	message += "Content-Transfer-Encoding: 7bit\r\n"
	message += fmt.Sprintf("\r\n%s", p.Body+"\r\n")

	// Connect to the SMTP Server
	// servername := "smtp.perfectvision.kinghost.net:587"
	// "atendimento@perfectvision.kinghost.net"
	host, _, _ := net.SplitHostPort(p.SMTP_Server)
	auth := smtp.PlainAuth("", p.SMTP_Mail, p.SMTP_Senha, host)

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	c, err := smtp.Dial(p.SMTP_Server)
	if err != nil {
		log.Panic(err)
	}

	c.StartTLS(tlsconfig)

	// Auth
	if err = c.Auth(auth); err != nil {
		log.Panic(err)
	}

	// To && From
	if err = c.Mail(p.From.Address); err != nil {
		log.Panic(err)
	}

	if err = c.Rcpt(p.To.Address); err != nil {
		log.Panic(err)
	}

	// Data
	w, err := c.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	c.Quit()

}
