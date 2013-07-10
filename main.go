package main

import (
	"github.com/garyburd/redigo/redis"
	"log"
	"flag"
	"time"
)

/*
Command-line flags.
*/
var (
	configServer = flag.String("configsrv", "", "Address for the a dedicated Redis server for the purpose of storing configurations")
	configDatabase = flag.Int("configdb", 0, "The keyspace/DB # to store the cluster configuration in")
	verbose = flag.Bool("v", false, "Toggle verbosity")
	debugMode = flag.Bool("debug", false, "Enable debugging mode (lots of useless output)")
	httpAddr = flag.String("http", ":8000", "Port for the HTTP interface to be served from")
	connectionTimeout = flag.String("ctimeout", "2s", "Specify a connection timeout")
	readTimeout = flag.String("rtimeout", "100ms", "Specify a timeout for read operations")
	writeTimeout = flag.String("wtimeout", "100ms", "Specify a timeout for write operations")
	listenAddr = flag.String("listen", ":6379", "The listen address")
)

/*
Other things...
*/
var (
	Config redis.Conn
	ConnectionTimeout time.Duration
	ReadTimeout time.Duration
	WriteTimeout time.Duration
	Nodes map[string]redis.Conn
)

func init() {
	flag.Parse()

	var err error

	// Parse the provided timeouts to make sure we were passed sane values.
	ConnectionTimeout, err = time.ParseDuration(*connectionTimeout)
	if err != nil {
		log.Fatal(err)
	}

	ReadTimeout, err = time.ParseDuration(*readTimeout)
	if err != nil {
		log.Fatal(err)
	}

	WriteTimeout, err = time.ParseDuration(*writeTimeout)
	if err != nil {
		log.Fatal(err)
	}

	Nodes = make(map[string]redis.Conn)
}

func main() {
	var err error

	Config, err = redis.DialTimeout("tcp", *configServer, ConnectionTimeout, ReadTimeout, WriteTimeout)
	if err != nil {
		log.Println("error connecting to configuration server")
		log.Fatal(err)
	}
	
	go StartHttp(*httpAddr)
	log.Fatal(StartRouter(*listenAddr))
}
