/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2020/9/16 17:04
 */
package transports

type IServer interface {
	Start() error
	Stop() error
}
