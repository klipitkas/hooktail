package common

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"syscall"
)

// ExecuteCommand runs a specific shell command in the target system.
func ExecuteCommand(cmd string, username string, workDir string, args ...string) (string, error) {

	command := exec.Command(cmd, args...)

	var buf bytes.Buffer
	mw := io.MultiWriter(&buf)

	if username != "" {
		credentials, err := UserCredentialsFromUsername(username)
		if err != nil {
			return "", fmt.Errorf("get user id, gid from username %v: %v", username, err)
		}
		command.SysProcAttr = &syscall.SysProcAttr{}
		command.SysProcAttr.Credential = credentials
	}

	command.Env = os.Environ()
	command.Dir = workDir
	command.Stdout = mw
	command.Stderr = mw

	if err := command.Start(); err != nil {
		return "", fmt.Errorf("start command %q: %v", command, err)
	}

	if err := command.Wait(); err != nil {
		return "", fmt.Errorf("wait for command %q: %v", command, err)
	}

	return buf.String(), nil
}

// UserCredentialsFromUsername returns user credentials based on the
// provided username string.
func UserCredentialsFromUsername(username string) (*syscall.Credential, error) {
	uid, err := UIDFromUsername(username)
	if err != nil {
		return &syscall.Credential{}, err
	}
	gid, err := GIDFromUsername(username)
	if err != nil {
		return &syscall.Credential{}, err
	}
	return &syscall.Credential{Uid: uid, Gid: gid}, nil
}

// UIDFromUsername returns a users id (uid) based on his username.
func UIDFromUsername(username string) (uint32, error) {
	user, err := user.Lookup(username)
	if err != nil {
		return 0, fmt.Errorf("lookup user %v: %v", username, err)
	}
	uid, err := strconv.ParseUint(user.Uid, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("parse uid from string: %v", err)
	}
	return uint32(uid), nil
}

// GIDFromUsername returns a users group id (gid) based on his username.
func GIDFromUsername(username string) (uint32, error) {
	user, err := user.Lookup(username)
	if err != nil {
		return 0, fmt.Errorf("lookup user %v: %v", username, err)
	}
	gid, err := strconv.ParseUint(user.Gid, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("parse gid from string: %v", err)
	}
	return uint32(gid), nil
}

// Sha1Hmac returns the hmac sha1 hash of message m based on secret s.
func Sha1Hmac(message, secret string) string {
	h := hmac.New(sha1.New, []byte(secret))
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}
