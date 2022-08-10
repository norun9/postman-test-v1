package config

import (
	"fmt"
	"strings"

	"path/filepath"
	"runtime"
	"time"

	"github.com/spf13/viper"
)

// Config :
type Config struct {
	Test AppConfig `mapstructure:"test"`
}

// CORS :
type CORS struct {
	AllowedOrigins []string `mapstructure:"allowed_origins" validate:"dive,required"`
}

// HTTP :
type HTTP struct {
	Port      int    `mapstructure:"port" validate:"required"`
	WebDomain string `mapstructure:"web_domain" validate:"required"`
	Protcol   string `mapstructure:"protcol" validate:"required"`
	CORS      CORS   `mapstructure:"cors" validate:"required"`
}

// AppConfig :
type AppConfig struct {
	HTTP  HTTP  `mapstructure:"http"`
	MySQL MySQL `mapstructure:"mysql"`
}

// MySQL :
type MySQL struct {
	DBHost          string        `mapstructure:"db_host" validate:"required"`
	DBName          string        `mapstructure:"db_name" validate:"required"`
	DBUserName      string        `mapstructure:"db_user_name" validate:"required"`
	DBPassword      string        `mapstructure:"db_password" validate:"required"`
	DBPort          string        `mapstructure:"db_port" validate:"required"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns" validate:"required"`
	MaxOpenConns    int           `mapstructure:"max_open_conns" validate:"required"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime" validate:"required"`
	Pseudo          bool
}

// Prepare :
func Prepare() AppConfig {
	viper.SetConfigName("config")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetConfigType("yaml")

	_, b, _, _ := runtime.Caller(0)
	configDir := filepath.Dir(b)
	pkgDir := filepath.Dir(configDir)
	backendDir := filepath.Dir(pkgDir)

	viper.AddConfigPath(backendDir)
	viper.AddConfigPath("./")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	viper.AutomaticEnv()

	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		panic(err)
	}

	var appConfig AppConfig
	env := viper.GetString("env.name")
	switch env {
	case "test":
		// no validate when test
		appConfig = c.Test
	default:
		panic(fmt.Sprintf("Unknown env: %s", env))
	}
	return appConfig
}
