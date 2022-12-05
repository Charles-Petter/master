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
	UserId_cx string
	jwt.StandardClaims
}

//使用默认声明创建jwt
func ReaeaseTokec(user_cx Model.Employee_cx) (string, error) {
	expirationTime_cx := time.Now().Add(7 * 24 * time.Hour) //过期时间
	//创建Claims
	claims_cx := &Claims_cx{
		UserId_cx: user_cx.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime_cx.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "cx",//签发人
			Subject:   "user token",
		},
	}
	// 生成token对象
	token_cx := jwt.NewWithClaims(jwt.SigningMethodHS256, claims_cx)
	// 生成签名字符串
	tokenString_cx, err_cx := token_cx.SignedString(jwtKey_cx)
	if err_cx != nil { //解析token失败
		return "", err_cx
	}
	return tokenString_cx, err_cx
}

// ParseToken 解析JWT
func ParseToken(tokenString_cx string) (*jwt.Token, *Claims_cx, error) {
	claims_cx := &Claims_cx{}
	token, err_cx := jwt.ParseWithClaims(tokenString_cx, claims_cx, func(token *jwt.Token) (i interface{}, err_cx error) {
		return jwtKey_cx, nil
	})
	return token, claims_cx, err_cx
}
