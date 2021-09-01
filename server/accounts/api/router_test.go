package api

import (
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/nmarcetic/gorchase/server/accounts"
)

func TestInitRouter(t *testing.T) {
	type args struct {
		r   *gin.RouterGroup
		svc accounts.Service
	}
	tests := []struct {
		name string
		args args
		want *gin.RouterGroup
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitRouter(tt.args.r, tt.args.svc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}
