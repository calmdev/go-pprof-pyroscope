package main

import (
	"fmt"
	"net/http"
	"net/http/pprof"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/{action}", pprof.Index)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)

	if err := http.ListenAndServe(":8081", mux); err != nil {
		panic(fmt.Errorf("Error when starting or running http server: %v", err))
	}
}
