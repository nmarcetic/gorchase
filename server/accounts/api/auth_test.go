package api

import (
	"reflect"
	"testing"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/nmarcetic/gorchase/server/accounts"
)

func TestAuthMiddleware(t *testing.T) {
	type args struct {
		svc accounts.Service
	}
	tests := []struct {
		name string
		args args
		want *jwt.GinJWTMiddleware
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AuthMiddleware(tt.args.svc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthMiddleware() = %v, want %v", got, tt.want)
			}
		})
	}
}
