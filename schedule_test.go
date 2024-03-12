package sidecarbackup

import (
	"reflect"
	"testing"
)

func TestNewScheduler(t *testing.T) {
	tests := []struct {
		name string
		want *Scheduler
	}{
		{
			name: "scheduler-test",
			want: &Scheduler{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewScheduler(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewScheduler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScheduler_Start(t *testing.T) {
	type fields struct {
		Error bool
	}
	type args struct {
		configFile string
		force bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "scheduler-test-good",
			fields: fields{
				Error: false,
			},
			args: args{
				configFile: "testdata/good.config.yaml",
				force: false,
			},
			want: true,
		},
		{
			name: "scheduler-test-bad",
			fields: fields{
				Error: false,
			},
			args: args{
				configFile: "testdata/bad.config.yaml",
				force: false,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scheduler{
				Error: tt.fields.Error,
			}
			if got := s.Start(tt.args.configFile, tt.args.force); got != tt.want {
				t.Errorf("Scheduler.Start() = %v, want %v", got, tt.want)
			}
		})
	}
}
