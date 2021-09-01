package accounts

import (
	"reflect"
	"testing"
)

func TestNewHasher(t *testing.T) {
	tests := []struct {
		name string
		want Hasher
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHasher(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHasher() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bcryptHasher_Hash(t *testing.T) {
	type args struct {
		pwd string
	}
	tests := []struct {
		name    string
		bh      *bcryptHasher
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.bh.Hash(tt.args.pwd)
			if (err != nil) != tt.wantErr {
				t.Errorf("bcryptHasher.Hash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("bcryptHasher.Hash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bcryptHasher_Compare(t *testing.T) {
	type args struct {
		plain  string
		hashed string
	}
	tests := []struct {
		name    string
		bh      *bcryptHasher
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.bh.Compare(tt.args.plain, tt.args.hashed); (err != nil) != tt.wantErr {
				t.Errorf("bcryptHasher.Compare() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
