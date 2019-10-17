package main_test

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseRequest(t *testing.T) {

	type args struct {
		request string
	}

	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := "", fmt.Errorf("test")
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %+v, want = %+v", got, tt.want)
			}
		})
	}

}
