package models

import "errors"

var (
	ErrorUserDoesntExist = errors.New("user doesnt exist")
	ErrorSendMail = errors.New("sorry, we cant change your password now, try it later please")
	ErrorNoRows = errors.New("")
)
