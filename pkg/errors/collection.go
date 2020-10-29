/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2020/10/29 19:28
 */
package errors

import "fmt"

type collection struct {
	list map[int]string
}

var c = newCollection()

func newCollection() *collection {
	return &collection{list: make(map[int]string)}
}

func (c *collection) add(e Error) {
	if _, ok := c.list[e.Code]; ok {
		panic(fmt.Errorf("错误编码:%d 已经存在", e.Code))
	}

	c.list[e.Code] = e.Message
}

func (c *collection) Export() map[int]string {
	return c.list
}
