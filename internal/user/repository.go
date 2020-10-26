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

type IUserRepo interface {
	GetUserByUsername(ctx context.Context, username string) (user *entity.UserEntity, err error)
	Save(ctx context.Context, user *entity.UserEntity) error
}
