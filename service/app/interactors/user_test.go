package interactors

import (
	"backend/app/inputs"
	"backend/app/outputs"
	repositories "backend/app/repositories/interface"
	"reflect"
	"testing"
)

func TestUserListInteractor_Handle(t *testing.T) {
	type fields struct {
		userRepo repositories.IUserRepository
	}
	type args struct {
		in *inputs.UserListInputData
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		want1   *outputs.UserListOutputData
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &UserListInteractor{
				userRepo: tt.fields.userRepo,
			}
			got, got1, err := r.Handle(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Handle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Handle() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Handle() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestUserGetInteractor_Handle(t *testing.T) {
	type fields struct {
		userRepo repositories.IUserRepository
	}
	type args struct {
		in *inputs.UserGetInputData
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		want1   *outputs.UserGetOutputData
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &UserGetInteractor{
				userRepo: tt.fields.userRepo,
			}
			got, got1, err := r.Handle(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Handle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Handle() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Handle() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestUserUpdateInteractor_Handle(t *testing.T) {
	type fields struct {
		userRepo repositories.IUserRepository
	}
	type args struct {
		in *inputs.UserUpdateInputData
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		want1   *outputs.UserUpdateOutputData
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &UserUpdateInteractor{
				userRepo: tt.fields.userRepo,
			}
			got, got1, err := r.Handle(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Handle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Handle() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Handle() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestUserDeleteInteractor_Handle(t *testing.T) {
	type fields struct {
		userRepo repositories.IUserRepository
	}
	type args struct {
		in *inputs.UserDeleteInputData
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &UserDeleteInteractor{
				userRepo: tt.fields.userRepo,
			}
			got, err := r.Handle(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Handle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Handle() got = %v, want %v", got, tt.want)
			}
		})
	}
}
