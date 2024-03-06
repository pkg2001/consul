package linkconsul

import (
	"strconv"

	"github.com/hashicorp/consul/api"
)

func LinkConsul(ip string, port int, servername string, serverId string) error {
	//使用默认配置
	config := api.DefaultConfig()
	//配置连接consul地址
	config.Address = "127.0.0.1:8500"
	//实例化客户端
	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
		return err
	}
	check := &api.AgentServiceCheck{
		GRPC:                           ip + ":" + strconv.Itoa(port),
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "10s",
	}

	//健康检查，检查的是谁？检查的我们注册的服务
	Registration := api.AgentServiceRegistration{}
	Registration.Address = ip
	Registration.Port = port
	Registration.Name = servername

	Registration.ID = serverId
	Registration.Check = check

	err = client.Agent().ServiceRegister(&Registration)
	if err != nil {
		panic(err)
		return err
	}
	return nil
}
