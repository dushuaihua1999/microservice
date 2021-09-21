package main

import (
	"context"
	"day02/pb/pb"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"strconv"
)

/**
 * @Author: dushuaihua
 * @Description:
 * @File:  client
 * @Version: 1.0.0
 * @Date: 2021/9/11 12:05
 */

func main()  {
	// 1.初始化consul配置
	consulConfig := api.DefaultConfig()

	// 2.创建consulConfig对象
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		fmt.Println("api.NewClient err:", err)
		return
	}

	// 3. 服务发现
	services, _, err := consulClient.Health().Service("grpc And Consul","grpc",true, nil)

	// 4.
	addr := services[0].Service.Address + ":" + strconv.Itoa(services[0].Service.Port)

	grpcConn, _ := grpc.Dial(addr, grpc.WithInsecure())

	clt := pb.NewHelloClient(grpcConn)

	rsp,_ := clt.SayHello(context.TODO(), &pb.Person{
		Name: "dushuaihua",
		Age:  22,
	})
	fmt.Println(rsp)
}
