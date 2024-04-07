package config

import (
	"github.com/yunbaifan/go-mall/lib/xorm"
	"github.com/zeromicro/go-zero/rest"
	"golang.org/x/text/language"
)

type Config struct {
	rest.RestConf
	// mysql database config
	MySQL xorm.DatabaseConf `json:"mysql" yaml:"mysql"`
	Auth  Auth              `json:"auth"`
	Lang  language.Tag      `json:"lang" default:"zh"`
}

type Auth struct {
	// jwt secret
	JwtSecret string `json:"jwtSecret" default:"@G5U5*hY2e3E" yaml:"jwtSecret"`
	Expire    int64  `json:"expire" default:"86400" yaml:"expire"`
}
