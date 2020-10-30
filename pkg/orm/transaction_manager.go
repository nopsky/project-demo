/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2020/10/13 20:37
 */
package orm

import (
	"context"
	"fmt"

	"github.com/nopsky/project-demo/pkg/util/id"
)

//自定义事务ID
type TxID int64

type TransactionManager struct {
	rm map[TxID]*transaction
}

func NewTransactionManager() *TransactionManager {
	return &TransactionManager{rm: make(map[TxID]*transaction)}
}

func (tm *TransactionManager) BeginTx(ctx context.Context) (context.Context, *transaction) {
	//生成事务id
	txID := tm.generateTxID()
	//生成事务上下文
	newCtx := NewTxContext(ctx, txID)
	//生成事务控制器
	tx := NewTransaction(newCtx)

	tm.rm[txID] = tx

	return newCtx, tx
}

func (tm *TransactionManager) generateTxID() TxID {
	return TxID(id.New())
}

func (tm *TransactionManager) getSession(dbName string, txID TxID) *Session {
	tx, ok := tm.rm[txID]

	if !ok {
		return nil
	}

	return tx.getSession(dbName)
}

func (tm *TransactionManager) addSession(dbName string, txID TxID, sess *Session) error {
	tx, ok := tm.rm[txID]
	if !ok {
		return fmt.Errorf("当前事务ID:`%d`不存在", txID)
	}

	tx.addSession(dbName, sess)

	return nil
}

func (tm *TransactionManager) close(ctx *TxContext) {
	ctx.clean()
	delete(tm.rm, ctx.GetTxID())
}

var tm = NewTransactionManager()

// 开启事务
func BeginTx(ctx context.Context) (context.Context, *transaction) {
	return tm.BeginTx(ctx)
}

func Transaction(ctx context.Context, fn useCaseFn) (interface{}, error) {
	ctx, tx := BeginTx(ctx)
	resp, err := fn(ctx)
	if err != nil {
		err = tx.Rollback()
		return nil, err
	}

	err = tx.Commit()

	if err != nil {
		return nil, err
	}
	return resp, nil
}
