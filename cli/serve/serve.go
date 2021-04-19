package serve

import (
	"errors"

	"github.com/mirror-media/yt-relay/cli"
	"github.com/mirror-media/yt-relay/relay"
	"github.com/mirror-media/yt-relay/server"
	"github.com/mirror-media/yt-relay/server/route"
)

var serveFlags = []string{"address", "port", "config"}

func serveMain(args []string, c cli.Conf) error {
	cfg := c.CFG
	if c.CFG == nil {
		return errors.New("config file is nil")
	}
	server, err := server.New(*cfg)
	if err != nil {
		return nil
	}

	relayService, err := relay.New(cfg.ApiKey)
	if err != nil {
		return err
	}

	_ = route.Set(server.Engine, cfg.AppName, relayService, server.APIWhitelist, cfg.Cache, server.Cache)

	return server.Run()
}

var Command = &cli.Command{Flags: serveFlags, Main: serveMain}
