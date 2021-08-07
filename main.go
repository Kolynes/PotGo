package main

import "github.com/Kolynes/PotGo/server"

func main() {
	s := server.Server{
		Host: "127.0.0.1",
		Port: 8000,
	}
	s.Initiate()
}
