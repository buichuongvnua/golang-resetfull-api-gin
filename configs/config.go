package config

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime"

	"github.com/spf13/viper"
)

// AppConfig : AppConfig
var AppConfig *Config

// Config : config for app
type Config struct {
	*viper.Viper
}

// New : return new config instance
func New(env string) *Config {
	log.Printf("Init app config for env: %s", env)
	config := &Config{
		Viper: viper.New(),
	}

	// Select the .env file
	_, filename, _, _ := runtime.Caller(0)
	folderPatch := path.Join(path.Dir(filename), "../settings")
	config.AddConfigPath(folderPatch)
	config.AddConfigPath("/app/settings")

	switch env {
	case "production":
		config.SetConfigName("settings.production")
	case "staging":
		config.SetConfigName("settings.staging")
	case "dev":
		config.SetConfigName("settings.dev")
	default:
		config.SetConfigName("settings.dev")
	}
	config.SetConfigType("toml")
	// Automatically refresh environment variables
	config.AutomaticEnv()

	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Println("failed to read configuration:", err.Error())
			os.Exit(1)
		}
	}

	return config
}
