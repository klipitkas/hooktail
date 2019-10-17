package common

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"syscall"
)

func ExecuteCommand(cmd string, username string, workDir string, args ...string) error {

	command := exec.Command(cmd, args...)

	u, err := user.Lookup(username)

	if err != nil {
		return fmt.Errorf("get user from username %s: %v", username, err)
	}

	uidInt, err := strconv.ParseUint(u.Uid, 10, 32)
	gidInt, err := strconv.ParseUint(u.Gid, 10, 32)

	command.SysProcAttr = &syscall.SysProcAttr{}
	command.SysProcAttr.Credential = &syscall.Credential{
		Uid: uint32(uidInt),
		Gid: uint32(gidInt),
	}

	log.Printf("Running command %s as user %s with id %v and gid %v.", cmd, username, uidInt, gidInt)

	command.Env = os.Environ()
	command.Dir = workDir

	var buf bytes.Buffer
	mw := io.MultiWriter(&buf)

	command.Stdout = mw
	command.Stderr = mw

	if err := command.Start(); err != nil {
		return fmt.Errorf("start command %q: %v", command, err)
	}

	if err := command.Wait(); err != nil {
		return fmt.Errorf("wait for command %q: %v", command, err)
	}

	log.Printf("%v", buf.String())

	return nil
}
