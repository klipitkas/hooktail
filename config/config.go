package config

import (
	"fmt"
	"io/ioutil"

	"github.com/klipitkas/hooktail/deployment"
	"gopkg.in/yaml.v2"
)

// Config contains all the configuration for the server.
type Config struct {
	// The port that the server will listen to.
	Port int `yaml:"port" json:"port"`
	// The list of deployments.
	Deployments []deployment.Deployment `yaml:"deployments,omitempty" json:"deployments,omitempty"`
}

// Parse parses a YAML file that contains the configuration
// and returns a Config struct as the result if the
// parsing is successful.
func Parse(config *Config, configPath string) error {
	// Read configuration
	b, err := ioutil.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("read yaml file %v: %v", configPath, err)
	}
	// Unmarshal to struct
	if err := yaml.Unmarshal(b, &config); err != nil {
		return fmt.Errorf("unmarshal yaml to struct: %v", err)
	}
	return nil
}
