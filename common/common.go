package common

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"syscall"
)

// ExecuteCommand runs a specific shell command in the target system.
func ExecuteCommand(cmd string, username string, workDir string, args ...string) error {

	var buf bytes.Buffer
	mw := io.MultiWriter(&buf)

	uid, gid, err := UserIDsFromUsername(username)
	if err != nil {
		return fmt.Errorf("getting user details: %v", err)
	}

	command := exec.Command(cmd, args...)
	command.SysProcAttr = &syscall.SysProcAttr{}
	command.SysProcAttr.Credential = &syscall.Credential{Uid: uid, Gid: gid}
	command.Env = os.Environ()
	command.Dir = workDir
	command.Stdout = mw
	command.Stderr = mw

	if err := command.Start(); err != nil {
		return fmt.Errorf("start command %q: %v", command, err)
	}

	if err := command.Wait(); err != nil {
		return fmt.Errorf("wait for command %q: %v", command, err)
	}

	return nil
}

// UserIDsFromUsername returns the user id and group id for
// a specific system user.
func UserIDsFromUsername(username string) (uint32, uint32, error) {

	user, err := user.Lookup(username)
	if err != nil {
		return 0, 0, fmt.Errorf("get user from username %s: %v", username, err)
	}

	uid, err := strconv.ParseUint(user.Uid, 10, 32)
	if err != nil {
		return 0, 0, fmt.Errorf("convert uid to uint64: %v", err)
	}

	gid, err := strconv.ParseUint(user.Gid, 10, 32)
	if err != nil {
		return 0, 0, fmt.Errorf("convert gid to uint64: %v", err)
	}

	return uint32(uid), uint32(gid), nil
}
