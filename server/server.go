package server

import (
	"fmt"

	ytrelay "github.com/mirror-media/yt-relay"
	"github.com/mirror-media/yt-relay/config"
	"github.com/mirror-media/yt-relay/whitelist"

	"github.com/gin-gonic/gin"
)

type Server struct {
	conf         *config.Conf
	Engine       *gin.Engine
	APIWhitelist ytrelay.APIWhitelist
}

func (s *Server) Run() error {
	return s.Engine.Run(fmt.Sprintf("%s:%d", s.conf.Address, s.conf.Port))
}

func New(c config.Conf) (*Server, error) {

	engine := gin.Default()

	s := &Server{
		conf:   &c,
		Engine: engine,
		APIWhitelist: &whitelist.API{
			Whitelist: c.Whitelists,
		},
	}
	return s, nil
}
