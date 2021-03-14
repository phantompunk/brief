package command

import (
	"flag"
	"reflect"
	"testing"
)

func TestNewVersionCommand(t *testing.T) {
	tests := []struct {
		name string
		want *VersionCommand
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewVersionCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVersionCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVersionCommand_Run(t *testing.T) {
	type fields struct {
		short bool
		fs    *flag.FlagSet
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &VersionCommand{
				short: tt.fields.short,
				fs:    tt.fields.fs,
			}
			c.Run()
		})
	}
}
