package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/url"
	"time"
)

type mysqlConf struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Database string `mapstructure:"database"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Timezone string `mapstructure:"timezone"`
}

var mysqlDbConfig = &mysqlConf{}

func initMysqlDb() *gorm.DB {

	if err := viper.UnmarshalKey("mysql", mysqlDbConfig); err != nil {
		log.Fatalf("Parse config.mysql segment error: %s\n", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=%s",
		mysqlDbConfig.User,
		mysqlDbConfig.Password,
		mysqlDbConfig.Host,
		mysqlDbConfig.Port,
		mysqlDbConfig.Database,
		url.QueryEscape(mysqlDbConfig.Timezone),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Connect mysql error: %s", err)
	}

	sqlDB, _ := db.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}
