package common

import (
	"Appitems/models"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// 密钥
var (
	jwtPrivateKey, _ = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	//jwtPublicKey     = &jwtPrivateKey.PublicKey
	jwtKey = []byte("a_secret_crect")
)

// 生成 token
func ReleaseToken(user models.UserDatabase) (string, error) {
	// Token 过期时间
	expirationTime := time.Now().Add(Week)

	// Token 中的用户信息
	claims := &models.Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "oceanlearn.tech",
			Subject:   "user access token",
		},
	}

	// 使用 ES256 算法创建 Token
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	// 使用私钥签名 Token
	tokenString, err := token.SignedString(jwtPrivateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// 对token解密看一下
/*
eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsImV4cCI6MTY4NzE0NTg0MSwiaWF0IjoxNjg2NTQxMDQxLCJpc3MiOiJvY2VhbmxlYXJuLnRlY2giLCJzdWIiOiJ1c2VyIGFjY2VzcyB0b2tlbiJ9.E--woGubSsKh0aHdm4KTLKWuTA6XuH9P-_zQN6sRhCi3zCMAmOpxywbdZgsIDMLP2_50OdJzYSCwDg01hfYlzg
第一部分是请求头，第二部分是封装的东西 第三是加密算法
*/

// 第二种

//生成token算法加密

func ReleaseTokenUser(user models.UserDatabase) (string, error) {
	// Token 过期时间
	expirationTime := time.Now().Add(Week)
	claims := &models.Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "oceanlearn.tech",
			Subject:   "user access token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	Tokenstring, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return Tokenstring, nil

}

// 解析token
func ParseToken(tokenString string) (*jwt.Token, *models.Claims, error) {
	Cliams := &models.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, Cliams, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, Cliams, err
}
