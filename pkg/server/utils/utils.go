package utils

import (
	"log"
	"net"
	"net/url"
	"time"

	"bitbucket.org/ltudorica/App/pkg/lbPolicy"
	"bitbucket.org/ltudorica/App/pkg/parsing"
)

var urls []*url.URL

func Init_lbPolicy() lbPolicy.LB {

	var servers = parsing.GetHostsFromService(0)
	for i := 0; i < len(servers); i++ {
		url, _ := url.Parse(servers[i].URL)
		urls = append(urls, url)
	}

	lb, err := lbPolicy.New(urls)
	if err != nil {
		panic(err)
	}

	return lb
}

func IsAlive(url *url.URL) bool {
	conn, err := net.DialTimeout("tcp", url.Host, time.Minute)
	if err != nil {
		log.Printf("Unreachable to %v, error: %v", url.Host, err.Error())
		return false
	}
	defer conn.Close()
	return true
}

func BadGateway() string {

	return `
	<html>
	<head>
	<style> 
	h1 { text-align: center; }
	h3 { text-align: center; }
	</style>
	</head>
	<body>

	<h1><strong>502 Bad Gateway</strong></h1>
	<hr>
	<h3>BadGateway</h3>

	</body>
	<html>
	`

}

func Successful() string {

	return `
	<html>
	<head>
	<style>
	h1 { text-align: center; }
	h3 { text-align: center; }
	</style>
	</head>
	<body>

	<h1><strong>Success</strong></h1>
	<hr>
	<h3>Successfully Installed</h3>

	</body>
	<html>
	`

}
