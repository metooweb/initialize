package initialize

import (
	"github.com/metooweb/goservice"
)

type Config struct {
	File  string
	Value interface{}
}

type Setting struct {
	Config   *Config
	Service  *goservice.Config
	TimeZone string
}
