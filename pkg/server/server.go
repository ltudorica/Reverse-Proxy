package server

import (
	"log"
	"net/http"
	"strconv"

	"bitbucket.org/ltudorica/App/pkg/parsing"
	"bitbucket.org/ltudorica/App/pkg/server/routes"
)

func Start() {
	var PORT int = parsing.GetProxyPort()

	server := &http.Server{
		Addr:    ":" + strconv.Itoa(PORT),
		Handler: http.HandlerFunc(routes.HandleRequest),
	}

	log.Printf("Starting Server at PORT %s", strconv.Itoa(PORT))
	log.Fatal(server.ListenAndServe())

}
