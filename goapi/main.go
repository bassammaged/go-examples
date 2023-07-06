package main

import "log"

func main() {
	server := NewServer()
	err := server.Run()
	if err != nil {
		log.Println(err)
		return
	}

}
