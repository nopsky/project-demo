# project-demo

项目目录结构demo


protoc  --govalidators_out=. --go_out=. --go-grpc_opt=paths=source_relative --go_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false:. -I/Users/nopsky/go/pkg/mod/github.com/mwitkow/go-proto-validators@v0.3.2 -I/usr/local/include/google/protobuf -I. -I/Users/nopsky/go/pkg/mod ./*.proto