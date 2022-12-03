package Global

import (
	"github.com/dgrijalva/jwt-go" //生成JWT和解析JWT的功能。
	"cx/Model"
	"time"
)

//跨域认证解决方案 规定了Token实现方式
//服务端认证通过后，会生成一个JSON对象，经过签名后得到一个Token（令牌）再发回给用户，
//用户后续请求只需要带上这个Token，服务端解密之后就能获取该用户的相关信息了。

// 用于签名的字符串
var jwtKey_cx = []byte("s_secret_crect")

type Claims_cx struct {
	UserId string
	jwt.StandardClaims
}

//使用默认声明创建jwt
func ReaeaseTokec(user Model.Employee_cx) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour) //过期时间
	//创建Claims
	claims := &Claims_cx{
		UserId: user.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "cx",//签发人
			Subject:   "user token",
		},
	}
	// 生成token对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 生成签名字符串
	tokenString, err := token.SignedString(jwtKey_cx)
	if err != nil { //解析token失败
		return "", err
	}
	return tokenString, err
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*jwt.Token, *Claims_cx, error) {
	claims := &Claims_cx{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey_cx, nil
	})
	return token, claims, err
}
