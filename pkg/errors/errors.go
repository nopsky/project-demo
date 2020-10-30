/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2020/10/28 10:47
 */
package errors

import perrors "github.com/pkg/errors"

type Error struct {
	Code     int
	Message  string
	otherErr error
}

func New(code int, message string) *Error {
	e := &Error{
		Code:    code,
		Message: message,
	}

	c.add(e)

	return e
}

func (e Error) Error() string {
	return e.Message
}

func (e Error) OtherError() error {
	return e.otherErr
}

func WithError(err error, e *Error) *Error {
	e.otherErr = err
	return e
}

func Wrap(err error, message string) error {
	return perrors.Wrap(err, message)
}

func Wrapf(err error, format string, args ...interface{}) error {
	return perrors.Wrapf(err, format, args...)
}

func Cause(err error) error {
	return perrors.Cause(err)
}

func Is(err error, target error) bool {
	return perrors.Is(err, target)
}

func As(err error, target interface{}) bool {
	return perrors.As(err, target)
}
