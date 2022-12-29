package sidecarbackup

import (
	"os/exec"
	"testing"
)

func TestCommand_GetName(t *testing.T) {
	type fields struct {
		Name    string
		Command string
		Options string
		Enable  bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "test-name",
			fields: fields{
				Name: "command-name",
			},
			want: "command-name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			job := Command{
				Name:    tt.fields.Name,
				Command: tt.fields.Command,
				Options: tt.fields.Options,
				Enable:  tt.fields.Enable,
			}
			if got := job.GetName(); got != tt.want {
				t.Errorf("Command.GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommand_Enabled(t *testing.T) {
	type fields struct {
		Name    string
		Command string
		Options string
		Enable  bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "test-enabled-false",
			fields: fields{
				Enable: true,
			},
			want: true,
		},
		{
			name: "test-enabled-false",
			fields: fields{
				Enable: false,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			job := Command{
				Name:    tt.fields.Name,
				Command: tt.fields.Command,
				Options: tt.fields.Options,
				Enable:  tt.fields.Enable,
			}
			if got := job.Enabled(); got != tt.want {
				t.Errorf("Command.Enabled() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestCommand_Execute(t *testing.T) {
	type fields struct {
		Name    string
		Command string
		Options string
		Enable  bool
	}
	type args struct {
		verbose bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "command-test-pass",
			fields: fields{
				Name:    "command-test-1",
				Command: "ls",
				Options: "",
				Enable:  true,
			},
			args: args{
				verbose: true,
			},
			wantErr: false,
		},
		{
			name: "command-test-2",
			fields: fields{
				Name:    "command-test-2",
				Command: "touch /tmp/sidecar-backup.test",
				Options: "",
				Enable:  true,
			},
			args: args{
				verbose: false,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		// Mock our external function calls
		execCommand = mockExecCommand
		defer func() { execCommand = exec.Command }()
		existsCommand = mockExistsCommand
		defer func() { existsCommand = Exists }()
		removeCommand = mockRemoveCommand
		defer func() { removeCommand = Remove }()

		t.Run(tt.name, func(t *testing.T) {
			job := Command{
				Name:    tt.fields.Name,
				Command: tt.fields.Command,
				Options: tt.fields.Options,
				Enable:  tt.fields.Enable,
			}
			if err := job.Execute(tt.args.verbose); (err != nil) != tt.wantErr {
				t.Errorf("Command.Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
