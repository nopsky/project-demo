/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2020/10/22 17:21
 */
package wrap

import perrors "github.com/pkg/errors"

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
	return response{
		Code:    0,
		Stat:    1,
		Message: perrors.Cause(err).Error(),
		Data:    perrors.WithStack(err).Error(),
	}
}
