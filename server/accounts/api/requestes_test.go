package api

import (
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_getCurrentUser(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name    string
		args    args
		wantU   UserContext
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotU, err := getCurrentUser(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("getCurrentUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotU, tt.wantU) {
				t.Errorf("getCurrentUser() = %v, want %v", gotU, tt.wantU)
			}
		})
	}
}
