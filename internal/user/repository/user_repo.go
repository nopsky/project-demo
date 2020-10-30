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
	"github.com/nopsky/project-demo/pkg/ecode"
	"github.com/nopsky/project-demo/pkg/errors"
	"github.com/nopsky/project-demo/pkg/orm"
)

type userRepo struct {
	db *orm.Manager
}

func NewUserRepo(db *orm.Manager) user.IUserRepo {
	return &userRepo{db: db}
}

func (ur userRepo) GetUserByUsername(ctx context.Context, username string) (userEntity *entity.UserEntity, err error) {
	db, err := ur.db.NewOrm(ctx)

	if err != nil {
		return nil, errors.WithError(err, ecode.DBNewOrmError)
	}

	var userModel string

	result := db.Where("username = ?", username).First(&userModel)

	if result.Error != nil {
		return nil, errors.WithError(result.Error, ecode.GetUserByUsernameError)
	}

	return userEntity, nil
}

func (userRepo) Save(ctx context.Context, user *entity.UserEntity) error {
	return nil
}
