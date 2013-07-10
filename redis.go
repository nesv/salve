package main

type RedisCommand struct {
	Command string
	Key string
	Rest []byte
}

func ParseRequest(m []byte) (r *RedisCommand, err error) {
	return
}
