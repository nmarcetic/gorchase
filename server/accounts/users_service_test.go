package accounts

import (
	"reflect"
	"testing"

	"github.com/nmarcetic/gorchase/pkg/logger"
)

func TestNewUsersService(t *testing.T) {
	type args struct {
		repo   UserRepository
		logger logger.Logger
	}
	tests := []struct {
		name string
		args args
		want Service
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUsersService(tt.args.repo, tt.args.logger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUsersService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_usersService_Register(t *testing.T) {
	type args struct {
		u User
	}
	tests := []struct {
		name    string
		svc     usersService
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.svc.Register(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("usersService.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("usersService.Register() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_usersService_Login(t *testing.T) {
	type args struct {
		email    string
		password string
	}
	tests := []struct {
		name    string
		svc     usersService
		args    args
		wantU   User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotU, err := tt.svc.Login(tt.args.email, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("usersService.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotU, tt.wantU) {
				t.Errorf("usersService.Login() = %v, want %v", gotU, tt.wantU)
			}
		})
	}
}

func Test_usersService_Get(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		svc     usersService
		args    args
		want    User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.svc.Get(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("usersService.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("usersService.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_usersService_GetByEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		svc     usersService
		args    args
		want    User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.svc.GetByEmail(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("usersService.GetByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("usersService.GetByEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_usersService_Update(t *testing.T) {
	type args struct {
		u User
	}
	tests := []struct {
		name    string
		svc     usersService
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.svc.Update(tt.args.u); (err != nil) != tt.wantErr {
				t.Errorf("usersService.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
