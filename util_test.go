package sidecarbackup

import (
	"fmt"
	"io/fs"
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


func mockUtilStatCommand(target string) (fs.FileInfo, error) {
	if strings.Contains(target, "exist") {
		return nil, nil
	}
	return nil, fmt.Errorf("stat-error")
}

func mockUtilIsNotExistCommand(err error) bool {
	return err != nil
}

func mockUtilRemoveCommand(target string) error {
	if strings.Contains(target, "error") {
		return fmt.Errorf("remove-error")
	}
	return nil
}

func TestExists(t *testing.T) {
	utilStatCommand = mockUtilStatCommand
	defer func() { utilStatCommand = os.Stat }()
	utilIsNotExistCommand = mockUtilIsNotExistCommand
	defer func() { utilIsNotExistCommand = os.IsNotExist }()

	type args struct {
		target string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "exists-test-true",
			args: args{
				target: "exists",
			},
			want: true,
		},
		{
			name: "exists-test-false",
			args: args{
				target: "dne",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Exists(tt.args.target); got != tt.want {
				t.Errorf("Exists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	utilRemoveCommand = mockUtilRemoveCommand
	defer func() { utilRemoveCommand = os.Remove }()
	type args struct {
		target string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "remove-test-pass",
			args: args{
				target: "pass",
			},
			wantErr: false,
		},
		{
			name: "remove-test-fail",
			args: args{
				target: "error",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Remove(tt.args.target); (err != nil) != tt.wantErr {
				t.Errorf("Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
