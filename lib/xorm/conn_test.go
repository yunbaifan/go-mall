package xorm

import (
	"gorm.io/gorm/logger"
	"testing"
)

var (
	_confg DatabaseConf
)

func TestMain(m *testing.M) {
	_confg = DatabaseConf{
		Source:        "root:123456@tcp(host.docker.internal:23306)/go-mall?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai",
		MaxIdleConns:  10,
		MaxOpenConns:  10,
		SlowThreshold: 100,
		LogLevel:      logger.LogLevel(4),
		Colorful:      false,
	}
	m.Run()
}

func TestConnectMysql(t *testing.T) {
	db, err := ConnectMysql(_confg)
	if err != nil {
		t.Error(err)
	}
	t.Logf("init db success: %v", db)
}
