/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2020/10/25 11:44
 */
package grpc

import (
	"github.com/nopsky/project-demo/di"
	"github.com/nopsky/project-demo/pkg/transports/rpc"
	pb "github.com/nopsky/project-demo/proto/user"
)

func Router(s *rpc.Server, d *di.UseCases) {
	pb.RegisterUserServiceServer(s.Server(), d.UserUseCase)
}
