package service

import (
	"bytes"
	"context"
	"html/template"

	"github.com/Egor-Tihonov/SandMailLogic/internal/models"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	gomail "gopkg.in/gomail.v2"
)

func (se *Service) SendEmail(ctx context.Context, email string) error {
	name, err := se.repo.GetEmailFromDB(ctx, email)
	if err != nil {
		return err
	}

	newPassword := se.CreatePassword()
	hashPassword, err := se.hashPassword(newPassword)
	if err != nil {
		return err
	}
	err = se.repo.UpdatePassword(ctx, email, hashPassword)
	if err != nil {
		return err
	}
	user := &models.UserMail{
		Email:    email,
		Name:     name,
		Password: newPassword,
	}

	err = se.CreateMail(user)
	if err != nil {
		return err
	}
	return nil
}

// hashPassword ...
func (se *Service) hashPassword(password string) (hashPassword string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return
	}
	hashPassword = string(bytes)
	return
}

func (se *Service) CreateMail(user *models.UserMail) error {
	t := template.New("mail.html")
	t, err := t.ParseFiles("C:\\Обучение\\4 курс\\Диплом\\SandMailLogic\\pkg\\templates\\mail.html")
	if err != nil {
		logrus.Errorf("error parse template: %e", err)
		return models.ErrorSendMail
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, user); err != nil {
		logrus.Errorf("error execute template: %e", err)
		return models.ErrorSendMail
	}

	result := tpl.String()

	m := gomail.NewMessage()
	m.SetHeader("From", "email")
	m.SetHeader("To", user.Email)
	m.SetAddressHeader("Cc", user.Email, user.Name)
	m.SetHeader("Subject", "Reset password at 451 Fahrenheit")
	m.SetBody("text/html", result)
	m.Attach("C:\\Обучение\\4 курс\\Диплом\\SandMailLogic\\pkg\\templates\\mail.html")
	d := gomail.NewDialer("smtp.gmail.com", 587, "email", "password")
	err = d.DialAndSend(m)
	if err != nil {
		logrus.Error("service: failed send email, %e", err)
		return models.ErrorSendMail
	}

	return nil
}
