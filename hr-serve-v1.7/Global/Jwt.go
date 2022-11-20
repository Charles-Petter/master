package Global

import (
	"github.com/dgrijalva/jwt-go"
	"m/Model"
	"time"
)

var jwtKey = []byte("s_secret_crect")

type Claims struct {
	UserId string
	jwt.StandardClaims
}

func ReaeaseTokec(user Model.Employee) (string, error){
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: user.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt: time.Now().Unix(),
			Issuer: "oceanlearn.tech",
			Subject: "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, err
}

//func ReaeaseTokec2(user Model.Operator) (string, error){
//	expirationTime := time.Now().Add(7 * 24 * time.Hour)
//	claims := &Claims{
//		UserId: user.Id,
//		StandardClaims: jwt.StandardClaims{
//			ExpiresAt: expirationTime.Unix(),
//			IssuedAt: time.Now().Unix(),
//			Issuer: "oceanlearn.tech",
//			Subject: "user token",
//		},
//	}
//
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//	tokenString, err := token.SignedString(jwtKey)
//	if err != nil {
//		return "", err
//	}
//	return tokenString, err
//}
//
//func ReaeaseTokec3(user Model.Designer) (string, error){
//	expirationTime := time.Now().Add(7 * 24 * time.Hour)
//	claims := &Claims{
//		UserId: user.Id,
//		StandardClaims: jwt.StandardClaims{
//			ExpiresAt: expirationTime.Unix(),
//			IssuedAt: time.Now().Unix(),
//			Issuer: "oceanlearn.tech",
//			Subject: "user token",
//		},
//	}
//
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//	tokenString, err := token.SignedString(jwtKey)
//	if err != nil {
//		return "", err
//	}
//	return tokenString, err
//}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error){
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, claims, err
}
