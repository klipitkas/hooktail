package config_test

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	config "github.com/klipitkas/hooktail/config"
)

func TestParse(t *testing.T) {

	type args struct {
		configPath string
	}

	tests := []struct {
		name    string
		args    args
		content string
		want    config.Config
		wantErr bool
	}{
		{
			"Parse valid yaml file configuration",
			args{
				configPath: "/tmp/valid.yml",
			},
			`port: 5042`,
			config.Config{
				Port: 5042,
			},
			false,
		},
		{
			"Parse invalid yaml file should fail",
			args{
				configPath: "/tmp/invalid.yml",
			},
			`port: -5042-`,
			config.Config{},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := []byte(tt.content)
			if err := ioutil.WriteFile(tt.args.configPath, b, 0644); err != nil {
				t.Errorf("writefile failed %v", err)
				return
			}
			got, err := config.Parse(tt.args.configPath)
			os.Remove(tt.args.configPath)
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
