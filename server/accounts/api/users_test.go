package api

import (
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/nmarcetic/gorchase/server/accounts"
)

func TestNewHandler(t *testing.T) {
	type args struct {
		svc accounts.Service
	}
	tests := []struct {
		name string
		args args
		want UserHandler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHandler(tt.args.svc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userHandler_Singup(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		h    userHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.Singup(tt.args.c)
		})
	}
}

func Test_userHandler_Get(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		h    userHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.Get(tt.args.c)
		})
	}
}
