# Logrus Filtered Formatter
Use this formatter to log json with filtered fields

## Usage
```go
package main

import (
	"github.com/Fs02/logrus-filtered-formatter"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	fields := []string{
		"password",
		"email",
	}
	log.Formatter = filtered.New(fields, &logrus.JSONFormatter{})

	log.WithFields(logrus.Fields{
		"_password_": "asdfasdf",
		"email_field":    "johndoe@gmail.com",
		"name":     "john doe",
	}).Info("new user created")
}
```
Output:
```json
{"email_field":"[FILTERED]","level":"info","msg":"new user created","name":"john doe","_password_":"[FILTERED]","time":"2017-10-05T16:05:29+07:00"}
```

Formatter filtered fields and log entries name can be in any case and contain whitespaces. 
Formatter will normalize both and check so if you have fields in logs like `email_address`, `old email`, 
`customer_email`, just add `email` field to formatter and it will filter all logs which contain emails.
