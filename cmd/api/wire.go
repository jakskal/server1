//+build wireinject

package main

import (
	"server1/cmd/handler"
	"server1/internal/auth"
	"server1/internal/register"
	"server1/internal/token"
	"server1/internal/user"

	"github.com/jinzhu/gorm"

	"github.com/google/wire"
)

var repositorySet = wire.NewSet(
	user.NewRepository,
	wire.Bind(new(user.RepositorySystem), new(*user.Repository)),
)

var serviceSet = wire.NewSet(
	user.NewService,
	auth.NewService,
	token.NewService,
	register.NewService,
)

var controllerSet = wire.NewSet(
	token.NewController,
	auth.NewController,
	register.NewController,
)

var handlerSet = wire.NewSet(
	handler.NewHandler,
)

var appSet = wire.NewSet(
	repositorySet,
	serviceSet,
	controllerSet,
	handlerSet,
)

func initHandler(db *gorm.DB) *handler.Handler {
	wire.Build(appSet)
	return &handler.Handler{}
}
