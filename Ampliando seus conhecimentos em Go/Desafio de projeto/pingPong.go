package main

import (
	"fmt"
	"time"
)

func main(){

	ping := make(chan bool)
	pong := make(chan bool)

	go func() {
		for {
			<-ping
			fmt.Println("ping")
			time.Sleep(1 * time.Second)
			pong <- true
		}
	}()
	go func () {
		for {
			<-pong
			fmt.Println("pong")
			time.Sleep(1 * time.Second)
			ping <- true
		}
	}()

	ping <- true

	select {}
}

