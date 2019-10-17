package request

import (
	"encoding/json"
	"fmt"
)

// Request contains any needed request information.
type Request struct {
	// The repository struct.
	Repository struct {
		// The SSH URL of the repository.
		SSHURL string `yaml:"ssh_url" json:"ssh_url"`
	} `yaml:"repository" json:"repository"`
}

// Parse the request from the body to the struct.
func (r *Request) Parse(body string) error {
	if err := json.Unmarshal([]byte(body), &r); err != nil {
		return fmt.Errorf("unmarshal request body to json: %v", err)
	}
	return nil
}
