package initialize

import (
	"time"
	"github.com/go-ini/ini"
	"github.com/pkg/errors"
	"github.com/metooweb/goservice"
	"path/filepath"
	"os"
)

var RootPath string

func Init(cfg *Config) (err error) {

	if cfg.TimeZone == "" {
		cfg.TimeZone = "Asia/Shanghai"
	}

	time.Local, _ = time.LoadLocation(cfg.TimeZone)

	if err = ini.MapTo(cfg.Config, cfg.Config.File); err != nil {
		err = errors.WithStack(err)
		return
	}

	err = goservice.Init(cfg.Service)

	return
}

func init() {

	RootPath = filepath.Dir(os.Args[0])

}
