package config

import (
	"fmt"

	"github.com/hashicorp/consul/api"
	"gopkg.in/yaml.v2"
)

type APIConfig struct {
	HTTPBindAddr string `yaml:"http_bind_addr" mapstructure:"http_bind_addr"`
	DBConn       string `yaml:"db_conn" mapstructure:"db_conn"`
}

func Get() (*APIConfig, error) {
	var a APIConfig
	client, err := api.NewClient(&api.Config{
		Address: "consul-server-bootstrap:8500",
	})
	if err != nil {
		return nil, err
	}

	kv := client.KV()

	// Lookup the pair
	pair, _, err := kv.Get("API_CONFIG", nil)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(pair.Value))
	err = yaml.Unmarshal(pair.Value, &a)
	fmt.Println(a)
	if err != nil {
		return nil, err
	}
	// agent := client.Agent()
	// err = agent.ServiceRegister(&api.AgentServiceRegistration{
	// 	Name: "api",
	// 	Tags: []string{"go"},
	// 	Port: 8080,
	// 	Check: &api.AgentServiceCheck{
	// 		CheckID:  "ping",
	// 		Name:     "Check web api",
	// 		HTTP:     "http://api:1234",
	// 		Interval: "5s",
	// 		Timeout:  "1s",
	// 	},
	// })

	return &a, err
}
