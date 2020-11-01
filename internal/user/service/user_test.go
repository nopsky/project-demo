/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2020/11/1 12:02
 */
package service

import (
	"context"
	"reflect"
	"testing"

	"github.com/nopsky/project-demo/internal/user"
	"github.com/nopsky/project-demo/internal/user/entity"
	"github.com/nopsky/project-demo/mocks"
)

func Test_userService_RegisterUser(t *testing.T) {

	ctx := context.Background()

	userRepo := new(mocks.IUserRepo)

	type fields struct {
		userRepo user.IUserRepo
	}
	type args struct {
		ctx  context.Context
		user *entity.UserEntity
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.UserEntity
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "register user service",
			fields: fields{userRepo: userRepo},
			args: args{ctx, &entity.UserEntity{
				ID:        1,
				CompanyID: 0,
				Username:  "test",
				Password:  "12346",
				Nickname:  "test",
				Role:      "admin",
			}},
			want: &entity.UserEntity{
				ID:        1,
				CompanyID: 0,
				Username:  "test",
				Password:  "12346",
				Nickname:  "test",
				Role:      "admin",
			},
			wantErr: false,
		},
	}

	userRepo.On("GetUserByUsername", ctx, "test").Return(nil, nil)

	userRepo.On("Save", ctx, &entity.UserEntity{
		ID:        1,
		CompanyID: 0,
		Username:  "test",
		Password:  "12346",
		Nickname:  "test",
		Role:      "admin",
	}).Return(nil)

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			u := &userService{
				userRepo: tt.fields.userRepo,
			}
			got, err := u.RegisterUser(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("RegisterUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegisterUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}
