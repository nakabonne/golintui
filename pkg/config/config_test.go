package config

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		name           string
		version        string
		commit         string
		date           string
		buildSource    string
		executable     string
		openCommandEnv string
		cfgDir         string
		debuggingFlag  bool
	}
	tests := []struct {
		name string
		args args
		want *Config
	}{
		// TODO: Add test cases.
		{
			name: "general case",
			args: args{
				name:           "name1",
				version:        "version1",
				commit:         "commit1",
				date:           "date1",
				buildSource:    "buildSource1",
				executable:     "executable1",
				openCommandEnv: "openCommandEnv1",
				cfgDir:         "dir1",
				debuggingFlag:  true,
			},
			want: &Config{
				Name:           "name1",
				Version:        "version1",
				Commit:         "commit1",
				BuildDate:      "date1",
				BuildSource:    "buildSource1",
				OpenCommandEnv: "openCommandEnv1",
				CfgDir:         "dir1",
				Executable:     "executable1",
				Debug:          true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.name, tt.args.version, tt.args.commit, tt.args.date, tt.args.buildSource, tt.args.executable, tt.args.openCommandEnv, tt.args.cfgDir, tt.args.debuggingFlag); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
