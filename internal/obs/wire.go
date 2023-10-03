//go:build wireinject
// +build wireinject

package obs

import (
	"github.com/cqroot/minstor/internal/config"
	"github.com/cqroot/minstor/internal/logger"
	"github.com/google/wire"
)

func InitObs() (*Obs, error) {
	wire.Build(New, config.NewObsConfig, logger.New)
	return &Obs{}, nil
}
