package core

import (
	"context"

	"olympics/lib/server"
	"olympics/pkg/core/entities"
	"olympics/pkg/storage"

	"github.com/heetch/confita"
)

type Config struct {
	Server  server.Config  `yaml:"server" config:"server"`
	Storage storage.Config `yaml:"storage"`
	Auth    struct {
		Disabled bool                                `yaml:"disabled"`
		Handlers []entities.HandlerDescriptionConfig `yaml:"handlers"`
	}
}

func ParseCfg(context context.Context, loader *confita.Loader) (*Config, error) {
	conf := &Config{}
	if loader != nil {
		err := loader.Load(context, conf)
		if err != nil {
			return nil, err
		}
	}
	return conf, nil
}
