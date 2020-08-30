package utils

import (
	"github.com/go-playground/validator/v10"
	"log"
	"regexp"
)

func UsernameValidator(fl validator.FieldLevel) bool {
	if username, ok := fl.Field().Interface().(string); ok {
		flag, err := regexp.Match("^[a-zA-Z]\\w{3,15}$", []byte(username))
		if err != nil {
			log.Println(err.Error())
			return false
		}
		return flag
	}
	return false
}

func TelValidator(fl validator.FieldLevel) bool {
	if tel, ok := fl.Field().Interface().(string); ok {
		flag, err := regexp.Match(
			"^(?:\\+?86)?1(?:3\\d{3}|5[^4\\D]\\d{2}|8\\d{3}|7(?:[35678]\\d{2}|4(?:0\\d|1[0-2]|9\\d))|9[189]\\d{2}|66\\d{2})\\d{6}$",
			[]byte(tel))
		if err != nil {
			log.Println(err.Error())
			return false
		}
		return flag
	}
	return false
}