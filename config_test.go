package sidecarbackup

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
				filename: "testdata/good.config.yaml",
			},
			wantErr: false,
		},
		{
			name: "Test Bad File",
			args: args{
				filename: "testdata/bad.config.yaml",
			},
			wantErr: true,
		},
		{
			name: "Test No File",
			args: args{
				filename: "testdata/dne.config.yaml",
			},
			wantErr: true,
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
