package models

import (
	"testing"
)

func TestUser_ValidatePassword(t *testing.T) {
	user := NewUser()
	err := user.SetNewUser("username", "username@example.com", "secret12345", "User", "Hello")
	if err != nil {
		return
	}

	type args struct {
		password string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "correct_password",
			args: args{
				password: "secret12345",
			},
			wantErr: false,
		},
		{
			name: "incorrect_password",
			args: args{
				password: "secret",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := user
			if err := r.ValidatePassword(tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("ValidatePassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
