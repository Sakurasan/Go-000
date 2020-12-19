package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config") // name of config file (without extension)
	//viper.AddConfigPath("/etc/appname/")   // path to look for the config file in
	//viper.AddConfigPath("$HOME/.appname")  // call multiple times to add many search paths
	viper.AddConfigPath(".")    // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		logrus.WithError(err).Error("application configuration'initialization is failed ")
	}
}

// GetString returns the value associated with the key as a string.
func GetString(key string) string {
	return viper.GetString(key)
}

// GetInt returns the value associated with the key as an integer.
func GetInt(key string) int {
	return viper.GetInt(key)
}

// GetBool returns the value associated with the key as a boolean.
func GetBool(key string) bool {
	return viper.GetBool(key)
}

func GetStringMapString(key string) map[string]string {
	return viper.GetStringMapString(key)
}
