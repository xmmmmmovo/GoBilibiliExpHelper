package utils

import "testing"

func TestPrintEnd(t *testing.T) {
	type args struct {
		tag string
	}
	tests := []struct {
		name string
		args args
	}{
		{"1", args{tag: "开始"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintStart(tt.args.tag)
		})
	}
}

func TestPrintStart(t *testing.T) {
	type args struct {
		tag string
	}
	tests := []struct {
		name string
		args args
	}{
		{"1", args{tag: "开始"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintEnd(tt.args.tag)
		})
	}
}
