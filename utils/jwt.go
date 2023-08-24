package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// jwtSecret JWT私钥
const jwtSecret = "moooooo"

// expireDuration JWT过期时间
var expireDuration = time.Hour * 72

//authClaims jwt claims
type authClaims struct {
	jwt.RegisteredClaims
	Phone  string `json:"phone"`
	UserID int    `json:"user_id"`
}

//TokenUserInfo token 解析后的user info
type TokenUserInfo struct {
	Phone  string `json:"phone"`
	UserID int    `json:"user_id"`
}

//GenerateJWT 产生jwt
func GenerateJWT(phone string, userID int) (string, error) {
	exp := &jwt.NumericDate{Time: time.Now().Add(expireDuration)}
	iat := &jwt.NumericDate{Time: time.Now()}

	claims := &authClaims{
		jwt.RegisteredClaims{
			ExpiresAt: exp,
			IssuedAt:  iat,
		},
		phone,
		userID,
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tenc, err := t.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	//jwt 无法主动token 过期 常见的场景为后台踢出用户或封禁用户。
	return tenc, nil
}

// ValidateJWT 验证jwt
func ValidateJWT(encToken string) (TokenUserInfo, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("JWT签名算法不一致，需要HS256，当前算法：%s", token.Header["alg"])
		}

		return []byte(jwtSecret), nil
	}

	token, err := jwt.ParseWithClaims(encToken, &authClaims{}, keyFunc)
	if err != nil {
		return TokenUserInfo{}, err
	}

	if claims, ok := token.Claims.(*authClaims); ok && token.Valid {

		return TokenUserInfo{
			claims.Phone,
			claims.UserID,
		}, nil
	}

	return TokenUserInfo{}, fmt.Errorf("JWT不可用")
}
