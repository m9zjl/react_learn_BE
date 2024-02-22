package api

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	log "log/slog"
	"net/http"
	"server/pkg/domain/dto"
	"server/pkg/domain/entity"
	"server/pkg/middleware/auth"
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
	Username string `valid:"Required; MaxSize(64)" json:"username"`
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
	email, passwd, username := req.Email, req.Passwd, req.Username
	user := &entity.User{
		Email:       email,
		Username:    username,
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

func (a *AuthService) LogoutHandler(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (a *AuthService) LoginHandler(c *gin.Context) {
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
		c.JSON(http.StatusOK, gin.H{
			"error":   "user doesn't exists",
			"success": false,
		})
		return
	}

	if passwd != user.Passwd {
		c.JSON(http.StatusOK, gin.H{
			"error": "email and password doesn't match",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &dto.UserAuthClaim{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 12).Unix(),
		},
		User: entity.User{
			Email:       user.Email,
			ID:          user.ID,
			Username:    user.Username,
			GmtModified: user.GmtModified,
			GmtCreate:   user.GmtCreate,
			Ext:         user.Ext,
		},
	})
	jwtToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.SetCookie("token", jwtToken, 3600*24, "/", "localhost", false, true)
	user.Passwd = ""
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"user":    user,
	})
}

func (a *AuthService) GetCurrentUserInfo(c *gin.Context) {
	user := auth.GetLoginUser(c)
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
	return
}
