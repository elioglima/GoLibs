package GoLibs

import (
	"crypto/tls"
	"fmt"
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

func SendSMTPMail(p SendSMTPMailST) (code int, message string, err error) {

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = p.From.String()
	headers["To"] = p.To.String()
	headers["Subject"] = p.Subj

	// Setup message
	msg := ""
	for k, v := range headers {
		msg += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	// message += "\r\n" + p.Body

	msg += "Content-Type: text/html; charset=\"utf-8\"\r\n"
	msg += "Content-Transfer-Encoding: 7bit\r\n"
	msg += fmt.Sprintf("\r\n%s", p.Body+"\r\n")

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
		return 0, "", err
	}

	c.StartTLS(tlsconfig)

	// Auth
	if err = c.Auth(auth); err != nil {
		return 0, "", err
	}

	// To && From
	if err = c.Mail(p.From.Address); err != nil {
		return 0, "", err
	}

	if err = c.Rcpt(p.To.Address); err != nil {
		return 0, "", err
	}

	defer c.Quit()
	// Data
	w, err := c.Data()
	if err != nil {
		return 0, "", err
	}
	defer w.Close()

	_, err = w.Write([]byte(msg))
	if err != nil {
		return 0, "", err
	}

	code, message, err = c.Text.ReadResponse(0)
	return

}
