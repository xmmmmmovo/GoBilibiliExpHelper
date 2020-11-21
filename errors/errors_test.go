package errors

import (
	"reflect"
	"testing"
)

func TestErr_Error(t *testing.T) {
	type fields struct {
		Code int
		Msg  string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"1", fields{
			Code: 0,
			Msg:  "aa",
		}, "{\"Code\":0,\"Msg\":\"aa\"}"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Err{
				Code: tt.fields.Code,
				Msg:  tt.fields.Msg,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		code int
		msg  string
	}
	tests := []struct {
		name string
		args args
		want *Err
	}{
		{"1", args{
			code: 12450,
			msg:  "aaa",
		}, &Err{
			Code: 12450,
			Msg:  "aaa",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.code, tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
