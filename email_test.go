package email

import (
	"testing"
)

const host = "mailtrap.io"
const port = 465
const user = "2771312eff9263ac4"
const pass = "d7f3fa84380b13"

func TestInvalidFrom(t *testing.T) {
	mail := Mail{
		From:    "internet.com",
		To:      "mirco.zeiss@gmail.com",
		Subject: "Hello there",
		HTML:    "<h1>Awesome</h1>",
	}
	err := mail.Send(host, port, user, pass)
	if err == nil {
		t.Error("from address not checked for valid format")
	}
}

func TestInvalidTo(t *testing.T) {
	mail := Mail{
		From:    "the@internet.com",
		To:      "mirco.zeiss",
		Subject: "Hello there",
		HTML:    "<h1>Awesome</h1>",
	}
	err := mail.Send(host, port, user, pass)
	if err == nil {
		t.Error("to address not checked for valid format")
	}
}

func TestSend(t *testing.T) {
	mail := Mail{
		From:    "the@internet.com",
		To:      "mirco.zeiss@gmail.com",
		Subject: "Hello there",
		HTML:    "<h1>Awesome</h1>",
	}
	err := mail.Send(host, port, user, pass)
	if err != nil {
		t.Fatal(err)
	}
}
