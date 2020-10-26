/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2020/10/20 13:05
 */
package user

import (
	"context"

	"github.com/nopsky/project-demo/proto/user"
)

type IUserUseCase interface {
	RegisterUser(ctx context.Context, req *user.RegisterUserRequest) (resp *user.RegisterUserResponse, err error)
}
