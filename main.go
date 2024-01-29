package main

import (
	"fmt"
	"sync"
	"time"
)

func sayHello(wg *sync.WaitGroup, pingChannel chan string, pongChannel chan string) {
	
    for i := 0; i < 100; i++ {
            fmt.Printf("\nHello ")
		    pingChannel <- "World!"
		    <-pongChannel
            time.Sleep(1 * time.Second)
	}
	wg.Done()
}
func sayWorld(wg *sync.WaitGroup, pingChannel chan string, pongChannel chan string) {
        for i := 0; i < 100; i++ {
		v := <-pingChannel

        fmt.Printf(v)
		pongChannel <- "pong"
        }
	wg.Done()
}
func main() {
	var wg sync.WaitGroup

	pingChannel := make(chan string)
	pongChannel := make(chan string)
	
	wg.Add(1)
	go sayHello(&wg, pingChannel, pongChannel)

	wg.Add(1)
    go sayWorld(&wg, pingChannel, pongChannel)

        wg.Wait()

        fmt.Println("\nFinished")
}
