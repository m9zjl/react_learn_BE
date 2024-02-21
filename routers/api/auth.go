package api

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	log "log/slog"
	"net/http"
	"server/pkg/domain/entity"
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

type registerEntity struct {
	authEntity
	Nickname string `valid:"Required; MaxSize(64)" json:"nickname"`
}

func (a *AuthService) Register(c *gin.Context) {
	// TODO login
	req := &registerEntity{}
	_ = c.BindJSON(req)
	yace := c.GetHeader("yace")

	if "y" == yace {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
		return
	}
	email, passwd, nickname := req.Email, req.Passwd, req.Nickname
	user := &entity.User{
		Email:       email,
		Nickname:    nickname,
		Passwd:      passwd,
		GmtCreate:   time.Now(),
		GmtModified: time.Now(),
	}
	success, err := a.userService.AddUser(user)
	if err != nil {
		log.Error("system error:%v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "system error",
		})
		return
	}
	if !success {
		c.JSON(http.StatusOK, gin.H{
			"error":   "user already exists",
			"success": false,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": success,
	})
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
		c.JSON(http.StatusNotFound, gin.H{
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
