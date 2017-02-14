package config

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"encoding/json"

	"github.com/stretchr/testify/assert"
	yaml "gopkg.in/yaml.v2"
)

func TestLoadYaml(t *testing.T) {
	yamlConfig := filepath.Join("testdata", "config.yml")

	yml, err := ioutil.ReadFile(yamlConfig)
	assert.Nil(t, err)

	config, err := LoadYaml(yml)
	assert.Nil(t, err)
	assert.NotNil(t, config)

	config, err = LoadYaml([]byte("This will totally not work"))
	assert.NotNil(t, config)
	assert.Error(t, err)
}

func TestLoadJSON(t *testing.T) {
	jsonConfig := filepath.Join("testdata", "config.json")

	j, err := ioutil.ReadFile(jsonConfig)
	assert.Nil(t, err)

	config, err := LoadJSON(j)
	assert.Nil(t, err)
	assert.NotNil(t, config)

	config, err = LoadJSON([]byte("This will totally not work"))
	assert.NotNil(t, config)
	assert.Error(t, err)
}

func TestConvertToJSON(t *testing.T) {
	heartbeat := Heartbeat{
		Timeout:   1,
		MaxMisses: 2,
	}

	config := Config{
		Heartbeats: &heartbeat,
	}

	actual, err := ConvertToJSON(config)
	assert.Nil(t, err)
	assert.NotNil(t, actual)

	expectedJSON, err := json.MarshalIndent(&config, "", "    ")
	assert.Nil(t, err)

	assert.JSONEq(t, string(expectedJSON), string(actual))
}

func TestConvertToYaml(t *testing.T) {
	heartbeat := Heartbeat{
		Timeout:   1,
		MaxMisses: 2,
	}

	config := Config{
		Heartbeats: &heartbeat,
	}

	actual, err := ConvertToYaml(config)
	assert.Nil(t, err)
	assert.NotNil(t, actual)

	expectedYaml, err := yaml.Marshal(config)
	assert.Nil(t, err)

	assert.Equal(t, string(expectedYaml), string(actual))
}
