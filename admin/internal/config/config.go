package config

import (
	"github.com/yunbaifan/go-mall/lib/xorm"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	// mysql database config
	MySQL xorm.DatabaseConf `json:"mysql" yaml:"mysql"`
}
