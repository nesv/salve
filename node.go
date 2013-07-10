package main

import (
	"github.com/garyburd/redigo/redis"
)

/*
A Node represents a partition in the cluster ring. A node consists of one
Redis master, and N Redis slaves.
*/
type Node struct {
	master redis.Conn
	slaves []redis.Conn
}

/*
NewNode creates a new "node" to be added to the cluster.
*/
func NewNode(addr string) (n *Node, err error) {
	return
}

/*
All requests to the node should be pass through the Run function.

This function handles the round-robin load balancing on read operations, and
ensuring any write operations are only performed on the master host, in the
node.
*/
func (n *Node) Run(c []byte) (r []byte, err error) {
	return
}

func (n *Node) write([]byte) (r []byte, err error) {
	return
}

func (n *Node) read([]byte) (r []byte, err error) {
	return
}
