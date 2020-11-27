package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret []byte

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// GenerateToken generate tokens used for auth
func GenerateToken(username, password string) (string, string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		EncodeMD5(username),
		EncodeMD5(password),
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "rpcx-usercenter",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, expireTime.Format("2006-01-02 15:04:05"), err
}

// ParseToken parsing token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*Claims)
		if  ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

// RefreshToken refresh token
func RefreshToken(token string) (string, string, error) {
	claims, parseErr := ParseToken(token)

	if parseErr != nil {
		return "", "", parseErr
	}

	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	newClaims := Claims{
		claims.Username,
		claims.Password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    claims.StandardClaims.Issuer,
		},
	}

	newTokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
	newToken, err := newTokenClaims.SignedString(jwtSecret)

	return newToken, expireTime.Format("2006-01-02 15:04:05"), err
}