package services

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/Gogotchuri/GoFast/app/models"
	"github.com/Gogotchuri/GoFast/app/services/cache"
	"github.com/Gogotchuri/GoFast/app/services/hash"
	"github.com/Gogotchuri/GoFast/app/services/misc"
	"github.com/Gogotchuri/GoFast/config"
	"github.com/twinj/uuid"
)

/*SendPasswordResetMail Sends link for reseting password to user's mail*/
func SendPasswordResetMail(user *models.User) {
	resetURL := config.GetInstance().Domain + "/password-reset?"
	params := url.Values{}
	token := uuid.NewV4().String()
	params.Add("token", token)
	resetURL += params.Encode()
	cache.GetRedisInstance().Set(token, strconv.Itoa(int(user.ID)), time.Minute*10)
	body := "<html><body>This email has been sent in response to password reset request. <br> If you haven't requested it at " +
		"<a href=\"" + config.GetInstance().Domain + "\">GoFast</a> just ignore this mail.<br>" +
		"Otherwise follow <a href=\"" + resetURL + "\">the link</a> to reset your password. <br>" +
		"Or enter url manually " + resetURL + "</body></html>"
	go misc.SendMail(user.Email, body, "GoFast password reset")
}

/*SetNewPassword Sets the new password to the owner of the email*/
func SetNewPassword(token, newPassword string) bool {
	idStr, err := cache.GetRedisInstance().Get(token).Result()
	if err != nil {
		_ = fmt.Errorf(err.Error())
		return false
	}
	userID, _ := strconv.ParseUint(idStr, 10, 64)
	user := models.GetUserByID(uint(userID))

	user.Password = hash.GetPasswordHash(newPassword)
	user.Save()
	cache.GetRedisInstance().Expire(token, 0)
	return true
}
