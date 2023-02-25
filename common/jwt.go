package common

import (
	beego "github.com/beego/beego/v2/server/web"
	j "github.com/golang-jwt/jwt/v5"
	"golang.org/x/sync/singleflight"
	"ruoyi-vue-beego/models/system"
	"time"
)

var g = &singleflight.Group{}

type JWT struct {
	SigningKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(beego.AppConfig.DefaultString("jwt.signingkey", "abcdefghijklmnopqrstuvwxyz")),
	}
}

type CustomClaims struct {
	system.LoginUser
	BufferTime int64
	j.RegisteredClaims
}

func (jwt *JWT) CreateClaims(loginUser system.LoginUser) CustomClaims {
	claims := CustomClaims{
		LoginUser: loginUser,
		RegisteredClaims: j.RegisteredClaims{
			NotBefore: j.NewNumericDate(time.Now()),                                                          // 签名生效时间
			ExpiresAt: j.NewNumericDate(time.Unix(beego.AppConfig.DefaultInt64("jwt.expirestime", 3600), 0)), // 过期时间
		},
	}
	return claims
}

func (jwt *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := j.NewWithClaims(j.SigningMethodHS256, claims)
	return token.SignedString(jwt.SigningKey)
}

func (jwt *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := j.ParseWithClaims(tokenString, &CustomClaims{}, func(token *j.Token) (i interface{}, e error) {
		return jwt.SigningKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims.(*CustomClaims), nil
}

// CreateTokenByOldToken 旧token 换新token 使用归并回源避免并发问题
func (jwt *JWT) CreateTokenByOldToken(oldToken string, claims CustomClaims) (string, error) {
	v, err, _ := g.Do("JWT:"+oldToken, func() (interface{}, error) {
		return jwt.CreateToken(claims)
	})
	return v.(string), err
}
