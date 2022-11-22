package Global

import (
	"github.com/dgrijalva/jwt-go" //生成JWT和解析JWT的功能。
	"m/Model"
	"time"
)

//跨域认证解决方案 规定了Token实现方式
//服务端认证通过后，会生成一个JSON对象，经过签名后得到一个Token（令牌）再发回给用户，
//用户后续请求只需要带上这个Token，服务端解密之后就能获取该用户的相关信息了。

// 用于签名的字符串
var jwtKey = []byte("s_secret_crect")

type Claims struct {
	UserId string
	jwt.StandardClaims
}

//使用默认声明创建jwt
func ReaeaseTokec(user Model.Employee) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour) //过期时间
	//创建Claims
	claims := &Claims{
		UserId: user.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "oceanlearn.tech",
			Subject:   "user token",
		},
	}
	// 生成token对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 生成签名字符串
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, err
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, claims, err
}
