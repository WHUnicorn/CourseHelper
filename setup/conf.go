package setup

import (
	"encoding/json"
	"github.com/spf13/viper"
	"log"
	"reflect"
)

var Config = struct {
	LogLevel     string `yaml:"logLevel,omitempty"`
	Cookie       string `yaml:"cookie,omitempty"`
	DatafilePath string `yaml:"datafilePath"`
	Port         string `yaml:"port"`
	IsUnix       string `yaml:"isUnix"`
	Test         struct {
		AA string `yaml:"aa"`
	} `yaml:"test"`
}{
	LogLevel:     "debug",
	DatafilePath: "./resources/trainingPlans/cs.course",
	Port:         "12345",
	IsUnix:       "true",
}

// Elem()用于获取指针指向的值，如果不是接口或指针会panics
// Addr()用于获得值的指针
func setConf(value reflect.Value, lastFields ...string) {
	for i := 0; i < value.Elem().NumField(); i++ {
		field := value.Elem().Field(i)
		if field.Kind() == reflect.String {
			resKey := ""
			for _, lastField := range lastFields {
				resKey += lastField + "."
			}
			resKey += value.Type().Elem().Field(i).Name
			if tempParam := viper.GetString(resKey); tempParam != "" {
				field.Set(reflect.ValueOf(tempParam))
			}
		} else {
			// 回溯 (前进 => 处理 => 回退)
			lastFields = append(lastFields, value.Elem().Type().Field(i).Name)
			setConf(field.Addr(), lastFields...)
			lastFields = lastFields[:len(lastFields)-1]
		}
	}
}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("请您添加配置文件（将 config.yaml.demo 重命名为 config.yaml 并补充 cookie 字段）")
		return // 自动退出
	}

	setConf(reflect.ValueOf(&Config))

	if Config.Cookie == "" {
		configStr, _ := json.Marshal(Config)
		log.Fatal("请配置cookie! 当前配置: ", string(configStr))
	}
}
