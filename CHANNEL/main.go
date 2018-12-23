package main

import (
	"fmt"
	"time"
)

func pinger(canal chan string) {
	for {
		canal <- "ping"
	}
}

func ponger(canal chan string) {
	for {
		canal <- "pong"
	}
}

func impressora(canal chan string) {
	for {
		msg := <-canal
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

//var orquestrador sync.WaitGroup

func main() {
	var canal chan string
	canal = make(chan string)
	//orquestrador.Add(2)
	go pinger(canal)
	go ponger(canal)
	go impressora(canal)
	//orquestrador.Wait()

	var entrada string
	fmt.Scanln(&entrada)
}
