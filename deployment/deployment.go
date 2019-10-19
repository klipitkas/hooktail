package deployment

import (
	"errors"
	"fmt"
	"os"
	"os/user"
	"path"

	"github.com/klipitkas/hooktail/common"
	"github.com/klipitkas/hooktail/config"
	"github.com/klipitkas/hooktail/request"
)

// Deployment is the specific deployment configuration.
type Deployment struct {
	// The username of the user that will perform the deployment.
	User string `yaml:"user,omitempty" json:"user,omitempty"`
	// The repository of the project that will be deployed.
	Repository string `yaml:"repository,omitempty" json:"repository,omitempty"`
	// The branch that will be deployed.
	Branch string `yaml:"branch,omitempty" json:"branch,omitempty"`
	// The path where the deployment will take place.
	Path string `yaml:"path,omitempty" json:"path,omitempty"`
	// Any script that should be ran before the deployment.
	BeforeScript string `yaml:"before_script,omitempty" json:"before_script,omitempty"`
	// Any script that should be ran after the deployment.
	AfterScript string `yaml:"after_script,omitempty" json:"after_script,omitempty"`
}

// Validate validates a specified deployment configuration.
func Validate(d Deployment) error {
	// Basic validation checks
	if d.User == "" {
		return errors.New("invalid or empty user")
	}
	if d.Repository == "" {
		return errors.New("invalid repository")
	}
	if d.Branch == "" {
		return errors.New("invalid branch")
	}
	if d.Path == "" {
		return errors.New("invalid deployment path")
	}

	// System validation checks
	// if _, err := os.Stat("git"); os.IsNotExist(err) {
	// 	return fmt.Errorf("check git command existence: %v", err)
	// }

	if _, err := user.Lookup(d.User); err != nil {
		return fmt.Errorf("check user existence: %v", err)
	}

	if _, err := os.Stat(d.Path); os.IsNotExist(err) {
		return fmt.Errorf("check path existence: %v", err)
	}

	if _, err := os.Stat(path.Join(d.Path, ".git")); os.IsNotExist(err) {
		return fmt.Errorf("check .git inside path existence: %v", err)
	}

	// Checking before and after script existence
	if d.BeforeScript != "" {
		if _, err := os.Stat(d.BeforeScript); os.IsNotExist(err) {
			return fmt.Errorf("check before script existence: %v", err)
		}
	}

	if d.AfterScript != "" {
		if _, err := os.Stat(d.BeforeScript); os.IsNotExist(err) {
			return fmt.Errorf("check after script existence: %v", err)
		}
	}

	return nil
}

// Deploy executes a specific deployment configuration.
func Deploy(d Deployment) error {

	// Validate the deployment first.
	if err := Validate(d); err != nil {
		return fmt.Errorf("validate deployment: %v", err)
	}

	// Execute any script that needs to be executed before
	// the deployment.
	if err := runBefore(d); err != nil {
		return fmt.Errorf("before deployment: %v", err)
	}

	// Execute the deployment
	if err := run(d); err != nil {
		return fmt.Errorf("run deployment: %v", err)
	}

	// Execute any script that needs to be executed after
	// the deployment.
	if err := runAfter(d); err != nil {
		return fmt.Errorf("after deployment: %v", err)
	}

	return nil
}

// run executes the core deployment commands.
func run(d Deployment) error {
	// Run the deployment
	args := []string{"remote", "update"}
	if err := common.ExecuteCommand("git", d.User, d.Path, args...); err != nil {
		return fmt.Errorf("git remote update: %v", err)
	}

	args = []string{"reset", "--hard", "origin/" + d.Branch}
	if err := common.ExecuteCommand("git", d.User, d.Path, args...); err != nil {
		return fmt.Errorf("hard reset to %v branch: %v", d.Branch, err)
	}

	return nil
}

// runScript runs a bash deployment script.
func runScript(path string, user string) error {
	args := []string{path}
	if err := common.ExecuteCommand("/bin/sh", user, "", args...); err != nil {
		return fmt.Errorf("run script: %v", err)
	}
	return nil
}

// runBefore runs the script that is specified to be ran
// before the deployment takes place.
func runBefore(d Deployment) error {
	// Run the before script
	if err := runScript(d.BeforeScript, d.User); err != nil {
		return fmt.Errorf("before script: %v", err)
	}
	return nil
}

// runAfter runs the script that is specified to be ran
// after the deployment takes place.
func runAfter(d Deployment) error {
	// Run the after script
	if err := runScript(d.AfterScript, d.User); err != nil {
		return fmt.Errorf("after script: %v", err)
	}
	return nil
}

// FindMatching searches for a matching deployment in the YAML
// configuration file when parsing the request.
func FindMatching(config config.Config, req request.Request) Deployment {
	deployment := Deployment{}
	for _, deps := range config.Deployments {
		deployment := deps["deployment"]
		if deployment.Repository == req.Repository.SSHURL {
			return deployment
		}
	}
	return deployment
}
