package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	common "github.com/klipitkas/hooktail/common"
	deployment "github.com/klipitkas/hooktail/deployment"

	"gopkg.in/yaml.v2"
)

// DeploymentRunner is the interface that implements
// the Run() method which is used for deployment.
type DeploymentRunner interface {
	Run(d deployment.Deployment) error
}

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

// The global configuration that will be used.
var config Config

func main() {

	// The path to the configuration file.
	var configPath string

	flag.StringVar(&configPath, "config", "config.yml", "The configuration file path.")
	flag.Parse()

	// Read configuration
	b, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalf("read %s: %v", configPath, err)
		return
	}

	// Unmarshal to struct
	if err := yaml.Unmarshal(b, &config); err != nil {
		log.Fatalf("parse yaml configuration details: %v", err)
		return
	}

	// The list of request handlers.
	http.HandleFunc("/", handleRequest)

	// Log the server start.
	log.Printf("Starting TLS server on port: %v", config.Port)

	// The server configuration.
	if err = http.ListenAndServeTLS(
		":"+strconv.Itoa(config.Port),
		config.TLSConfig.PubKeyPath,
		config.TLSConfig.PrivKeyPath,
		nil,
	); err != nil {
		log.Fatalf("listen on port %d: %v", config.Port, err)
	}

}

func handleRequest(w http.ResponseWriter, req *http.Request) {
	_, err := ioutil.ReadAll(req.Body)

	if err != nil {
		log.Printf("parsing request: %v", err)
		return
	}

	// request := Request{}

	// if err := json.Unmarshal(body, &request); err != nil {
	// 	log.Printf("cannot unmarshal string: %v", err)
	// 	return
	// }

	d := findMatchingDeployment("git@github.com:klipitkas/remove_greek_accents.git")

	if err = d.Run(); err != nil {
		log.Printf("run deployment: %v", err)
	}

	// requestDump, err := httputil.DumpRequest(req, true)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(requestDump))
}

func findMatchingDeployment(sshURL string) deployment.Deployment {
	for _, d := range config.Deployments {
		deployment := d["deployment"]
		if deployment.Repository == sshURL {
			return deployment
		}
	}

	return deployment.Deployment{}
}

func runDeployment(d deployment.Deployment) error {
	// Run the before script
	args := []string{d.BeforeScript}

	if err := common.ExecuteCommand("/bin/sh", d.User, "", args...); err != nil {
		return fmt.Errorf("execute before script: %v", err)
	}

	// Run the deployment
	args = []string{"remote", "update"}

	if err := common.ExecuteCommand("git", d.User, d.Path, args...); err != nil {
		return fmt.Errorf("execute remote update: %v", err)
	}

	args = []string{"reset", "--hard", "origin/" + d.Branch}

	if err := common.ExecuteCommand("git", d.User, d.Path, args...); err != nil {
		return fmt.Errorf("execute hard reset: %v", err)
	}

	// Run the after script
	args = []string{d.AfterScript}

	if err := common.ExecuteCommand("/bin/sh", d.User, "", args...); err != nil {
		return fmt.Errorf("execute before script: %v", err)
	}

	return nil
}
