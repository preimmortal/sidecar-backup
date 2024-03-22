package sidecarbackup

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	grsync "github.com/preimmortal/grsync"
	mocks "github.com/preimmortal/sidecar-backup/mocks"
)

func TestRsync_GetName(t *testing.T) {
	type fields struct {
		Name    string
		Source  string
		Dest    string
		Options grsync.RsyncOptions
		Enable  bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "name-test",
			fields: fields{
				Name: "rsync-name",
			},
			want: "rsync-name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			job := Rsync{
				Name:    tt.fields.Name,
				Source:  tt.fields.Source,
				Dest:    tt.fields.Dest,
				Options: tt.fields.Options,
				Enable:  tt.fields.Enable,
			}
			if got := job.GetName(); got != tt.want {
				t.Errorf("Rsync.GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRsync_Enabled(t *testing.T) {
	type fields struct {
		Name    string
		Source  string
		Dest    string
		Options grsync.RsyncOptions
		Enable  bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "rsync-enabled-test-true",
			fields: fields{
				Enable: true,
			},
			want: true,
		},
		{
			name: "rsync-enabled-test-false",
			fields: fields{
				Enable: false,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			job := Rsync{
				Name:    tt.fields.Name,
				Source:  tt.fields.Source,
				Dest:    tt.fields.Dest,
				Options: tt.fields.Options,
				Enable:  tt.fields.Enable,
			}
			if got := job.Enabled(); got != tt.want {
				t.Errorf("Rsync.Enabled() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRsync_Execute(t *testing.T) {
	existsCommand = mockExistsCommand
	defer func() { existsCommand = Exists }()
	type fields struct {
		Name    string
		Source  string
		Dest    string
		Options grsync.RsyncOptions
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
			name: "rsync-execute-test",
			fields: fields{
				Name:    "rsync-execute-test-1",
				Source:  "/exists/",
				Dest:    "/dest",
				Options: grsync.RsyncOptions{},
				Enable:  false,
			},
			args: args{
				verbose: true,
			},
			wantErr: false,
		},
		{
			name: "rsync-execute-test-dne",
			fields: fields{
				Name:    "rsync-execute-test-2",
				Source:  "/dne/",
				Dest:    "/dest",
				Options: grsync.RsyncOptions{},
				Enable:  false,
			},
			args: args{
				verbose: false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			job := Rsync{
				Name:    tt.fields.Name,
				Source:  tt.fields.Source,
				Dest:    tt.fields.Dest,
				Options: tt.fields.Options,
				Enable:  tt.fields.Enable,
			}
			if err := job.Execute(tt.args.verbose); (err != nil) != tt.wantErr {
				t.Errorf("Rsync.Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRsync_runTask(t *testing.T) {
	type fields struct {
		Name    string
		Source  string
		Dest    string
		Options grsync.RsyncOptions
		Enable  bool
	}
	type args struct {
		verbose bool
		task    Task
	}
	type modifiers struct {
		runErr   bool
		runDelay bool
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		modifiers modifiers
		wantErr   bool
	}{
		{
			name:   "rsync-runtask-test",
			fields: fields{},
			args: args{
				verbose: true,
				task:    nil,
			},
			modifiers: modifiers{
				runErr:   false,
				runDelay: true,
			},
			wantErr: false,
		},
		{
			name:   "rsync-runtask-test",
			fields: fields{},
			args: args{
				verbose: true,
				task:    nil,
			},
			modifiers: modifiers{
				runErr:   false,
				runDelay: false,
			},
			wantErr: false,
		},
		{
			name:   "rsync-runtask-test-2",
			fields: fields{},
			args: args{
				verbose: true,
				task:    nil,
			},
			modifiers: modifiers{
				runErr:   true,
				runDelay: false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mocks.NewMockTask(ctrl)

			if tt.modifiers.runErr {
				m.EXPECT().
					Run().
					Return(fmt.Errorf("run-error"))
			}

			if tt.modifiers.runDelay {
				m.EXPECT().
					Run().
					DoAndReturn(func() error {
						time.Sleep(10010 * time.Millisecond)
						return nil
					}).
					AnyTimes()
			} else {
				m.EXPECT().
					Run().
					Return(nil).
					AnyTimes()
			}

			m.EXPECT().
				State().
				Return(grsync.State{}).
				AnyTimes()

			m.EXPECT().
				Log().
				Return(grsync.Log{}).
				AnyTimes()

			tt.args.task = m

			job := Rsync{
				Name:    tt.fields.Name,
				Source:  tt.fields.Source,
				Dest:    tt.fields.Dest,
				Options: tt.fields.Options,
				Enable:  tt.fields.Enable,
			}

			if err := job.runTask(tt.args.verbose, tt.args.task); (err != nil) != tt.wantErr {
				t.Errorf("Rsync.runTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
