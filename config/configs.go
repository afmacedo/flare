// Package config implements config encoding and decoding.
package config

import (
	"encoding/json"

	yaml "gopkg.in/yaml.v2"
)

// Heartbeat represents hartbeat configuration between hosts.
type Heartbeat struct {
	// Timeout in seconds to wait between missing hartbeats
	// before giving up.
	Timeout uint32 `json:"timeout" yaml:"timeout"`

	// MaxMisses is the maximum heartbeat misses allowed.
	MaxMisses uint32 `json:"max_misses" yaml:"max_misses"`
}

// Config is the main configuration structure.
type Config struct {
	Heartbeats *Heartbeat `json:"heartbeat" yaml:"heartbeat"`
}

// LoadYaml parses a yaml config file.
func LoadYaml(config []byte) (*Config, error) {
	c := Config{}
	return &c, yaml.Unmarshal(config, &c)
}

// ConvertToYaml converts config into Yaml config.
func ConvertToYaml(config Config) ([]byte, error) {
	return yaml.Marshal(config)
}

// LoadJSON parses a JSON config file.
func LoadJSON(config []byte) (*Config, error) {
	c := Config{}
	return &c, json.Unmarshal(config, &c)
}

// ConvertToJSON converts config into JSON config.
func ConvertToJSON(config Config) ([]byte, error) {
	return json.MarshalIndent(config, "", "    ")
}
