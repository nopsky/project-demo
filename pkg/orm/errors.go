/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2020/10/30 18:04
 */
package orm

import (
	"gorm.io/gorm"
)

func IsNotFoundError(err error) bool {
	if err == gorm.ErrRecordNotFound {
		return true
	}

	return false
}
