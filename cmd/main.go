/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2020/10/20 13:01
 */
package main

import (
	"flag"
	"log"

	"github.com/nopsky/project-demo/di"
	"github.com/nopsky/project-demo/internal/user/api/grpc"
	"github.com/nopsky/project-demo/internal/user/api/rest"
	"github.com/nopsky/project-demo/pkg/app"
	"github.com/nopsky/project-demo/pkg/config"
	"github.com/nopsky/project-demo/pkg/transports/http"
	"github.com/nopsky/project-demo/pkg/transports/rpc"
)

var configFile = flag.String("f", "config.yml", "set config file which viper will loading.")

func main() {

	flag.Parse()

	//加载配置文件
	config.Load(*configFile)

	hs, err := http.New()

	if err != nil {
		log.Fatal("#1", err)
	}

	if err != nil {
		log.Fatal("#2", err)
	}

	rs, err := rpc.New()

	if err != nil {
		log.Fatal("#3", err)
	}

	ucs, err := di.MakeUseCase()

	if err != nil {
		log.Fatal("#4", err)
	}

	setProvider(hs, rs, ucs)

	a, err := app.New(app.HttpServer(hs), app.RpcServer(rs))

	if err != nil {
		log.Fatal("#5", err)
	}

	if err := a.Start(); err != nil {
		panic(err)
	}

	a.AwaitSignal()

}

func setProvider(s *http.Server, r *rpc.Server, ucs *di.UseCases) {
	//rest
	rest.Router(s, ucs)

	//grpc
	grpc.Router(r, ucs)
}
