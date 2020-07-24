package hash

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"

	"github.com/Gogotchuri/GoFast/config"
)

/*GetPasswordHash Returns hashed version of the passed password*/
func GetPasswordHash(password string) string {
	// Create a new hmac
	mac := hmac.New(sha256.New, []byte(config.GetInstance().HMACKey))
	// Write our data with added pepper to it
	mac.Write([]byte(password + config.GetInstance().Pepper))

	return base64.URLEncoding.EncodeToString(mac.Sum(nil))
}

/*IsPasswordValid Returns whether passed password's hash is equal to passed hash*/
func IsPasswordValid(password, hash string) bool {
	passwordHash := GetPasswordHash(password)
	return passwordHash == hash
}
