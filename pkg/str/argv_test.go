package str

import (
	"reflect"
	"testing"
)

func TestToArgv(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToArgv(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToArgv() = %v, want %v", got, tt.want)
			}
		})
	}
}
