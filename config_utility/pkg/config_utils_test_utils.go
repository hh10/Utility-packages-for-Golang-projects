package config_utils

import (
	"net/url"
	"reflect"
)

type TestInnerConfig struct {
	SomeInnerString string `mapstructure:"SOME_INNER_STRING" default:"default_inner_str"`
	SomeInnerInt int32 `mapstructure:"SOME_INNER_INT" default:"6268"`
}

type TestConfig struct {
	TestInnerConfig  `mapstructure:",squash"`
	SomeURL      *url.URL `mapstructure:"SOME_URL" default:"http://127.0.0.1:30000"`
	SomePassword []byte `mapstructure:"SOME_PASSWORD" default:"password"`
	SomeString   string `mapstructure:"SOME_STRING" default:"default_str"`
	SomeInt      int32  `mapstructure:"SOME_INT" default:"30000"`
	SomeBool     bool   `mapstructure:"SOME_BOOL" default:"true"`
}

var DefaultTestConfig = TestConfig{
	TestInnerConfig: TestInnerConfig{
		SomeInnerString: "default_inner_str",
		SomeInnerInt: 6268,
	},
	SomeURL: &url.URL{Scheme: "http", Host: "127.0.0.1:30000"},
	SomePassword: []byte("password"),
	SomeString: "default_str",
	SomeInt: 30000,
	SomeBool: true,
}

func LoadTestConfig(override_config_path string) (*TestConfig, error) {
	config := &TestConfig{}
	vconfig, err := LoadConfig(reflect.TypeOf(config), override_config_path)
	if err != nil {
		return nil, err
	}
	if err = vconfig.Unmarshal(config); err != nil {
		return nil, err
	}
	return config, nil
}
