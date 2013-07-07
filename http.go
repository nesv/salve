package main

import (
	"net/http"
	"log"
	"github.com/bmizerany/pat"
	"fmt"
)

func StartHttp(laddr string) error {
	// Create a new pattern muxer.
	pmux := pat.New()
	pmux.Post("/nodes/", http.HandlerFunc(httpAddNode))

	// Get the HTTP server listening, on the provided address.
	http.Handle("/", pmux)
	return http.ListenAndServe(laddr, nil)
}

/*
Adds a node to the cluster.
*/
func httpAddNode(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "tbd")
}
