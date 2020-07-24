package services

import (
	"github.com/Gogotchuri/GoFast/app/services/cache"
	"github.com/Gogotchuri/GoFast/config"

	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
)

/*JWTAccessDetails Defines JWT token details*/
type JWTAccessDetails struct {
	UserID  uint
	Token   string
	UUID    string
	Expires int64
}

/*JWTTokens defines struct for easy refresh and access token detail manipulation*/
type JWTTokens struct {
	Access  JWTAccessDetails
	Refresh JWTAccessDetails
}

/*JWTCreateToken Creates and returns JWT token*/
func JWTCreateToken(userID uint) (*JWTTokens, error) {
	var err error

	cfg := config.GetInstance().JWT
	td := &JWTTokens{}
	//Expiration times and uuids
	td.Access.Expires = time.Now().Add(time.Minute * cfg.AccessExp).Unix()
	td.Access.UUID = uuid.NewV4().String()
	td.Refresh.Expires = time.Now().Add(time.Hour * 24 * cfg.RefreshExp).Unix()
	td.Refresh.UUID = uuid.NewV4().String()

	//Creating Access Token
	aClaims := jwt.MapClaims{}
	aClaims["authorized"] = true
	aClaims["access_uuid"] = td.Access.UUID
	aClaims["user_id"] = userID
	aClaims["exp"] = td.Access.Expires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, aClaims)
	td.Access.Token, err = at.SignedString([]byte(cfg.AccessSecret))
	if err != nil {
		return nil, err
	}

	//Creating Refresh Token
	rClaims := jwt.MapClaims{}
	rClaims["refresh_uuid"] = td.Refresh.UUID
	rClaims["user_id"] = userID
	rClaims["exp"] = td.Refresh.Expires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rClaims)
	td.Refresh.Token, err = rt.SignedString([]byte(cfg.RefreshSecret))
	if err != nil {
		return nil, err
	}
	td.Access.UserID = userID
	td.Refresh.UserID = userID
	return td, nil
}

/*Save saves token into fast storage (Currently Redis)*/
func (jwtTD *JWTAccessDetails) Save() error {
	//Set times to automatically delete tokens after expiration
	at := time.Unix(jwtTD.Expires, 0)
	now := time.Now()

	err := cache.GetRedisInstance().Set(jwtTD.UUID, strconv.Itoa(int(jwtTD.UserID)), at.Sub(now)).Err()
	if err != nil {
		return err
	}
	return nil
}

/*Delete deletes token from fast storage (Currently Redis)*/
func (jwtTD *JWTAccessDetails) Delete() error {
	_, err := cache.GetRedisInstance().Del(jwtTD.UUID).Result()
	if err != nil {
		return err
	}
	return nil
}

/*Save saves tokens into fast storage (Currently Redis)*/
func (jwtTD *JWTTokens) Save() error {
	err := jwtTD.Access.Save()
	if err != nil {
		return err
	}
	return jwtTD.Refresh.Save()
}

/*JWTHasValidRefreshToken checks whether given token is valid*/
func JWTHasValidRefreshToken(jwtToken string) (*JWTAccessDetails, error) {
	return hasValidToken(jwtToken, false)
}

/*JWTHasValidToken checks whether given token is valid*/
func JWTHasValidToken(jwtToken string) (*JWTAccessDetails, error) {
	return hasValidToken(jwtToken, true)
}

/*JWTGetTokenDetails Extracts token details*/
func JWTGetTokenDetails(jwtToken string) (*JWTAccessDetails, error) {
	return getTokenDetails(jwtToken, true)
}

/*Check if JWT token is valid and active given a token and a string*/
func hasValidToken(jwtToken string, isAccess bool) (*JWTAccessDetails, error) {
	secret := config.GetInstance().JWT.AccessSecret
	if !isAccess {
		secret = config.GetInstance().JWT.RefreshSecret
	}
	token, err := verifyToken(jwtToken, secret)
	if err != nil {
		return nil, err
	}
	//Check claims set
	if _, ok := token.Claims.(jwt.Claims); !ok || !token.Valid {
		return nil, fmt.Errorf("claims not set correctly")
	}
	td, err := getTokenDetails(jwtToken, isAccess)
	if err != nil {
		return nil, err
	}
	_, err = checkTokenInStorage(td)
	if err != nil {
		return nil, err
	}

	return td, nil
}

/*getTokenDetails Extracts token details*/
func getTokenDetails(jwtToken string, isAccess bool) (*JWTAccessDetails, error) {
	secret := config.GetInstance().JWT.AccessSecret
	if !isAccess {
		secret = config.GetInstance().JWT.RefreshSecret
	}
	token, err := verifyToken(jwtToken, secret)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("Token not presented with correct claims")
	}
	accessUUID, ok := claims["access_uuid"].(string)
	if !isAccess {
		accessUUID, ok = claims["refresh_uuid"].(string)
	}
	if !ok {
		return nil, fmt.Errorf("Token uuid not found in claims")
	}

	userID, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
	if err != nil {
		return nil, err
	}

	exp, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["exp"]), 10, 64)
	if err != nil {
		return nil, err
	}

	return &JWTAccessDetails{
		UUID:    accessUUID,
		UserID:  uint(userID),
		Token:   token.Raw,
		Expires: exp,
	}, nil
}

func checkTokenInStorage(ad *JWTAccessDetails) (uint64, error) {
	idStr, err := cache.GetRedisInstance().Get(ad.UUID).Result()
	if err != nil {
		return 0, err
	}
	userID, _ := strconv.ParseUint(idStr, 10, 64)
	if ad.UserID != uint(userID) {
		return 0, fmt.Errorf("User ID doesn't match")
	}
	return userID, nil
}

func verifyToken(tokenString, secret string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

/*JWTExtractToken extracts token from authorization header*/
func JWTExtractToken(authHeader string) string {
	tokSplit := strings.Split(authHeader, " ")
	if len(tokSplit) == 2 {
		return tokSplit[1]
	}
	return ""
}
