package obs

import (
	"github.com/cqroot/minstor/internal/config"
	"github.com/cqroot/minstor/internal/logger"
)

type Obs struct {
	config *config.ObsConfig
	logger logger.Logger
}

func New(config *config.ObsConfig, logger logger.Logger) *Obs {
	return &Obs{
		config: config,
		logger: logger,
	}
}

func (s Obs) Run() {
	s.logger.Info().Msg("Run OBS!")
}
