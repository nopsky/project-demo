/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2020/10/8 12:03
 */
package http

type Options struct {
	Port int
	Host string
	Mode string
}

type OptionFn func(*Options)

func newOptions() *Options {
	opts := &Options{
		Port: 80,
		Host: "0.0.0.0",
		Mode: "debug",
	}
	return opts
}

func Mode(mode string) OptionFn {
	return func(opts *Options) {
		opts.Mode = mode
	}
}
