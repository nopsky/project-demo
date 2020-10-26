/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2020/10/14 19:21
 */
package orm

import (
	"context"
	"sync"
)

type useCaseFn func(ctx context.Context) (interface{}, error)

type transaction struct {
	ctx  *TxContext
	sm   map[string]*Session
	lock sync.RWMutex
}

func NewTransaction(ctx *TxContext) *transaction {
	return &transaction{
		ctx: ctx,
		sm:  make(map[string]*Session),
	}
}

func (tx *transaction) Commit() error {
	var err error
	for k, v := range tx.sm {
		err = v.Commit().Error
		if err != nil {
			return err
		}
		//当存在多个db的连接时，如果某个db提交成功了，就删除次db的事务信息
		//如果需要保持多个db事务的一致性，则需要采用分布式事务方式
		delete(tx.sm, k)
	}

	tm.close(tx.ctx)
	return nil
}

func (tx *transaction) Rollback() error {

	defer tm.close(tx.ctx)

	var err error
	for k, v := range tx.sm {
		err = v.Rollback().Error
		if err != nil {
			return err
		}
		delete(tx.sm, k)
	}

	return nil
}

//返回原始的ctx
func (tx *transaction) GetContext() context.Context {
	return tx.ctx.GetContext()
}

func (tx *transaction) getSession(dbName string) *Session {
	tx.lock.RLock()
	defer tx.lock.RUnlock()
	if val, ok := tx.sm[dbName]; ok {
		return val
	}

	return nil
}

func (tx *transaction) addSession(dbName string, sess *Session) {
	tx.lock.Lock()
	tx.sm[dbName] = sess
	tx.lock.Unlock()
}
