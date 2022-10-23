package routes

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"

	"bitbucket.org/ltudorica/App/pkg/lbPolicy"
	"bitbucket.org/ltudorica/App/pkg/server/utils"
)

var mu sync.Mutex

var lb lbPolicy.LB = utils.Init_lbPolicy()

func HandleRequest(res http.ResponseWriter, req *http.Request) {

	log.Printf("%s: %s%s", req.Method, req.Host, req.URL.Path)

	if lb.CountServers() == 0 {
		res.WriteHeader(200)
		res.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(res, utils.Successful())

		return
	}

	var counter int
	backend := *lb.GetServers()

	for i := 0; i < lb.CountServers(); i++ {
		if !utils.IsAlive(backend[i].URL) {
			counter++
		}
	}

	if counter == len(*lb.GetServers()) {
		res.WriteHeader(502)
		res.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(res, utils.BadGateway())
		return
	}

	mu.Lock()
	currentServer := lb.Next()

	if currentServer.GetIsDead() {
		currentServer = lb.Next()
	}

	log.Printf("Redirecting to Proxy URL %s", currentServer.URL)

	var url *url.URL = currentServer.URL

	mu.Unlock()
	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("%v is dead", url)
		currentServer.SetDead(true)
		HandleRequest(w, r)
	}

	req.URL.Host = url.Host
	req.URL.Scheme = url.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = url.Host

	proxy.ServeHTTP(res, req)
}
