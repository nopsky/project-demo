/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2020/10/20 13:05
 */
package user

import (
	"context"

	"github.com/nopsky/project-demo/internal/user/entity"
)

type IUserService interface {
	RegisterUser(ctx context.Context, user *entity.UserEntity) (*entity.UserEntity, error)
}
