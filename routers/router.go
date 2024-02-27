package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/pkg/middleware"
	"server/routers/api"
	v1 "server/routers/api/v1"
)

type Server struct {
	engine *gin.Engine
}

func NewServer(
	authService *api.AuthService,
	articleService *v1.ArticleService,
) *Server {
	gin.SetMode(gin.DebugMode)
	r := gin.New()

	// Use logger from Gin
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	//r.Use(cors.Default())
	r.Use(middleware.Cors())

	//// Swagger docs
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	//
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	})
	//Request
	r.POST("/login", authService.LoginHandler)

	r.POST("logout", authService.LogoutHandler)

	r.POST("/register", authService.Register)

	// user api
	{
		r.GET("/current_user", authService.GetCurrentUserInfo)
	}
	// article api
	{
		r.GET("/articles", articleService.GetArticles)
		r.POST("/articles", articleService.GetArticles)
		r.DELETE("/article:id", articleService.DeleteArticle)
		r.PUT("/article", articleService.Add)
	}

	userRouterGroup := r.Group("/user")
	{
		userRouterGroup.GET("/get")
	}

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
