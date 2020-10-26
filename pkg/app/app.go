/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2020/9/16 18:00
 */
package app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/nopsky/project-demo/pkg/config"
	"github.com/nopsky/project-demo/pkg/logger"
	"github.com/pkg/errors"
)

type Application struct {
	opts Options
}

func New(optFn ...OptionFn) (*Application, error) {
	opts := newOptions()

	app := &Application{
		opts: opts,
	}

	//读取配置文件
	if err := config.UnmarshalKey("application", &app.opts); err != nil {
		return nil, err
	}

	for _, fn := range optFn {
		fn(&app.opts)
	}

	return app, nil
}

func (a *Application) Start() error {
	if a.opts.httpServer != nil {
		if err := a.opts.httpServer.Start(); err != nil {
			return errors.Wrap(err, "start http server error")
		}
	}

	if a.opts.rpcServer != nil {
		if err := a.opts.rpcServer.Start(); err != nil {
			return errors.Wrap(err, "start grpc server error")
		}
	}

	return nil
}

func (a *Application) AwaitSignal() {
	c := make(chan os.Signal, 1)
	signal.Reset(syscall.SIGTERM, syscall.SIGINT)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	select {
	case s := <-c:
		logger.Info("receive a signal", s.String())
		if a.opts.httpServer != nil {
			if err := a.opts.httpServer.Stop(); err != nil {
				logger.Info("stop http server error", err.Error())
			}
		}

		if a.opts.rpcServer != nil {
			if err := a.opts.rpcServer.Stop(); err != nil {
				logger.Info("stop grpc server error", err.Error())
			}
		}

		os.Exit(0)
	}
}
