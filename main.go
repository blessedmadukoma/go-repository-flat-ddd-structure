package main

import (
	"goRepositoryPattern/server"
)

func main() {
	server.HandleArgs()

	if err := server.Initialize(); err != nil {
		panic(err)
	}

	server.Run()
}
