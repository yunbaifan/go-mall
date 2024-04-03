package xorm

import (
	"database/sql"
	xlogger "github.com/yunbaifan/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type DatabaseConf struct {
	Source        string          `json:"source" yaml:"source" default:""`
	MaxIdleConns  int             `json:"maxIdleConns" yaml:"maxIdleConns" default:"10"`      // 空闲中的最大连接数
	MaxOpenConns  int             `json:"maxOpenConns" yaml:"maxOpenConns" default:"10"`      // 打开到数据库的最大连接数
	SlowThreshold int             `json:"slowThreshold" yaml:"slowThreshold" default:"100ms"` // 慢查询阈值
	LogLevel      logger.LogLevel `json:"logLevel" yaml:"logLevel" default:"Info"`            // 日志级别
	Colorful      bool            `json:"colorful" yaml:"colorful" default:"false"`           // 是否显示颜色
}

func ConnectMysql(conf DatabaseConf) (db *gorm.DB, err error) {
	if db, err = gorm.Open(mysql.Open(conf.Source), &gorm.Config{
		Logger: xlogger.NewGormLogger(
			xlogger.WithSlowThreshold(time.Duration(conf.SlowThreshold)),
			xlogger.WithLogLevel(conf.LogLevel),
			xlogger.WithColorful(conf.Colorful),
		),
	}); err != nil {
		return nil, err
	}
	var (
		sqlDB *sql.DB
	)
	if sqlDB, err = db.DB(); err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(conf.MaxIdleConns)
	sqlDB.SetMaxOpenConns(conf.MaxOpenConns)

	return db, nil
}
