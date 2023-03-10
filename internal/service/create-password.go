package service

import "math/rand"

const voc string = "abcdfghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numbers string = "0123456789"
const symbols string = "?/!)(*+-"

func (se *Service) CreatePassword() string {
	chars := voc
	chars = chars + numbers
	chars = chars + symbols
	return se.GeneratePassword(14, chars)
}

func (se *Service) GeneratePassword(length int, chars string) string {
	password := ""
	for i := 0; i < length; i++ {
		password += string([]rune(chars)[rand.Intn(len(chars))])
	}
	return password
}
