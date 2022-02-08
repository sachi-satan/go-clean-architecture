package interactors

import (
	"backend/app/inputs"
	"backend/app/outputs"
	repositories "backend/app/repositories/interface"
	"backend/app/services"
	"reflect"
	"testing"
)

func TestAuthSignUpInteractor_Handle(t *testing.T) {
	type fields struct {
		userRepo     repositories.IUserRepository
		jwtService   *services.Jwt
		mysqlService *services.MySql
	}
	type args struct {
		in *inputs.AuthSignUpInputData
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		want1   *outputs.AuthSignUpOutputData
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &AuthSignUpInteractor{
				userRepo:     tt.fields.userRepo,
				jwtService:   tt.fields.jwtService,
				mysqlService: tt.fields.mysqlService,
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

func TestAuthSingInInteractor_Handle(t *testing.T) {
	type fields struct {
		userRepo   repositories.IUserRepository
		jwtService *services.Jwt
	}
	type args struct {
		in *inputs.AuthSignInInputData
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		want1   *outputs.AuthSignInOutputData
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &AuthSingInInteractor{
				userRepo:   tt.fields.userRepo,
				jwtService: tt.fields.jwtService,
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
