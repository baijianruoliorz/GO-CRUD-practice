package common

import (
	"github.com/dgrijalva/jwt-go"
	"oceanlearn.teach/ginessential/model"
	"time"
)

var jwtkey=[]byte("a_secret_crect")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}
func ReleaseToken(user model.User)(string,error){
	expriationTime:=time.Now().Add(7*24*time.Hour)
	claims:=&Claims{
		UserId:         user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expriationTime.Unix(),
			IssuedAt: time.Now().Unix(),
			Issuer: "oceanlearn.tech",
			Subject: "user token",
		},
	}
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	tokenString,err:=token.SignedString(jwtkey)
	if err!=nil {
		return "", err
	}
	return tokenString,nil
	}
func ParseToken(tokenString string) (*jwt.Token,*Claims,error) {
	claims:=&Claims{}
	token,err:=jwt.ParseWithClaims(tokenString,claims, func(token *jwt.Token) (i interface{},err error) {
		return jwtkey,nil
	})
	return token,claims,err
}