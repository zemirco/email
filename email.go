// Package email makes email sending simple.
//
//  mail := Mail{
//    From:    "mirco.zeiss@gmail.com",
//    To:      "john@wayne.com",
//    Subject: "Hey John",
//    HTML:    "<h1>Good job!</h1>",
//  }
//  err := mail.Send("smtp.amazonaws.com", 587, "user", "password")
//  if err != nil {
//    panic(err)
//  }
package email

import (
	"bytes"
	"fmt"
	"net/mail"
	"net/smtp"
	"net/textproto"
)

type Mail struct {
	From    string
	To      string
	Subject string
	HTML    string
}

// Send does what it is supposed to do.
func (m *Mail) Send(host string, port int, user, pass string) error {

	// validate from address
	from, err := mail.ParseAddress(m.From)
	if err != nil {
		return err
	}

	// validate to address
	to, err := mail.ParseAddress(m.To)
	if err != nil {
		return err
	}

	// set headers for html email
	header := textproto.MIMEHeader{}
	header.Set(textproto.CanonicalMIMEHeaderKey("from"), from.Address)
	header.Set(textproto.CanonicalMIMEHeaderKey("to"), to.Address)
	header.Set(textproto.CanonicalMIMEHeaderKey("content-type"), "text/html; charset=UTF-8")
	header.Set(textproto.CanonicalMIMEHeaderKey("mime-version"), "1.0")
	header.Set(textproto.CanonicalMIMEHeaderKey("subject"), m.Subject)

	// init empty message
	var buffer bytes.Buffer

	// write header
	for key, value := range header {
		buffer.WriteString(fmt.Sprintf("%s: %s\r\n", key, value[0]))
	}

	// write body
	buffer.WriteString(fmt.Sprintf("\r\n%s", m.HTML))

	// send email
	addr := fmt.Sprintf("%s:%d", host, port)
	auth := smtp.PlainAuth("", user, pass, host)
	return smtp.SendMail(addr, auth, from.Address, []string{to.Address}, buffer.Bytes())
}
