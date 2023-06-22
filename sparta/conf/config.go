package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./conf/")

	if err := viper.ReadInConfig(); err != nil {
		panic(any(fmt.Sprintf("Read config file failed at:%s", err.Error())))
	}
}
