package config

import (
	"log"

	"github.com/spf13/viper"
)

type Configuration struct {
	Server     Server     `mapstructure:"server" json:"server" yaml:"server"`
	DataSource DataSource `mapstructure:"datasource" json:"datasource" yaml:"datasource"`
}

type Server struct {
	Host string `mapstructure:"host" json:"host" yaml:"host"`
	Post string `mapstructure:"post" json:"post" yaml:"post"`
	Mode string `mapstructure:"mode" json:"mode" yaml:"mode"`
}

type DataSource struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	User     string `mapstructure:"user" json:"user" yaml:"user"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Name     string `mapstructure:"name" json:"name" yaml:"name"`
}

var Cfg Configuration

func InitConfig() {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("configs")

	err := v.ReadInConfig()
	if err != nil {
		log.Fatalf("读取配置文件失败: %v", err)
	}

	if err := v.Unmarshal(&Cfg); err != nil {
		log.Fatalf("解析配置文件失败: %v", err)
	}

	log.Println("配置文件加载成功")
}
