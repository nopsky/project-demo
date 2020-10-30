/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2020/10/16 09:59
 */
package id

import (
	"os"

	"github.com/bwmarrin/snowflake"
	"github.com/nopsky/project-demo/pkg/logger"
)

var node *snowflake.Node

func init() {
	n, err := snowflake.NewNode(1)
	if err != nil {
		logger.Panic("load snowflake err", err.Error())
		os.Exit(1)
	}
	node = n
}

func New() int64 {
	return node.Generate().Int64()
}
