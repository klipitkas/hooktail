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
	// The secret that is used for the deployment.
	Secret string `yaml:"secret,omitempty" json:"secret,omitempty"`
	// The TLS configuration.
	TLSConfig struct {
		// The path to the public key.
		PubKeyPath string `yaml:"public_key_path" json:"public_key_path"`
		// The path to the private key.
		PrivKeyPath string `yaml:"private_key_path" json:"private_key_path"`
	} `yaml:"tlsconfig" json:"tlsconfig"`
	// The list of deployments.
	Deployments []map[string]deployment.Deployment `yaml:"deployments,omitempty" json:"deployments,omitempty"`
}

// Parse parses a YAML file that contains the configuration
// and returns a Config struct as the result if the
// parsing is successful.
func Parse(configPath string) (Config, error) {

	// The default configuration is empty.
	config := Config{}

	// Read configuration
	b, err := ioutil.ReadFile(configPath)
	if err != nil {
		return config, fmt.Errorf("read yaml file %v: %v", configPath, err)
	}

	// Unmarshal to struct
	if err := yaml.Unmarshal(b, &config); err != nil {
		return config, fmt.Errorf("unmarshal yaml to struct: %v", err)
	}

	return config, nil
}
