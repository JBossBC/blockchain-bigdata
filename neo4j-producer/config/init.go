package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigType("xml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("读取配置文件出错:%s", err.Error()))
	}
}

func GetConfig(key string) any {
	return viper.Get(key)
}
func GetString(key string) string {
	return viper.GetString(key)
}
func GetBool(key string) bool {
	return viper.GetBool(key)
}
