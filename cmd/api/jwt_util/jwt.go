package jwtutil

import (
	"douyin-micro/pkg/constants"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/golang-jwt/jwt/v4"
)

type CustomClaims struct {
	UID int64 `json:"uid"`
	jwt.RegisteredClaims
}

func CreateToken(uid int64) (string, error) {
	claims := CustomClaims{
		uid,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(constants.JwtTtl)),
			Issuer:    constants.JwtIssuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	tokenStr, err := token.SignedString([]byte(constants.JwtSecret))
	if err != nil {
		klog.Error(err)
		return "", err
	}
	return tokenStr, nil
}

func ParseToken(tokenStr string) (int64, string, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(t *jwt.Token) (any, error) {
		return []byte(constants.JwtSecret), nil
	})
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims.UID, claims.Issuer, nil
	}

	klog.Error(err)
	return 0, "", nil
}
