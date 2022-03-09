package main

import (
	"reflect"
	"testing"
)

func Test_getConfig(t *testing.T) {
	tests := []struct {
		name    string
		want    *config
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getConfig()
			if (err != nil) != tt.wantErr {
				t.Errorf("getConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}
