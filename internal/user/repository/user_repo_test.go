/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2020/11/1 11:02
 */
package repository

import (
	"context"
	"reflect"
	"testing"

	"github.com/nopsky/project-demo/internal/user/entity"
	"github.com/nopsky/project-demo/pkg/config"
	"github.com/nopsky/project-demo/pkg/orm"
)

func getDB() (*orm.Manager, error) {
	var configFile = "../../../config/config.yml"

	//加载配置文件
	config.Load(configFile)

	return orm.New()

}
func Test_userRepo_GetUserByUsername(t *testing.T) {

	db, err := getDB()

	if err != nil {
		t.Fatal("test new user repo err", err)
	}

	type fields struct {
		db *orm.Manager
	}
	type args struct {
		ctx      context.Context
		username string
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantUserEntity *entity.UserEntity
		wantErr        bool
	}{
		// TODO: Add test cases.
		{
			name:           "get user by username",
			fields:         fields{db: db},
			args:           args{context.Background(), "test"},
			wantUserEntity: nil,
			wantErr:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ur := userRepo{
				db: tt.fields.db,
			}
			gotUserEntity, err := ur.GetUserByUsername(tt.args.ctx, tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotUserEntity, tt.wantUserEntity) {
				t.Errorf("GetUserByUsername() gotUserEntity = %v, want %v", gotUserEntity, tt.wantUserEntity)
			}
		})
	}
}
