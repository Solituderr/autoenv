package conf

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

var Config Conf

type Conf struct {
	Service struct {
		Name string `yaml:"service_name"`
	}
	Etcd struct {
		Host string `yaml:"host"`
	}
}

func Init() {
	file, err := os.ReadFile("conf/config.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(file, &Config)
	fmt.Println(Config)
}
