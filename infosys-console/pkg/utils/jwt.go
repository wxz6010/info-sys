package utils

import (
	"com.alex.infosys/pkg/errs"
	"com.alex.infosys/pkg/g"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const secretKey = "345678hjjhan,;a1221jd--=mvjhak"

func GenJwtToken(claim *g.LoginClaims) (string, error) {

	// 将 uid，用户角色， 过期时间作为数据写入 token 中
	claim.ExpiresAt = time.Now().Add(time.Hour*24*7).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	// SecretKey 用于对用户数据进行签名，不能暴露
	return token.SignedString([]byte(secretKey))
}

func ParserToken(tokenString string) (*g.LoginClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &g.LoginClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errs.Unauthorized.New("token不可用")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errs.Unauthorized.New("token过期")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errs.Unauthorized.New("无效的token")
			} else {
				return nil, errs.Unauthorized.New("token不可用")
			}
		}
	}

	if claim, ok := token.Claims.(*g.LoginClaims); ok && token.Valid {
		return claim, nil
	}

	return nil, errs.Unauthorized.New("无效的token")
}
