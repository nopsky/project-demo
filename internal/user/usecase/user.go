/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2020/10/22 15:20
 */
package usecase

import (
	"context"

	"github.com/nopsky/project-demo/internal/user"
	"github.com/nopsky/project-demo/internal/user/entity"
	"github.com/nopsky/project-demo/pkg/util/id"
	pb "github.com/nopsky/project-demo/proto/user"
	perrors "github.com/pkg/errors"
)

type userUseCase struct {
	userService user.IUserService
}

func NewUserUseCase(us user.IUserService) user.IUserUseCase {
	return &userUseCase{userService: us}
}

func (uuc userUseCase) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (resp *pb.RegisterUserResponse, err error) {
	user := &entity.UserEntity{
		ID:        id.New(),
		CompanyID: 0,
		Username:  req.Username,
		Password:  req.Password,
		Nickname:  "",
		Role:      "",
	}

	user, err = uuc.userService.RegisterUser(ctx, user)

	if err != nil {
		return nil, perrors.Wrap(err, "注册用户")
	}

	resp = new(pb.RegisterUserResponse)

	resp.Username = user.Username
	resp.Id = user.ID

	return resp, nil
}
