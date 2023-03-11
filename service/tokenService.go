package service

import (
	"ruoyi-vue-beego/common"
	"ruoyi-vue-beego/models/system"
)

var tokenServiceImpl *tokenService

func init() {
	tokenServiceImpl = &tokenService{}
}

type tokenService struct {
}

func GetTokenService() *tokenService {
	return tokenServiceImpl
}

func (tokenService *tokenService) createToken(loginUser system.LoginUser) (string, error) {
	jwt := common.NewJWT()
	claims := jwt.CreateClaims(loginUser)
	token, err := jwt.CreateToken(claims)
	return token, err
}
