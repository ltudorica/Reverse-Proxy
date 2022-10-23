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
	// Load the file; returns []byte
	f, err := os.ReadFile(filepath.Join("./config/", "config.yaml"))

	if err != nil {
		log.Fatal(err)
	}
	// Create an empty Config to be are target of unmarshalling
	var raw interface{}

	// Unmarshal our input YAML file into empty interface
	if err := yaml.Unmarshal(f, &raw); err != nil {
		log.Fatal(err)
	}

	// Use mapstructure to convert our interface{} to Car (var c)
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

// func PrintStuff() {
// 	// Load the file; returns []byte
// 	f, err := os.ReadFile(filepath.Join("./config/", "config.yaml"))

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// Create an empty Config to be are target of unmarshalling
// 	var c Config
// 	var raw interface{}

// 	// Unmarshal our input YAML file into empty interface
// 	if err := yaml.Unmarshal(f, &raw); err != nil {
// 		log.Fatal(err)
// 	}

// 	// Use mapstructure to convert our interface{} to Car (var c)
// 	decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{WeaklyTypedInput: true, Result: &c})
// 	if err := decoder.Decode(raw); err != nil {
// 		log.Fatal(err)
// 	}

// 	// Print out the new struct
// 	fmt.Printf("%+v\n", c)
// }
