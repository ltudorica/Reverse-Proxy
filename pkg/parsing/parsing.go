package parsing

import (
	"log"
	"os"
	"path/filepath"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v3"
)

type Host struct {
	URL string `mapstructure:"URL"`
}

type Service struct {
	Name   string `mapstructure:"name"`
	Domain string `mapstructure:"domain"`
	Host   []Host `mapstructure:"hosts"`
}

type Listen struct {
	Adress    string `mapstructure:"address"`
	Port      int    `mapstructure:"port"`
	LB_Policy string `mapstructure:"lbPolicy"`
}

type Proxy struct {
	Listen  Listen    `mapstructure:"listen"`
	Service []Service `mapstructure:"services"`
}

type Config struct {
	Proxy Proxy `mapstructure:", squash"`
}

func GetProxy() Proxy {
	var config Config
	f, err := os.ReadFile(filepath.Join("./config/", "config.yaml"))

	if err != nil {
		log.Fatal(err)
	}
	var raw interface{}

	if err := yaml.Unmarshal(f, &raw); err != nil {
		log.Fatal(err)
	}

	decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{WeaklyTypedInput: true, Result: &config})
	if err := decoder.Decode(raw); err != nil {
		log.Fatal(err)
	}

	return config.Proxy
}

func GetProxyPort() int {
	return GetProxy().Listen.Port
}

func GetProxyAddress() string {
	return GetProxy().Listen.Adress
}

func GetLB_Policy() string {
	return GetProxy().Listen.LB_Policy
}

func GetService(i int) Service {
	return GetProxy().Service[i]
}

func GetServiceName(i int) string {
	return GetService(i).Name
}

func GetServiceDomain(i int) string {
	return GetService(i).Domain
}

func GetHostsFromService(i int) []Host {
	return GetService(i).Host
}
