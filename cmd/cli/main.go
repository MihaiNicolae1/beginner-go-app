package main

import (
	"github.com/MihaiNicolae1/beginner-go-app/network"
	"log"
)

func main() {
	log.Println("Awesome CLI v0.01")
	network.Ping("127.0.0.1")
}
