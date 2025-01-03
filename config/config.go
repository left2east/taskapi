package config

import (
	"bytes"
	"embed"

	"github.com/spf13/viper"
)


type Db struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Pass     string `yaml:"pass"`
	Dbname   string `yaml:"dbname"`
}

// Config 配置文件结构体
type Config struct {
	Db          Db          `yaml:"db"`
}

var AppConfig Config

// InitConfig 初始化配置文件
func InitConfig(embeddedConfigFiles embed.FS) {
	configFile := "config/config.yml"
	viper.SetConfigType("yml")
	configBytes, err := embeddedConfigFiles.ReadFile(configFile)
	if err != nil {
		panic(err)
	}

	err = viper.ReadConfig(bytes.NewBuffer(configBytes))
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		panic(err)
	}
}
