package main

import (
	"github.com/Kong/go-pdk"
	"github.com/Kong/go-pdk/server"
)

const Version = "1.0.0"
const Priority = 1

type Config struct {
}

func New() interface{} {
	return &Config{}
}

func main() {
	server.StartServer(New, Version, Priority)
}

func (conf *MyConfig) Access(kong *pdk.PDK) {
}

func (conf *MyConfig) Response(kong *pdk.PDK) {
}
