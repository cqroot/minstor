package obs

import (
	"fmt"

	"github.com/cqroot/minstor/internal/config"
	"github.com/cqroot/minstor/internal/logger"
	"github.com/gin-gonic/gin"
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

func (s Obs) Run() error {
	r := gin.Default()

	return r.Run(fmt.Sprintf("%s:%d", s.config.BindIp(), s.config.BindPort()))
}
