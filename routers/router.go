package routers

import (
	"github.com/gin-gonic/gin"
	"server/routers/api"
)

type Server struct {
	engine *gin.Engine
}

func NewServer(authService *api.AuthService) *Server {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	// Use logger from Gin
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	//// Swagger docs
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	//Request
	r.POST("/login", authService.LoginHandler)

	r.POST("/register", authService.Register)
	//
	//// Auth middleware
	//api := r.Group("/api", middleware.AuthorizationMiddleware)
	//
	//api.GET("users", userHandler.FindAll)
	//api.GET("users/:id", userHandler.FindByID)
	//api.POST("users", userHandler.Save)
	//api.DELETE("users/:id", userHandler.Delete)

	return &Server{engine: r}
}

func (s *Server) Run(addr ...string) error {
	return s.engine.Run(addr...)
}
