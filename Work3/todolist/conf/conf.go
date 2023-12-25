package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

type Conf struct {
	MysqlConfig MysqlConfig
	RedisConfig RedisConfig
}

type MysqlConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
	Charset  string `yaml:"charset"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
}
type RedisConfig struct {
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
}

var Config = Conf{}

func InitConfig() {
	reader := viper.New()
	reader.AddConfigPath("./conf/local")
	reader.SetConfigType("yaml")
	reader.SetConfigName("config")

	if err := reader.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := reader.Unmarshal(&Config); err != nil {
		panic(err)
	}

	fmt.Println("Config loaded")
}
