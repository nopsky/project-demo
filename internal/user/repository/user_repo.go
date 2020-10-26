/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2020/10/20 14:31
 */
package repository

import (
	"context"

	"github.com/nopsky/project-demo/internal/user"
	"github.com/nopsky/project-demo/internal/user/entity"
	"github.com/nopsky/project-demo/pkg/orm"
)

type userRepo struct {
	db *orm.Manager
}

func NewUserRepo(db *orm.Manager) user.IUserRepo {
	return &userRepo{db: db}
}

func (ur userRepo) GetUserByUsername(ctx context.Context, username string) (user *entity.UserEntity, err error) {
	return nil, nil
}

func (userRepo) Save(ctx context.Context, user *entity.UserEntity) error {
	return nil
}
