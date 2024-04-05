package config

import (
	"github.com/yunbaifan/go-mall/lib/xorm"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	// mysql database config
	MySQL xorm.DatabaseConf `json:"mysql" yaml:"mysql"`
	Auth  Auth              `json:"auth"`
}

type Auth struct {
	// jwt secret
	JwtSecret string `json:"jwtSecret" default:"@G5U5*hY2e3E" yaml:"jwtSecret"`
	Expire    int64  `json:"expire" default:"86400" yaml:"expire"`
}
