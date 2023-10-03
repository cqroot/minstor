//go:build wireinject
// +build wireinject

package obs

import (
	"github.com/cqroot/minstor/internal/config"
	"github.com/cqroot/minstor/internal/logger"
	"github.com/google/wire"
)

var ConfigSet = wire.NewSet(
	config.NewObsConfig,
	wire.Bind(new(config.Config), new(*config.ObsConfig)),
)

func InitObs() (*Obs, error) {
	wire.Build(New, ConfigSet, logger.New)
	return &Obs{}, nil
}
