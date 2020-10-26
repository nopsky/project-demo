/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2020/10/14 16:05
 */
package orm

import (
	"gorm.io/gorm"
)

type Session struct {
	*gorm.DB
}

func NewSession(engine *gorm.DB) *Session {
	return &Session{engine.Session(&gorm.Session{WithConditions: false})}
}
