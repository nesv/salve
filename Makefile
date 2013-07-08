gopath = `pwd`

all: $(wildcard *.go)
	GOPATH=${gopath} go build -o=bin/salve-server $^

clean:
	rm -rf bin pkg src

deps:
	GOPATH=${gopath} go get github.com/garyburd/redigo/redis
	GOPATH=${gopath} go get github.com/bmizerany/pat
