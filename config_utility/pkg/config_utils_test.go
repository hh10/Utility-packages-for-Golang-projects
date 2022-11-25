package config_utils

import (
	"os"
	"reflect"
	"testing"
	"io/ioutil"
)

func TestDefaultConfigLoad(t *testing.T) {
	cfg, err := LoadTestConfig("")
	if err != nil {
		t.Fatal(err.Error())
	}
	reflect.DeepEqual(cfg, DefaultTestConfig)
}

func TestConfigOverrideByEnv(t *testing.T) {
	os.Setenv("SOME_STRING", "test_overides_str")
	os.Setenv("SOME_INNER_INT", "0000")
	DefaultTestConfig.SomeString = "test_overides_str"
	DefaultTestConfig.SomeInnerInt = 0000

	cfg, err := LoadTestConfig("")
	if err != nil {
		t.Fatal(err.Error())
	}
	reflect.DeepEqual(cfg, DefaultTestConfig)
}

func TestConfigOverrideByConfigFileEnv(t *testing.T) {
	// make tempfile with config to override in /tmp
	cfg_file, err := ioutil.TempFile("/tmp", "app_config_test_*.env")
	if err != nil {
		t.Fatal(err.Error())
	}
	defer os.Remove(cfg_file.Name())
	
	config_override := "SOME_INNER_STRING=\"inner_string_changed\"\nSOME_STRING='changed'\nSOME_BOOL=false\nSOME_INT=60000\nSOME_URL='http://localhost:3232'\n"
	if _, err = cfg_file.WriteString(config_override); err != nil {
		t.Fatal(err.Error())
	}
	os.Setenv("SOME_INNER_STRING", "inner_string_changed_again")
	DefaultTestConfig.SomeInnerString = "inner_string_changed_again"
	DefaultTestConfig.SomeString = "changed"
	DefaultTestConfig.SomeBool = false
	DefaultTestConfig.SomeInt = 0000
	DefaultTestConfig.SomeURL.Host = "localhost:3232"

	cfg, err := LoadTestConfig(cfg_file.Name())
	if err != nil {
		t.Fatal(err.Error())
	}
	reflect.DeepEqual(cfg, DefaultTestConfig)
}
