package conf

import (
	"log"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

// MysqlClient is mysql client
var MysqlClient *gorm.DB

// RedisClient is redis client
var RedisClient *redis.Client

// AppLog application panic or other error log
var AppLog *logrus.Logger

// AccessLog api calls
var AccessLog *logrus.Logger

// Load load config and init mysql and redis client
func Load() {

	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Read config file error: %s", err)
	}

	initServerConf()
	initConfigParamConf()

	// init mysql and redis
	MysqlClient = initMysqlDb()
	//RedisClient = initRedisConf()

	// init all the log targets
	AppLog = initApplicationLog()
	AccessLog = initAccessLog()
}
