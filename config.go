package initialize

import (
	"github.com/metooweb/goservice"
)

type Config struct {
	TimeZone string
	Service  *goservice.Config
	Config struct {
		File  string
		Value interface{}
	}
}
