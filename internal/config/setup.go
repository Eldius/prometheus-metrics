package config

import (
	"fmt"
	"github.com/eldius/prometheus-metrics/internal/logging"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	serviceName = "prometheus-metrics"
)

// Setup configures app parameters
func Setup(cfgFile string) error {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".mock-server-go" (without extension).
		viper.AddConfigPath(filepath.Join(home, ".basic-auth"))
		viper.AddConfigPath(filepath.Join(home))
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName("application")
	}

	SetDefaults()
	MapEnvVars()

	viper.AutomaticEnv() // read in environment variables that match

	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using config file:", viper.ConfigFileUsed())
	}

	if err := logging.SetupLogs(serviceName, viper.GetBool("debug")); err != nil {
		err = fmt.Errorf("failed to configure logs: %w", err)
		return err
	}

	return nil
}

// SetDefaults sets default configuration values
func SetDefaults() {
	viper.SetDefault("auth.key", "1234567890")
}

// MapEnvVars sets up environment variables mapping
func MapEnvVars() {
	viper.SetEnvPrefix("metrics")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}
