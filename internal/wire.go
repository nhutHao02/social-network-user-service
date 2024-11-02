//go:build wireinject
// +build wireinject

package internal

import (
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/nhutHao02/social-network-user-service/config"
	"github.com/nhutHao02/social-network-user-service/internal/api"
	"github.com/nhutHao02/social-network-user-service/internal/api/http"
	"github.com/nhutHao02/social-network-user-service/internal/api/http/v1"
	"github.com/nhutHao02/social-network-user-service/internal/application/imp"
	"github.com/nhutHao02/social-network-user-service/internal/infrastructure/user"
	"github.com/nhutHao02/social-network-user-service/pkg/redis"
)

var serverSet = wire.NewSet(
	api.NewSerVer,
)

var itemServerSet = wire.NewSet(
	http.NewHTTPServer,
)

var httpHandlerSet = wire.NewSet(
	v1.NewUserHandler,
)

var serviceSet = wire.NewSet(
	imp.NewUserService,
)

var repositorySet = wire.NewSet(
	user.NewUserCommandRepository,
	user.NewUserQueryRepository,
)

func InitializeServer(cfg *config.Config, db *sqlx.DB, rdb *redis.RedisClient) *api.Server {
	wire.Build(serverSet, itemServerSet, httpHandlerSet, serviceSet, repositorySet)
	return &api.Server{}
}
