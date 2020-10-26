/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2020/9/16 17:33
 */
package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nopsky/project-demo/pkg/config"
	"github.com/pkg/errors"
)

type Server struct {
	so         *Options
	router     *gin.Engine
	httpServer http.Server
}

type ServerRouter func(s *Server)

func New(optFn ...OptionFn) (*Server, error) {
	var err error

	//读取默认值
	s := &Server{
		so: newOptions(),
	}

	//读取配置文件
	if err = config.UnmarshalKey("http", s.so); err != nil {
		return nil, err
	}

	//更新配置
	for _, fn := range optFn {
		fn(s.so)
	}

	// 配置gin
	gin.SetMode(s.so.Mode)

	r := gin.New()

	//自定义logo
	r.Use(ginLogger())

	s.router = r

	return s, nil
}

func (s *Server) Start() error {

	if s.so.Port == 0 {
		return errors.New("指定监听的端口号")
	}

	if s.so.Host == "" {
		s.so.Host = "0.0.0.0"
	}

	addr := fmt.Sprintf("%s:%d", s.so.Host, s.so.Port)

	s.httpServer = http.Server{Addr: addr, Handler: s.router}

	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return
		}
	}()

	return nil
}

func (s *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5) // 平滑关闭,等待5秒钟处理

	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		return errors.Wrap(err, "shutdown http server error")
	}

	return nil
}
