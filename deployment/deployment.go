package deployment

import "errors"

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

// Run is the method that runs a specific deployment.
func (d *Deployment) Run() error {

	if d.User == "" {
		return errors.New("invalid user")
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

	return nil
}
