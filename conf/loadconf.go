package conf

import (
	"github.com/spf13/viper"
)

var config *Conf

func Config() *Conf {
	return config
}

type Conf struct {
	Server ServerConf `yaml:"server"`
	Mysql  MysqlConf  `yaml:"mysql"`
	Logger LoggerConf `yaml:"logger"`
}

type ServerConf struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Name string `yaml:"name"`
	Mode string `yaml:"mode"`
}

type MysqlConf struct {
	DataSourceName string `yaml:"dataSourceName"`
	MaxOpenConns   int    `yaml:"maxOpenConns"`
	MaxIdleConns   int    `yaml:"maxIdleConns"`
}

type LoggerConf struct {
	Level       string `yaml:"level"`
	FilePath    string `yaml:"filePath"`
	FileName    string `yaml:"fileName"`
	MaxFileSize uint64 `yaml:"maxFileSize"`
	ToFile      bool   `yaml:"toFile"`
	Formatter   string `yaml:"formatter"`
	Caller      bool   `yaml:"caller"`
}

func LoadConf() error {
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	config = &Conf{}
	err = viper.Unmarshal(config)
	if err != nil {
		panic(err)
	}

	return nil
}
