package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

/**
 * @Author: dushuaihua
 * @Description:
 * @File:  consul_deregister
 * @Version: 1.0.0
 * @Date: 2021/9/12 9:34
 */

func main()  {
	// 1. 初始化 consul 配置
	consulConfig := api.DefaultConfig()

	//2.创建 consul 对象
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		fmt.Println("api.NewClient err:", err)
		return
	}

	//3. 注销服务
	consulClient.Agent().ServiceDeregister("bj38")


}