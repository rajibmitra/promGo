package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

        //Set up a connection to the server.
        
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
