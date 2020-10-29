/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2020/10/22 17:21
 */
package wrap

import (
	"github.com/nopsky/project-demo/pkg/errors"
)

type response struct {
	Code    int         `json:"code"`
	Stat    int         `json:"stat"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(data interface{}) interface{} {
	return response{
		Code:    0,
		Stat:    0,
		Message: "ok",
		Data:    data,
	}
}

func Fail(err error) interface{} {

	err = errors.Cause(err)

	code := -1

	e, ok := err.(*errors.Error)

	if ok {
		code = e.Code
	}

	return response{
		Code:    code,
		Stat:    1,
		Message: err.Error(),
		Data:    nil,
	}
}
