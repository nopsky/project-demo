/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2020/10/29 15:05
 */
package ecode

import "github.com/nopsky/project-demo/pkg/errors"

var (
	DBNewOrmError          = errors.New(100000, "获取数据库失败")
	GetUserByUsernameError = errors.New(200000, "根据用户名获取用户数据失败")
	UserIsExists           = errors.New(200001, "用户已存在")
	SaveUserInfoError      = errors.New(200002, "保存用户信息失败")
)
