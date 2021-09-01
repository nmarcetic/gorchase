package postgres

import (
	"reflect"
	"testing"

	"github.com/nmarcetic/gorchase/server/accounts"
	"gorm.io/gorm"
)

func TestUserDB_TableName(t *testing.T) {
	tests := []struct {
		name string
		u    UserDB
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.u.TableName(); got != tt.want {
				t.Errorf("UserDB.TableName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toDBObject(t *testing.T) {
	type args struct {
		u accounts.User
	}
	tests := []struct {
		name    string
		args    args
		want    UserDB
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := toDBObject(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("toDBObject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toDBObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fromDBObject(t *testing.T) {
	type args struct {
		udb     UserDB
		secrets bool
	}
	tests := []struct {
		name    string
		args    args
		want    accounts.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fromDBObject(tt.args.udb, tt.args.secrets)
			if (err != nil) != tt.wantErr {
				t.Errorf("fromDBObject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fromDBObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewUserRepository(t *testing.T) {
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want accounts.UserRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userRepository_Create(t *testing.T) {
	type args struct {
		u accounts.User
	}
	tests := []struct {
		name    string
		repo    userRepository
		args    args
		wantId  string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotId, err := tt.repo.Create(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("userRepository.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotId != tt.wantId {
				t.Errorf("userRepository.Create() = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}

func Test_userRepository_Get(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		repo    userRepository
		args    args
		wantU   accounts.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotU, err := tt.repo.Get(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("userRepository.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotU, tt.wantU) {
				t.Errorf("userRepository.Get() = %v, want %v", gotU, tt.wantU)
			}
		})
	}
}

func Test_userRepository_GetByEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		repo    userRepository
		args    args
		want    accounts.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.repo.GetByEmail(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("userRepository.GetByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userRepository.GetByEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userRepository_Update(t *testing.T) {
	type args struct {
		u accounts.User
	}
	tests := []struct {
		name    string
		repo    userRepository
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.repo.Update(tt.args.u); (err != nil) != tt.wantErr {
				t.Errorf("userRepository.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_userRepository_Delete(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		repo    userRepository
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.repo.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("userRepository.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
