package api

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_respErr(t *testing.T) {
	type args struct {
		c    *gin.Context
		code int
		err  error
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			respErr(tt.args.c, tt.args.code, tt.args.err)
		})
	}
}

func Test_respOK(t *testing.T) {
	type args struct {
		c    *gin.Context
		data interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			respOK(tt.args.c, tt.args.data)
		})
	}
}
