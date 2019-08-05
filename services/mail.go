package services

import (
	"bytes"
	"fmt"
	"log"
	"path"

	"html/template"

	"github.com/denisbakhtin/projectmanager/config"
	"github.com/denisbakhtin/projectmanager/models"
	"github.com/gin-gonic/gin"
	gomail "gopkg.in/gomail.v1"
)

//SendUserActivationMessage sends email message with account activation instructions, refer to config.yml for mail settings
func SendUserActivationMessage(c *gin.Context, user *models.User) error {
	scheme := c.Request.URL.Scheme
	if len(scheme) == 0 {
		scheme = "http"
	}
	link := fmt.Sprintf("%s://%s/#!/activate/%s", scheme, c.Request.Host, user.Token)
	msg := gomail.NewMessage()
	msg.SetHeader("From", config.Settings.SMTPReply)
	msg.SetHeader("To", user.Email)
	msg.SetHeader("Subject", fmt.Sprintf("Your account on %s requires activation", config.Settings.ProjectName))
	tmpl, err := template.New("").ParseFiles(path.Join(config.AppDir, "views", "email", "activation.tmpl"))
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buf, "email/activation.tmpl", gin.H{"Name": user.Name, "ProjectName": config.Settings.ProjectName, "Link": link}); err != nil {
		return err
	}
	msg.SetBody("text/html", buf.String())

	mailer := gomail.NewMailer(config.Settings.SMTPServer, config.Settings.SMTPLogin, config.Settings.SMTPPassword, config.Settings.SMTPPort)
	if err := mailer.Send(msg); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

//SendUserRegistrationMessage sends email message with account registration notice, refer to config.yml for mail settings
func SendUserRegistrationMessage(c *gin.Context, user *models.User) error {
	scheme := c.Request.URL.Scheme
	if len(scheme) == 0 {
		scheme = "http"
	}
	msg := gomail.NewMessage()
	msg.SetHeader("From", config.Settings.SMTPReply)
	msg.SetHeader("To", user.Email)
	msg.SetHeader("Subject", fmt.Sprintf("You have been successfully registered on %s", config.Settings.ProjectName))
	tmpl, err := template.New("").ParseFiles(path.Join(config.AppDir, "views", "email", "registration.tmpl"))
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buf, "email/registration.tmpl", gin.H{"Name": user.Name, "ProjectName": config.Settings.ProjectName}); err != nil {
		return err
	}
	msg.SetBody("text/html", buf.String())

	mailer := gomail.NewMailer(config.Settings.SMTPServer, config.Settings.SMTPLogin, config.Settings.SMTPPassword, config.Settings.SMTPPort)
	if err := mailer.Send(msg); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

//SendPasswordResetMessage sends email message with a password reset link, refer to config.yml for mail settings
func SendPasswordResetMessage(c *gin.Context, user *models.User) error {
	scheme := c.Request.URL.Scheme
	if len(scheme) == 0 {
		scheme = "http"
	}
	link := fmt.Sprintf("%s://%s/#!/reset/%s", scheme, c.Request.Host, user.Token)
	msg := gomail.NewMessage()
	msg.SetHeader("From", config.Settings.SMTPReply)
	msg.SetHeader("To", user.Email)
	msg.SetHeader("Subject", fmt.Sprintf("Password reset instructions on %s", config.Settings.ProjectName))
	tmpl, err := template.New("").ParseFiles(path.Join(config.AppDir, "views", "email", "reset.tmpl"))
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buf, "email/reset.tmpl", gin.H{"Name": user.Name, "ProjectName": config.Settings.ProjectName, "Link": link}); err != nil {
		return err
	}
	msg.SetBody("text/html", buf.String())

	mailer := gomail.NewMailer(config.Settings.SMTPServer, config.Settings.SMTPLogin, config.Settings.SMTPPassword, config.Settings.SMTPPort)
	if err := mailer.Send(msg); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

//SendPasswordResetConfirmation sends email message with a password reset confirmation
func SendPasswordResetConfirmation(c *gin.Context, user *models.User) error {
	scheme := c.Request.URL.Scheme
	if len(scheme) == 0 {
		scheme = "http"
	}
	msg := gomail.NewMessage()
	msg.SetHeader("From", config.Settings.SMTPReply)
	msg.SetHeader("To", user.Email)
	msg.SetHeader("Subject", fmt.Sprintf("Password reset confirmation on %s", config.Settings.ProjectName))
	tmpl, err := template.New("").ParseFiles(path.Join(config.AppDir, "views", "email", "reset_confirmation.tmpl"))
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buf, "email/reset_confirmation.tmpl", gin.H{"Name": user.Name, "ProjectName": config.Settings.ProjectName}); err != nil {
		return err
	}
	msg.SetBody("text/html", buf.String())

	mailer := gomail.NewMailer(config.Settings.SMTPServer, config.Settings.SMTPLogin, config.Settings.SMTPPassword, config.Settings.SMTPPort)
	if err := mailer.Send(msg); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
