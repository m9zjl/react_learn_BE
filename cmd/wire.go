//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/google/wire"
	"server/pkg/infra/repository"
	"server/pkg/repo"
	"server/routers"
	"server/routers/api"
	"server/service"
)

func InitApp() (*routers.Server, error) {
	wire.Build(routers.NewServer, api.NewAuthService, service.NewUserService, repo.NewUserRepo, repository.InitDB)
	return &routers.Server{}, nil
}
