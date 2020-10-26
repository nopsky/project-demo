/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2020/10/17 10:29
 */
package rpc

import (
	"time"

	"github.com/spf13/viper"
)

type ServerOptions struct {
	Port int
	Host string
}

func newServiceOptions() *ServerOptions {
	return &ServerOptions{
		Port: 8080,
		Host: "0.0.0.0",
	}
}

type OptionFn func(*ServerOptions)

func ServerPort(port int) OptionFn {
	return func(o *ServerOptions) {
		o.Port = port
	}
}

type ClientOptions struct {
	Timeout  time.Duration
	Services map[string]ServerOptions
}

type ClientOption func(*ClientOptions)

func newClientOptions(v *viper.Viper) (opts *ClientOptions, err error) {
	opts = new(ClientOptions)

	if err = v.UnmarshalKey("grpc.client", opts); err != nil {
		return nil, err
	}

	return opts, nil
}

func WithTimeout(d time.Duration) ClientOption {
	return func(o *ClientOptions) {
		o.Timeout = d
	}
}
