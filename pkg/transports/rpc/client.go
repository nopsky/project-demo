/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2020/10/17 20:12
 */
package rpc

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

type Client struct {
	opts *ClientOptions
}

func NewClient(v *viper.Viper, options ...ClientOption) (c *Client, err error) {

	c = new(Client)

	c.opts, err = newClientOptions(v)
	if err != nil {
		return nil, err
	}

	for _, option := range options {
		option(c.opts)
	}

	return c, nil
}

func (c *Client) Dial(service string, grpcOptions ...grpc.DialOption) (*grpc.ClientConn, error) {

	if c.opts.Timeout == 0 {
		c.opts.Timeout = 30
	}

	ctx, cancel := context.WithTimeout(context.Background(), c.opts.Timeout*time.Second)
	defer cancel()

	server, ok := c.opts.Services[service]
	if !ok {
		return nil, fmt.Errorf("service `%s` is not found, please check config", service)
	}

	addr := fmt.Sprintf("%s:%d", server.Host, server.Port)

	conn, err := grpc.DialContext(ctx, addr, grpcOptions...)

	if err != nil {
		return nil, err
	}

	return conn, nil
}
