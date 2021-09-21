package main

import (
	"context"
	"day02/pb/pb"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"net"
)

/**
 * @Author: dushuaihua
 * @Description:
 * @File:  server
 * @Version: 1.0.0
 * @Date: 2021/9/11 12:01
 */

type Children struct {}

func (c *Children) SayHello(ctx context.Context, p *pb.Person) (*pb.Person, error) {
	p.Name = "hello " + p.Name
	fmt.Println()
	return p, nil
}

func main() {
	//把grpc对象服务，注册到consul上
	//1.初始化consul
	consulConfig := api.DefaultConfig()

	//创建consul对象
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		fmt.Println("api.NewClient err:", err)
		return
	}
	
	//3. 告诉consul 即将注册的服务的配置
	reg := api.AgentServiceRegistration{
		ID:"bj38",
		Tags:[]string{"grpc", "consul"},
		Name:"grpc And Consul",
		Address: "127.0.0.1",
		Port: 8800,
		Check: &api.AgentServiceCheck{
			CheckID: "consul grpc test",
			TCP: "127.0.0.1:8800",
			Timeout: "1s",
			Interval: "5s",
		},
	}
	//4. 注册 grpc 服务到 consul上
	consulClient.Agent().ServiceRegister(&reg)
	grpcSrv := grpc.NewServer()

	pb.RegisterHelloServer(grpcSrv, new(Children))

	listener, err := net.Listen("tcp", "127.0.0.1:8800")
	if err != nil {
		fmt.Println("listener err :", err)
	}
	defer listener.Close()

	grpcSrv.Serve(listener)
}