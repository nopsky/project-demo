/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2020/10/14 18:09
 */
package orm

import (
	"context"
)

type TxContext struct {
	context.Context
	txID TxID
}

func NewTxContext(ctx context.Context, txID TxID) *TxContext {
	return &TxContext{Context: ctx, txID: txID}
}

// return origin context
func (ctx *TxContext) GetContext() context.Context {
	return ctx.Context
}

func (ctx *TxContext) GetTxID() TxID {
	return ctx.txID
}

func (ctx *TxContext) clean() {
	ctx.txID = 0
}
