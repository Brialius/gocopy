package main

import "testing"

func Test_validateArgs(t *testing.T) {
	tests := []struct {
		name    string
		a       Args
		wantErr bool
	}{
		{
			"OK",
			Args{"from", "to", 0, 0, true},
			false,
		},
		{
			"Fail 1",
			Args{"", "to", 0, 0, true},
			true,
		},
		{
			"Fail 2",
			Args{
				"from", "", 0, 0, true},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateArgs(tt.a); (err != nil) != tt.wantErr {
				t.Errorf("validateArgs() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
