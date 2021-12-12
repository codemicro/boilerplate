package state

import (
	"fmt"
	"github.com/codemicro/mars/internal/db"
	"github.com/gofiber/fiber/v2"
	"github.com/kkyr/fig"
	"github.com/spf13/afero"
)

type State struct {
	Config *Config
    Data *db.DB
	FS afero.Fs
	HTTP *fiber.App
}

func New() *State {
	return &State{
		FS: afero.NewOsFs(),
	}
}

type Config struct {
	Server struct {
		Host string `fig:"host" default:"127.0.0.1"`
		Port int `fig:"port" default:"8080"`
	}
	Database struct {
		Filename string `fig:"filename" default:"mars-imaging.bolt.db"`
	}
}

func LoadConfig() (*Config, error) {
	cfg := new(Config)
	if err := fig.Load(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

func (cfg *Config) GetHTTPServeAddress() string {
	return fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
}