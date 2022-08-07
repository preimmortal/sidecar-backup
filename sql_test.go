package sidecarbackup

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func mockExecCommand(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=CommandHelperProcess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}
	return cmd
}

func CommandHelperProcess(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}
	os.Exit(0)
}

func mockExistsCommand(target string) bool {
	return strings.Contains(target, "exist")
}

func mockRemoveCommand(target string) error {
	if strings.Contains(target, "error") {
		return nil
	}
	return fmt.Errorf("remove command error")
}

func TestSql_GetName(t *testing.T) {
	type fields struct {
		Name    string
		Source  string
		Dest    string
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
				Name: "sql-name",
			},
			want: "sql-name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			job := Sql{
				Name:    tt.fields.Name,
				Source:  tt.fields.Source,
				Dest:    tt.fields.Dest,
				Options: tt.fields.Options,
				Enable:  tt.fields.Enable,
			}
			if got := job.GetName(); got != tt.want {
				t.Errorf("Sql.GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSql_Enabled(t *testing.T) {
	type fields struct {
		Name    string
		Source  string
		Dest    string
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
			job := Sql{
				Name:    tt.fields.Name,
				Source:  tt.fields.Source,
				Dest:    tt.fields.Dest,
				Options: tt.fields.Options,
				Enable:  tt.fields.Enable,
			}
			if got := job.Enabled(); got != tt.want {
				t.Errorf("Sql.Enabled() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestSql_Execute(t *testing.T) {
	type fields struct {
		Name    string
		Source  string
		Dest    string
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
			name: "sql-test-pass",
			fields: fields{
				Name:    "sql-test-1",
				Source:  "/exists",
				Dest:    "/exists",
				Options: "",
				Enable:  true,
			},
			args: args{
				verbose: true,
			},
			wantErr: false,
		},
		{
			name: "sql-test-fail-1",
			fields: fields{
				Name:    "sql-test-fail-1",
				Source:  "/dne",
				Dest:    "/exists",
				Options: "",
				Enable:  true,
			},
			args: args{
				verbose: false,
			},
			wantErr: true,
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
			job := Sql{
				Name:    tt.fields.Name,
				Source:  tt.fields.Source,
				Dest:    tt.fields.Dest,
				Options: tt.fields.Options,
				Enable:  tt.fields.Enable,
			}
			if err := job.Execute(tt.args.verbose); (err != nil) != tt.wantErr {
				t.Errorf("Sql.Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
