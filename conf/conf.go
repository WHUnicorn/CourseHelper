package conf

import (
	"encoding/json"
	"github.com/spf13/viper"
	"log"
)

var Config = struct {
	LogLevel     string `yaml:"logLevel,omitempty"`
	Cookie       string `yaml:"cookie,omitempty"`
	DatafilePath string `yaml:"datafilePath"`
}{
	LogLevel:     "debug",
	DatafilePath: "./data/default.yaml",
}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading config file:", err)
		return // 自动退出
	}

	Config.LogLevel = viper.GetString("logLevel")
	Config.Cookie = viper.GetString("cookie")
	Config.DatafilePath = viper.GetString("datafilePath")

	if Config.Cookie == "" {
		configStr, _ := json.Marshal(Config)
		log.Fatal("请配置cookie! 当前配置: ", string(configStr))
	}
}
