package util

import (
	consulapi "github.com/hashicorp/consul/api"
	"log"
)

var consulClient *consulapi.Client

func init()  {
	config := consulapi.DefaultConfig()
	config.Address= "127.0.0.1:8500"
	client, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}
	consulClient = client
}

func RegService() {


	reg := consulapi.AgentServiceRegistration{}
	reg.ID = "user.service"
	reg.Name= "user.service"
	reg.Address= "192.168.3.5"
	reg.Port = 8080
	reg.Tags = []string{"primary"}

	check := consulapi.AgentServiceCheck{}
	check.Interval = "5s"
	check.HTTP= "http://192.168.3.5:8080/health"
	reg.Check = &check

	consulClient.Agent().ServiceRegister(&reg)
}

func UnRegService(){
	consulClient.Agent().ServiceDeregister("user.service")
}