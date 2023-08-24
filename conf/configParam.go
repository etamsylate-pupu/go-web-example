package conf

import (
	"github.com/spf13/viper"
	"log"
)

type configParams struct {
	AesKey      string `mapstructure:"aeskey"`
	BuildInUser string `mapstructure:"buildInUser"` //内置用户
	SuperRole   string `mapstructure:"superRole"`
}

// ConfigParamsConf is config of api need
var ConfigParamsConf = &configParams{}

func initConfigParamConf() {
	if err := viper.UnmarshalKey("configParam", ConfigParamsConf); err != nil {
		log.Fatalf("Parse config.configParam segment error: %s", err)
	}
}
