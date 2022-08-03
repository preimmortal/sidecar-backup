package sidecarbackup

import (
	"testing"
)

func TestSql_Enabled(t *testing.T) {
	type fields struct {
		Name    string
		Source  string
		Dest    string
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
				Options: "",
				Enable:  true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := Sql{
				Name:    tt.fields.Name,
				Source:  tt.fields.Source,
				Dest:    tt.fields.Dest,
				Options: tt.fields.Options,
				Enable:  tt.fields.Enable,
			}
			if err := j.Execute(); (err != nil) != tt.wantErr {
				t.Errorf("Sql.Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
