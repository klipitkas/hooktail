package deployment_test

import (
	"testing"

	"github.com/klipitkas/hooktail/deployment"
)

func TestValidateDeployment(t *testing.T) {

	type args struct {
		dep deployment.Deployment
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"Test running a deployment without a user",
			args{
				dep: deployment.Deployment{
					User: "",
				},
			},
			"",
			true,
		},
		{
			"Test running a deployment without a repository",
			args{
				dep: deployment.Deployment{
					User: "test",
				},
			},
			"",
			true,
		},
		{
			"Test running a deployment without a branch",
			args{
				dep: deployment.Deployment{
					User:       "test",
					Repository: "test",
				},
			},
			"",
			true,
		},
		{
			"Test running a deployment without a path",
			args{
				dep: deployment.Deployment{
					User:       "test",
					Repository: "test",
					Branch:     "master",
				},
			},
			"",
			true,
		},
		{
			"Test running a deployment without a valid local user",
			args{
				dep: deployment.Deployment{
					User:       "test01828731827",
					Repository: "test",
					Branch:     "master",
					Path:       "/tmp",
				},
			},
			"",
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := deployment.Deploy(tt.args.dep)
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
		})
	}
}
