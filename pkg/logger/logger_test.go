package logger

import (
	"bytes"
	"testing"

	"github.com/nakabonne/golintui/pkg/config"
)

func TestNewLogger(t *testing.T) {
	type args struct {
		cfg *config.Config
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "general case",
			args: args{cfg: &config.Config{Debug: true}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			_ = NewLogger(tt.args.cfg, w)
			if gotW := w.String(); gotW == "" {
				t.Errorf("NewLogger() gotW = %v, want non-empty strign", gotW)
			}
		})
	}
}
