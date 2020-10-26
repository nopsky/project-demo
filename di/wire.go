// +build wireinject

/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2020/10/26 19:38
 */
package di

import (
	"github.com/google/wire"
)

func MakeUseCase() (*UseCases, error) {
	panic(wire.Build(
		wire.Struct(new(UseCases), "*"),
		DBProviderSet,
		RepoProviderSet,
		ServiceProviderSet,
		UseCaseProviderSet,
	))
}
