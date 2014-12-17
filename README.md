
# email

[![Build Status](https://travis-ci.org/zemirco/email.svg)](https://travis-ci.org/zemirco/email)
[![GoDoc](https://godoc.org/github.com/zemirco/email?status.svg)](https://godoc.org/github.com/zemirco/email)

Send emails easily.

## Example

```go
package main

import "github.com/zemirco/email"

func main() {
  mail := email.Mail{
    From:    "the@internet.com",
    To:      "mirco.zeiss@gmail.com",
    Subject: "Hello there",
    HTML:    "<h1>Awesome</h1>",
  }
  err := mail.Send("smtp.amazonaws.com", 587, "user", "password")
  if err != nil {
    panic(err)
  }
}
```

## Test

`go test`

## License

MIT
