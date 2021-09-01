package accounts

import "testing"

func TestUser_Validate(t *testing.T) {
	tests := []struct {
		name    string
		u       User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.u.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("User.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_isEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isEmail(tt.args.email); got != tt.want {
				t.Errorf("isEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}
