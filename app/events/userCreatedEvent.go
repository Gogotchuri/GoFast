package events

import (
	"github.com/Gogotchuri/GoFast/app/models"
	"github.com/Gogotchuri/GoFast/app/services"
)

func FireUserCreated(user *models.User) {
	services.SendVerificationMail(user)
}
