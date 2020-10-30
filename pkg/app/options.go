/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2020/10/8 14:16
 */
package app

import (
	"github.com/nopsky/project-demo/pkg/transports"
	"github.com/nopsky/project-demo/pkg/transports/http"
	"github.com/nopsky/project-demo/pkg/transports/rpc"
)

type Options struct {
	name       string
	httpServer transports.IServer
	rpcServer  transports.IServer
}

type OptionFn func(*Options)

func newOptions(optFn ...OptionFn) Options {
	opts := Options{name: "default"}
	return opts
}

func Name(n string) OptionFn {
	return func(o *Options) {
		o.name = n
	}
}

func HttpServer(svr *http.Server) OptionFn {
	return func(o *Options) {
		o.httpServer = svr
	}
}

func RpcServer(svr *rpc.Server) OptionFn {
	return func(o *Options) {
		o.rpcServer = svr
	}
}
