/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2020/10/11 13:29
 */
package orm

type Options struct {
	Databases map[string]*DBOptions
}

type DBOptions struct {
	Dns          string
	MaxIdleConns int
	MaxOpenConns int
	Mode         string
}

type OptionFn func(*Options)

func newOptions() *Options {
	return &Options{Databases: make(map[string]*DBOptions)}
}
