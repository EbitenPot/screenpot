// Created by EldersJavas(EldersJavas&gmail.com)

package test

import (
	"testing"
)

func TestSerJSON(t *testing.T) {
	type args struct {
		map1 map[string]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test", args: args{map1: map[string]string{
			"name":   "名1",
			"config": "{\"is:_ok\",\"ye:s\"}",
		}}, want: `{"name" :"名1","config":{"is_ok","yes"}}`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SerJSON(tt.args.map1); got != tt.want {
				t.Errorf("SerJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
