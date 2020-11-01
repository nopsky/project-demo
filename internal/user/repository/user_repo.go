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
	"github.com/nopsky/project-demo/internal/user/repository/models"
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

	var userModel models.Users

	err = db.Where("username = ?", username).First(&userModel).Error

	if err != nil {
		if !orm.IsNotFoundError(err) {
			return nil, errors.WithError(err, ecode.GetUserByUsernameError)
		}
		return nil, nil
	}

	userEntity = &entity.UserEntity{
		ID:        userModel.ID,
		CompanyID: userModel.CompanyID,
		Username:  userModel.Username,
		Nickname:  userModel.Nickname,
		Role:      userModel.Role,
	}

	return userEntity, nil
}

func (ur userRepo) Save(ctx context.Context, user *entity.UserEntity) error {
	var userModel models.Users

	db, err := ur.db.NewOrm(ctx)

	if err != nil {
		return errors.WithError(err, ecode.DBNewOrmError)
	}

	userModel.ID = user.ID
	userModel.Username = user.Username

	if err := db.Save(userModel).Error; err != nil {
		return errors.WithError(err, ecode.SaveUserInfoError)
	}

	return nil
}
