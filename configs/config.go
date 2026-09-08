package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

var Conf *Config

func InitConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("读取配置文件失败")
		return err
	}
	if err := viper.Unmarshal(&Conf); err != nil {
		fmt.Println("反序列化失败")
		return err
	}
	return nil
}
