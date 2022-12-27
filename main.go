package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"

	"github.com/valyala/gorpc"
)

type Storage struct {
	pool map[string]interface{}
}

// Storage of type string
func (ac *Storage) AddString(args [2]string) {
	ac.pool[args[0]] = args[1]
}
func (ac *Storage) GetString(key string) string {
	return ac.pool[key].(string)
}

// Storage of type int
func (ac *Storage) AddInt(args [2]string) {
	ac.pool[args[0]], _ = strconv.Atoi(args[1])
}
func (ac *Storage) GetInt(key string) int {
	return ac.pool[key].(int)
}

// Storage of type bytes
func (ac *Storage) AddBytes(args [2][]byte) {
	ac.pool[string(args[0])] = args[1]
}
func (ac *Storage) GetBytes(key string) []byte {
	return ac.pool[key].([]byte)
}

var (
	bindPort = 3000
)

func main() {
	flag.IntVar(&bindPort, "port", 3000, "Specify the port server binds to.")
	flag.Parse()

	d := gorpc.NewDispatcher()
	st := new(Storage)
	st.pool = make(map[string]interface{})

	d.AddService("Storage", st)
	tcpServer := gorpc.NewTCPServer(
		fmt.Sprintf(":%d", bindPort),
		d.NewHandlerFunc(),
	)

	fmt.Printf("[INFO] LilyBox Server Binded on: %d\n", bindPort)
	err := tcpServer.Serve()
	if err != nil {
		log.Fatalln(err)
	}
	tcpServer.Stop()
}
