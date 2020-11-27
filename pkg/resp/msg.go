package resp

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func defaultMsg(s string, msg ...string) string {
	message := s
	if len(msg) > 0 {
		message = msg[0]
	}
	return message
}

func defaultParamMsg(s string, msg ...interface{}) string {
	message := s
	if len(msg) > 0 {
		switch msg[0].(type) {
		case string:
			message = msg[0].(string)
		case validator.ValidationErrors:
			message = "invalid " + msg[0].(validator.ValidationErrors)[0].Field()
		default:
			message = fmt.Sprint(msg[0])
		}
	}
	return message
}
