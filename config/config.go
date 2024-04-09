package config

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"strings"
)

//go:embed default.yaml
var defaultConfig []byte

type Config struct {
	Base `mapstructure:",squash"`
}

type Base struct {
	Environment string
	Server      ServerConfig `yaml:"server" mapstructure:"server"`
	Postgres    *Postgres    `yaml:"postgres" mapstructure:"postgres"`
}

// ServerConfig hold http/grpc server config
type ServerConfig struct {
	Grpc ServerAddress `yaml:"grpc" mapstructure:"grpc"`
	Http ServerAddress `yaml:"http" mapstructure:"http"`
}

type ServerAddress struct {
	Host string `yaml:"host" mapstructure:"host"`
	Port int    `yaml:"port" mapstructure:"port"`
}

func (s *ServerAddress) String() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

type Postgres struct {
	Host     string `yaml:"host" mapstructure:"host"`
	Port     int    `yaml:"port" mapstructure:"port"`
	Database string `yaml:"database" mapstructure:"database"`
	Username string `yaml:"username" mapstructure:"username"`
	Password string `yaml:"password" mapstructure:"password"`
}

func (m *Postgres) FormatDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Ho_Chi_Minh",
		m.Host, m.Username, m.Password, m.Database, m.Port)
}

func Load() *Config {
	var cfg = &Config{}

	viper.SetConfigType("yaml")
	if err := viper.ReadConfig(bytes.NewBuffer(defaultConfig)); err != nil {
		log.Fatalf(fmt.Sprintf("Error reading config file: %s", err.Error()))
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf(fmt.Sprintf("viper.Unmarshal err: %v", err))
	}

	return cfg
}
