package goservice

import (
	"github.com/kardianos/service"
)

type program struct {
	config *Config
}

func (p *program) Start(s service.Service) error {

	go p.run()

	return nil
}
func (p *program) run() {

	var err error

	if p.config.Run != nil {
		err = p.config.Run()
	}

	if err != nil {
		panic(err)
	}

	return
}

func (p *program) Stop(s service.Service) (err error) {

	if p.config.Stop != nil {
		err = p.config.Stop()
	}

	return
}
