//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/google/wire"
	"server/pkg/infra/repository"
	"server/pkg/repo/impl"
	"server/routers"
	"server/routers/api"
	v1 "server/routers/api/v1"
	"server/service"
)

func InitApp() (*routers.Server, error) {
	wire.Build(
		routers.NewServer,
		api.NewAuthService,
		service.NewUserService,
		service.NewArticleService,
		impl.NewUserRepo,
		impl.NewArticleRepo,
		repository.InitDB,
		v1.NewArticleService,
	)
	return &routers.Server{}, nil
}
