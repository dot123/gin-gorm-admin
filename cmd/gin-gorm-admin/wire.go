//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	v1 "github.com/dot123/gin-gorm-admin/api/v1"
	jwt "github.com/dot123/gin-gorm-admin/internal/middleware/jwt"
	"github.com/dot123/gin-gorm-admin/internal/models"
	"github.com/dot123/gin-gorm-admin/internal/service"
	"github.com/dot123/gin-gorm-admin/pkg/fileStore"
	"github.com/google/wire"
)

func BuildInjector(*fileStore.Local) (*Injector, func(), error) {
	wire.Build(InitGormDB, InitGinEngine, models.RepoSet, service.ProviderSet, jwt.JWTSet, v1.ProviderSet, RouterSet, InjectorSet)
	return new(Injector), nil, nil
}
