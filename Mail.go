package GoLibs

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/mail"
	"net/smtp"
	"strconv"
)

type SendSMTPMailDadosST struct {
	From        mail.Address
	To          mail.Address
	Subj        string
	Body        string
	SMTP_Server string
	SMTP_Mail   string
	SMTP_Senha  string
}

type SendSMTPMailServerST struct {
	Server string
	Mail   string
	Port   int
	Senha  string
}

type SendSMTPMailST struct {
	SMTP       SendSMTPMailServerST
	Dados      *[]SendSMTPMailDadosST
	smtpClient *smtp.Client
}

func New(smtp SendSMTPMailServerST) (*SendSMTPMailST, error) {
	s := &SendSMTPMailST{}
	s.SMTP = smtp
	err := s.Connect()
	return s, err
}

func (s *SendSMTPMailST) Connect() error {

	var err error

	// Connect to the SMTP Server
	// servername := "smtp.perfectvision.kinghost.net:587"
	// "atendimento@perfectvision.kinghost.net"
	host, _, _ := net.SplitHostPort(s.SMTP.Server + ":" + strconv.Itoa(s.SMTP.Port))
	auth := smtp.PlainAuth("", s.SMTP.Mail, s.SMTP.Senha, host)

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	s.smtpClient, err = smtp.Dial(s.SMTP.Server)
	if err != nil {
		return err
	}

	s.smtpClient.StartTLS(tlsconfig)

	// Auth
	if err = s.smtpClient.Auth(auth); err != nil {
		return err
	}

	return nil
}

func (s *SendSMTPMailST) Close() {
	s.smtpClient.Quit()
}

func (s *SendSMTPMailST) Send() error {

	// Setup headers
	for _, dados := range *s.Dados {

		headers := make(map[string]string)
		headers["From"] = dados.From.String()
		headers["To"] = dados.To.String()
		headers["Subject"] = dados.Subj

		// Setup message
		msg := ""
		for k, v := range headers {
			msg += fmt.Sprintf("%s: %s\r\n", k, v)
		}

		msg += "Content-Type: text/html; charset=\"utf-8\"\r\n"
		msg += "Content-Transfer-Encoding: 7bit\r\n"
		msg += fmt.Sprintf("\r\n%s", dados.Body+"\r\n")

		// To && From
		if err := s.smtpClient.Mail(dados.From.Address); err != nil {
			return err
		}

		if err := s.smtpClient.Rcpt(dados.To.Address); err != nil {
			return err
		}

		// Data
		w, err := s.smtpClient.Data()
		if err != nil {
			return err
		}

		defer w.Close()

		_, err = w.Write([]byte(msg))
		if err != nil {
			return err
		}
	}

	return nil
}
