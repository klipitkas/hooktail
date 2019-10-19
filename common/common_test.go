package common_test

import (
	"reflect"
	"syscall"
	"testing"

	"github.com/klipitkas/hooktail/common"
)

func TestExecuteCommand(t *testing.T) {

	type args struct {
		command     string
		commandArgs []string
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"Test running a simple echo command",
			args{
				command:     "echo",
				commandArgs: []string{"Hello world"},
			},
			"Hello world\n",
			false,
		},
		{
			"Test running a more complex echo command",
			args{
				command:     "echo",
				commandArgs: []string{"Hello world\n\n"},
			},
			"Hello world\n\n\n",
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := common.ExecuteCommand(tt.args.command, "", "", tt.args.commandArgs...)
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %+v (%T), want = %+v (%T)", got, got, tt.want, tt.want)
			}
		})
	}
}

func TestUIDFromUsername(t *testing.T) {

	type args struct {
		username string
	}

	tests := []struct {
		name    string
		args    args
		want    uint32
		wantErr bool
	}{
		{
			"Get the root uid",
			args{
				username: "root",
			},
			0,
			false,
		},
		{
			"Get a random user uid should fail",
			args{
				username: "x-root-0851",
			},
			0,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := common.UIDFromUsername(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %+v (%T), want = %+v (%T)", got, got, tt.want, tt.want)
			}
		})
	}
}

func TestGIDFromUsername(t *testing.T) {

	type args struct {
		username string
	}

	tests := []struct {
		name    string
		args    args
		want    uint32
		wantErr bool
	}{
		{
			"Get the root gid",
			args{
				username: "root",
			},
			0,
			false,
		},
		{
			"Get a random user gid should fail",
			args{
				username: "x-root-0851",
			},
			0,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := common.GIDFromUsername(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %+v (%T), want = %+v (%T)", got, got, tt.want, tt.want)
			}
		})
	}
}

func TestUserCredentialsFromUsername(t *testing.T) {

	type args struct {
		username string
	}

	tests := []struct {
		name    string
		args    args
		want    syscall.Credential
		wantErr bool
	}{
		{
			"Get the root user credentials",
			args{
				username: "root",
			},
			syscall.Credential{Uid: 0, Gid: 0},
			false,
		},
		{
			"Get a random user credentials should fail",
			args{
				username: "x-root-0851",
			},
			syscall.Credential{},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := common.UserCredentialsFromUsername(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, &tt.want) {
				t.Errorf("got = %+v (%T), want = %+v (%T)", got, got, tt.want, tt.want)
			}
		})
	}
}
