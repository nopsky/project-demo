/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2020/9/27 17:56
 */
package logger

type options struct {
	driver     string
	format     string
	output     string
	level      string
	filename   string
	maxSize    int
	maxBackups int
	maxAge     int
	compress   bool
}

type OptionFunc func(*options)

func newOptions(optsFunc ...OptionFunc) options {
	opt := options{
		driver:     "zap",
		output:     "console",
		level:      "debug",
		format:     "text",
		filename:   "./access.log",
		maxSize:    10,
		maxBackups: 5,
		maxAge:     30,
		compress:   false,
	}

	for _, o := range optsFunc {
		o(&opt)
	}

	return opt
}

func Driver(driver string) OptionFunc {
	return func(o *options) {
		o.driver = driver
	}
}

func Format(format string) OptionFunc {
	return func(o *options) {
		o.format = format
	}
}

func Output(output string) OptionFunc {
	return func(o *options) {
		o.output = output
	}
}

func Level(level string) OptionFunc {
	return func(o *options) {
		o.level = level
	}
}

func Filename(filename string) OptionFunc {
	return func(o *options) {
		o.filename = filename
	}
}

func MaxSize(maxSize int) OptionFunc {
	return func(o *options) {
		o.maxSize = maxSize
	}
}

func MaxBackups(maxBackups int) OptionFunc {
	return func(o *options) {
		o.maxBackups = maxBackups
	}
}

func MaxAge(maxAge int) OptionFunc {
	return func(o *options) {
		o.maxAge = maxAge
	}
}

func Compress(compress bool) OptionFunc {
	return func(o *options) {
		o.compress = compress
	}
}
