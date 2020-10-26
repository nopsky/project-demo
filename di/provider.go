/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2020/10/26 19:34
 */
package di

import (
	"github.com/google/wire"
	"github.com/nopsky/project-demo/internal/user/repository"
	"github.com/nopsky/project-demo/internal/user/service"
	"github.com/nopsky/project-demo/internal/user/usecase"
	"github.com/nopsky/project-demo/pkg/orm"
)

var (
	DBProviderSet      = wire.NewSet(orm.New)
	RepoProviderSet    = wire.NewSet(repository.NewUserRepo)
	ServiceProviderSet = wire.NewSet(service.NewUserService)
	UseCaseProviderSet = wire.NewSet(usecase.NewUserUseCase)
)
