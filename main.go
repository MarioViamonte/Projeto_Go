package main

import "github.com/MarioViamonte/Projeto/server"


func main(){
	server := server.NewServer()

	server.Run()
}