package main

import (
	"net"
	"net/http"
	"log"
	"github.com/bmizerany/pat"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func StartHttp(laddr string) error {
	// Create a new pattern muxer.
	pmux := pat.New()
	pmux.Post("/nodes/:addr", http.HandlerFunc(httpAddNode))
	pmux.Get("/nodes/", http.HandlerFunc(httpListNodes))

	// Get the HTTP server listening, on the provided address.
	http.Handle("/", pmux)
	log.Println("http: listening on", laddr)
	return http.ListenAndServe(laddr, nil)
}

/*
Adds a node to the cluster.
*/
func httpAddNode(w http.ResponseWriter, r *http.Request) {
	nodeAddr := r.URL.Query().Get(":addr")

	// First, try and connect to the node.
	if conn, err := redis.Dial("tcp", nodeAddr); err != nil {
		log.Printf("error while trying to connect to %q: %s", nodeAddr, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		Nodes[nodeAddr] = conn
	}
	
	// Add the new node address to an ordered list of masters.
	log.Printf("adding node %q to ring", nodeAddr)
	if _, err := Config.Do("RPUSH", "salve:masters", nodeAddr); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// Now, create a key for the node.
	if _, err := Config.Do("SET", fmt.Sprintf("salve:master.%s", nodeAddr), "alive"); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

/*
List all registered nodes.
*/
func httpListNodes(w http.ResponseWriter, r *http.Request) {
	nodes, err := redis.Strings(Config.Do("LRANGE", "salve:masters", "0", "-1"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	for _, n := range nodes {
		fmt.Fprintf(w, "%s\n", n)
	}
}
