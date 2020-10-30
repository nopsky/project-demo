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
)
