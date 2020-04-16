package mocks

import (
	"github.com/denisbakhtin/projectmanager/models"
	"github.com/gin-gonic/gin"
)

//EmailMock is a email service mock
type EmailMock struct {
	ActivationSent        bool
	RegistrationSent      bool
	ResetSent             bool
	ResetConfirmationSent bool
}

func (e *EmailMock) SendUserActivationMessage(c *gin.Context, user *models.User) error {
	e.ActivationSent = true
	return nil
}

func (e *EmailMock) SendUserRegistrationMessage(c *gin.Context, user *models.User) error {
	e.RegistrationSent = true
	return nil
}

func (e *EmailMock) SendPasswordResetMessage(c *gin.Context, user *models.User) error {
	e.ResetSent = true
	return nil
}

func (e *EmailMock) SendPasswordResetConfirmation(c *gin.Context, user *models.User) error {
	e.ResetConfirmationSent = true
	return nil
}
