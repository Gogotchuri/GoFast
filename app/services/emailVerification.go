package services

import (
	"strings"
	"time"

	"github.com/Gogotchuri/GoFast/app/models"
	"github.com/Gogotchuri/GoFast/app/services/cache"
	"github.com/Gogotchuri/GoFast/app/services/misc"
	"github.com/Gogotchuri/GoFast/config"
)

/*SendVerificationMail Generates random code and sends to passed user's email*/
func SendVerificationMail(user *models.User) {
	otacConf := config.GetInstance().Redis.OTAC

	code := misc.RandCode()
	cache.GetRedisInstance().Set(otacConf.EntryPrefix+user.Email, code, time.Duration(otacConf.Expires)*time.Minute)

	go misc.SendMail(user.Email, "Your verification code is: "+code, "GoFast email verfication")
}

/*VerifyEmailByCode Verifies user if the passed code is correct*/
func VerifyEmailByCode(user *models.User, code string) bool {
	otacConf := config.GetInstance().Redis.OTAC
	otacKey := otacConf.EntryPrefix + user.Email

	cachedOTAC, err := cache.GetRedisInstance().Get(otacKey).Result()
	if err != nil || strings.ToLower(cachedOTAC) != strings.ToLower(code) {
		return false
	}

	user.EmailVerifiedAt = time.Now()
	user.Save()
	cache.GetRedisInstance().Expire(otacKey, 0)
	return true
}
