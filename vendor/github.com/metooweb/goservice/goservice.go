package goservice

import (
	"github.com/kardianos/service"
	"os"
)

type Config struct {
	service.Config
	Run  func() error
	Stop func() error
}

func Init(cfg *Config) (err error) {

	var s service.Service

	prg := &program{
		config: cfg,
	}

	if s, err = service.New(prg, &cfg.Config); err != nil {
		return
	}

	if len(os.Args) > 1 {

		if err = service.Control(s, os.Args[1]); err != nil {
			return
		}

		return
	}

	if err = s.Run(); err != nil {
		return
	}

	return nil
}
