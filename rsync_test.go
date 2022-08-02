package main

import (
	"testing"
)

func TestRsync_IsEnabled(t *testing.T) {
	type fields struct {
		Name    string
		Source  string
		Dest    string
		Ignored string
		Options string
		Enable  bool
		Result  string
		Error   error
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Enabled Test Positive",
			fields: fields{
				Name:    "Test",
				Source:  "/testSource",
				Dest:    "/testDest",
				Options: "",
				Enable:  true,
			},
			want: true,
		},
		{
			name: "Enabled Test Negative",
			fields: fields{
				Name:    "Test",
				Source:  "/testSource",
				Dest:    "/testDest",
				Options: "",
				Enable:  false,
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
				Ignored: tt.fields.Ignored,
				Options: tt.fields.Options,
				Enable:  tt.fields.Enable,
				Result:  tt.fields.Result,
				Error:   tt.fields.Error,
			}
			if got := job.Enabled(); got != tt.want {
				t.Errorf("Rsync.IsEnabled() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRsync_Execute(t *testing.T) {
	type fields struct {
		Name    string
		Source  string
		Dest    string
		Ignored string
		Options string
		Enable  bool
		Result  string
		Error   error
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Execute Test",
			fields: fields{
				Name:    "Test",
				Source:  "/testSource",
				Dest:    "/testDest",
				Ignored: "*ignored*",
				Options: "",
				Enable:  true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := Rsync{
				Name:    tt.fields.Name,
				Source:  tt.fields.Source,
				Dest:    tt.fields.Dest,
				Ignored: tt.fields.Ignored,
				Options: tt.fields.Options,
				Enable:  tt.fields.Enable,
				Result:  tt.fields.Result,
				Error:   tt.fields.Error,
			}
			if err := j.Execute(); (err != nil) != tt.wantErr {
				t.Errorf("Rsync.Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
