package sidecarbackup

import (
	"testing"
)

func Test_verifyConfig(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "config-verify-test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			verifyConfig()
		})
	}
}

func Test_ReadConfig(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "config-test-good",
			args: args{
				filename: "testdata/good.config.yaml",
			},
			wantErr: false,
		},
		{
			name: "config-test-bad",
			args: args{
				filename: "testdata/bad.config.yaml",
			},
			wantErr: true,
		},
		{
			name: "config-test-no",
			args: args{
				filename: "testdata/dne.config.yaml",
			},
			wantErr: true,
		},
		{
			name: "config-test-neg-interval",
			args: args{
				filename: "testdata/neg.interval.config.yaml",
			},
			wantErr: false,
		},
		{
			name: "config-test-neg-worker",
			args: args{
				filename: "testdata/neg.workers.config.yaml",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ReadConfig(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("readConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
