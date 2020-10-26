/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2020/10/17 10:12
 */
package rpc

import (
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/nopsky/project-demo/pkg/config"
	"gitlab.jiayunhui.com/service/api/pkg/logger"
	"google.golang.org/grpc"
)

type Server struct {
	opts *ServerOptions
	grpc *grpc.Server
}

func New(optFn ...OptionFn) (*Server, error) {
	var err error
	s := &Server{
		opts: newServiceOptions(),
	}

	if err = config.UnmarshalKey("grpc.server", s.opts); err != nil {
		return nil, err
	}

	for _, fn := range optFn {
		fn(s.opts)
	}

	s.grpc = grpc.NewServer()

	return s, nil
}

func (s *Server) Server() *grpc.Server {
	return s.grpc
}

func (s *Server) Start() error {
	logger.Info("grpc server starting ...")
	if s.opts.Port == 0 {
		return errors.New("指定监听的端口号")
	}

	if s.opts.Host == "" {
		s.opts.Host = "0.0.0.0"
	}

	addr := fmt.Sprintf("%s:%d", s.opts.Host, s.opts.Port)

	go func() {
		lis, err := net.Listen("tcp", addr)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		if err = s.grpc.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	return nil
}

func (s *Server) Stop() error {
	logger.Info("grpc server stopping ...")
	s.grpc.GracefulStop()
	return nil
}
