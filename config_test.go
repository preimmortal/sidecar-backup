package main

import "testing"

func Test_readConfig(t *testing.T) {
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
			name: "Test Good File",
			args: args{
				filename: "test/good.config.yaml",
			},
			wantErr: false,
		},
		{
			name: "Test Bad File",
			args: args{
				filename: "test/bad.config.yaml",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := readConfig(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("readConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
