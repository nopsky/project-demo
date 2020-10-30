/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2020/10/21 10:42
 */
package service

import (
	"context"

	"github.com/nopsky/project-demo/internal/user"
	"github.com/nopsky/project-demo/internal/user/entity"
	"github.com/nopsky/project-demo/pkg/ecode"
	"github.com/nopsky/project-demo/pkg/util/id"
	perrors "github.com/pkg/errors"
)

type userService struct {
	userRepo user.IUserRepo
}

func NewUserService(ur user.IUserRepo) user.IUserService {
	return &userService{userRepo: ur}
}

func (u *userService) RegisterUser(ctx context.Context, user *entity.UserEntity) (*entity.UserEntity, error) {
	item, err := u.userRepo.GetUserByUsername(ctx, user.Username)
	if err != nil {
		return nil, perrors.Wrap(err, "")
	}

	if item != nil {
		return nil, ecode.UserIsExists
	}

	user.ID = id.New()

	if err = u.userRepo.Save(ctx, user); err != nil {
		return nil, perrors.Wrap(err, "注册用户失败")
	}

	return user, nil

}

func (u *userService) Login(ctx context.Context, username, password string) (user entity.UserEntity, err error) {
	return
}

func (u *userService) Save(ctx context.Context, user entity.UserEntity) (entity.UserEntity, error) {
	return user, nil
}

func (u *userService) ListUser(ctx context.Context, companyID int64) ([]entity.UserEntity, error) {
	return nil, nil
}
