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

func Init(setting *Setting) (err error) {

	if setting.TimeZone == "" {
		setting.TimeZone = "Asia/Shanghai"
	}

	time.Local, _ = time.LoadLocation(setting.TimeZone)

	if err = ini.MapTo(setting.Config.Value, setting.Config.File); err != nil {
		err = errors.WithStack(err)
		return
	}
	if setting.Service != nil {
		err = goservice.Init(setting.Service)
	}
	return
}

func init() {

	RootPath = filepath.Dir(os.Args[0])

}
