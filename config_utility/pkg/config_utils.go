package config_utils

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"net/url"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// LoadConfig populates config with struct defaults, config file at config_path and env variables (in that order)
func LoadConfig(config_type reflect.Type, config_path string) (*viper.Viper, error) {
	vconfig := viper.New()

	// Set all defaults as per config struct definition
	config_elems := config_type.Elem()
	for i := 0; i < config_elems.NumField(); i++ {
		field := config_elems.Field(i)
		tag := field.Tag
		key := tag.Get("mapstructure")
		defaultValue := tag.Get("default")
		if key == "" && defaultValue != "" {
			return nil, fmt.Errorf("loadConfig depends on the struct elem tags having a mapstructure tag for populating defaults. Found key %s with default value but no mapstructure key specified.", key)
		}
		log.Debugf("reading in env var: %s, default %s", key, defaultValue)
		vconfig.SetDefault(key, defaultValue)
		vconfig.BindEnv(key)
	}

	// Overide defaults with provided config file if any
	if config_path != "" {
		if _, err := os.Stat(config_path); err != nil {
			return nil, err
		}
		vconfig.AddConfigPath(filepath.Dir(config_path))
		config_extension := filepath.Ext(config_path) // with dot
		vconfig.SetConfigName(strings.TrimSuffix(filepath.Base(config_path), config_extension))
		vconfig.SetConfigType(config_extension[1:])
		if err := vconfig.MergeInConfig(); err != nil {
			return nil, err
		}
	}

	// Overide as per env var values set in shell
	vconfig.AutomaticEnv()

	// all settings loaded as per precedence, adjusting it as per config struct definition before unmarshalling
	config_map := vconfig.AllSettings()
	// all passwords converted to bytes
	for k, v := range config_map {
		if strings.HasSuffix(k, "_password") {
			vconfig.Set(k, []byte(fmt.Sprint(v)))
		} else if strings.HasSuffix(k, "_url") {
			url, err := url.Parse(fmt.Sprint(v))
			if err != nil {
				return nil, fmt.Errorf("Provided URL string %s couldn't be parsed to a valid URL, err: %v.", v, err)
			}
			vconfig.Set(k, url)
		}
	}

	return vconfig, nil
}
