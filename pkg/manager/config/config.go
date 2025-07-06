package config

import (
	"github.com/jinzhu/configor"
)

type Config struct {
	APPName string `default:"sentinel-manager"`

	Host       Host       `yaml:"host"`
	Database   Database   `yaml:"database"`
	Migrate    Migrate    `yaml:"migrate"`
	CORSConfig CORSConfig `yaml:"cors-config"`
	Encryption string     `yaml:"encryption" default:"enigma.yml"`
}

func New(c *Config, configPath string) (*Config, error) {
	if c == nil {
		c = &Config{}
	}
	if err := configor.Load(c, configPath); err != nil {
		return nil, err
	}
	return c, nil
}

type Host struct {
	Port                   int32  `env:"SENTINEL_HOST_PORT"             yaml:"port"             default:"8099"`
	TlsEnable              bool   `env:"SENTINEL_HOST_TLS_ENABLE"       yaml:"tls-enable"       default:"false"`
	TlsCertificateFilename string `env:"SENTINEL_HOST_TLS_CRT_FILENAME" yaml:"tls-crt-filename" default:"manager.crt"`
	TlsPrivateKeyFilename  string `env:"SENTINEL_HOST_TLS_KEY_FILENAME" yaml:"tls-key-filename" default:"manager.key"`

	XAuthToken bool `default:"false"`
}

type Database struct {
	Type            string `default:"mysql"`
	Protocol        string `default:"tcp"`
	Host            string `env:"SENTINEL_DB_HOST"`
	Port            string `env:"SENTINEL_DB_PORT"`
	DBName          string `env:"SENTINEL_DB_SCHEME"`
	Username        string `env:"SENTINEL_DB_SERVER_USERNAME"`
	Password        string `env:"SENTINEL_DB_SERVER_PASSWORD"`
	MaxOpenConns    int    `default:"15"`
	MaxIdleConns    int    `default:"5"`
	MaxConnLifeTime int    `default:"1"`
	ShowSQL         bool   `default:"false"`
	LogLevel        string `default:"warn"`
}

type Migrate struct {
	Source string `yaml:"source" default:"./schema"`
}

type CORSConfig struct {
	AllowOrigins string `env:"SENTINEL_CORSCONFIG_ALLOW_ORIGINS" yaml:"allow-origins,omitempty"`
	AllowMethods string `env:"SENTINEL_CORSCONFIG_ALLOW_METHODS" yaml:"allow-methods,omitempty"`
}
