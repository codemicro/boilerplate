package app

import (
	"github.com/codemicro/mars/internal/db"
	"github.com/codemicro/mars/internal/state"
	"github.com/gofiber/fiber/v2"
)

func Start() error {

	// Setup

	st := state.New()

	if cfg, err := state.LoadConfig(); err != nil {
		return err
	} else {
		st.Config = cfg
	}

	if data, err := db.New(st.Config.Database.Filename); err != nil {
		return err
	} else {
		st.Data = data
	}

	if err := setupHTTPServer(st); err != nil {
		return err
	}

	// Run

	if err := st.HTTP.Listen(st.Config.GetHTTPServeAddress()); err != nil {
		return err
	}

	return nil
}

func setupHTTPServer(st *state.State) error {

	st.HTTP = fiber.New()

	e := &endpoints{state: st}

	st.HTTP.Get("/", e.uiIndex)

	return nil
}

type endpoints struct {
	state *state.State
}