/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2020/10/14 14:15
 */
package orm

import (
	"context"
	"sync"

	"gorm.io/gorm"
)

type DB struct {
	name string
	sm   sync.Map //管理当前DB下的事务对象 map[context.Context]*Session
	orm  *gorm.DB
}

func (d *DB) NewOrm(ctx context.Context) (*Session, error) {
	txCtx, ok := ctx.(*TxContext)
	if !ok {
		return NewSession(d.orm), nil
	}

	txID := txCtx.GetTxID()

	if txID == 0 {
		return NewSession(d.orm), nil
	}

	//由于当前的架构模式的下repository下的db为单例对象，所以需要判断是否已经有repo开启了事务
	session := tm.getSession(d.name, txID)
	if session != nil {
		return session, nil
	}

	session = NewSession(d.orm)

	err := session.WithContext(txCtx).Begin().Error
	if err != nil {
		return nil, err
	}

	err = tm.addSession(d.name, txID, session)

	if err != nil {
		return nil, err
	}

	return session, nil
}
