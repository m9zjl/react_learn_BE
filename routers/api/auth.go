package api

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"server/service"
	"time"
)

type AuthService struct {
	userService *service.UserService
}

func NewAuthService(userService *service.UserService) *AuthService {
	return &AuthService{
		userService: userService,
	}
}

type authEntity struct {
	Email  string `valid:"Required; MaxSize(128)" json:"email"`
	Passwd string `valid:"Required; MaxSize(128)" json:"passwd"`
}

func (a *AuthService) LoginHandler(c *gin.Context) {
	// TODO login
	req := &authEntity{}
	_ = c.BindJSON(req)
	email, passwd := req.Email, req.Passwd

	user, err := a.userService.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "can't get user by email",
		})
		return
	}
	if user == nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "user doesn't exists",
		})
		return
	}

	if passwd != user.Passwd {
		c.JSON(http.StatusConflict, gin.H{
			"error": "email and password doesn't match",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.MapClaims{
		"email":     user.Email,
		"id":        user.ID,
		"ExpiresAt": time.Now().Add(time.Hour * 12).Unix(),
	})
	jwtToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"token": jwtToken,
	})
}
