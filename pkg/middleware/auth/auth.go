package auth

import (
	"github.com/duke-git/lancet/v2/strutil"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"server/pkg/domain/dto"
	"server/pkg/domain/entity"
)

func GetLoginUser(c *gin.Context) *entity.User {
	token, _ := c.Cookie("token")
	if strutil.IsBlank(token) {
		return nil
	}

	parseToken, _ := jwt.ParseWithClaims(token, &dto.UserAuthClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	authClaims := parseToken.Claims.(*dto.UserAuthClaim)
	err := authClaims.Valid()
	if err != nil {
		return nil
	}
	c.Set("user", authClaims.User)
	return &authClaims.User
}

func IsLogin(c *gin.Context) bool {
	return GetLoginUser(c) != nil
}
