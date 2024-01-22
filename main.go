package main

import (
        "fmt"
	"sync"
)

func sayHello(wg *sync.WaitGroup, pingChannel chan string, pongChannel chan string) {
        for i := 0; i < 5; i++ {
                fmt.Println("Hello")
		pingChannel <- "World"
		<-pongChannel 
        }
	wg.Done()
}
func sayWorld(wg *sync.WaitGroup, pingChannel chan string, pongChannel chan string) {
        for i := 0; i < 5; i++ {
		v := <-pongChannel
                fmt.Println(v)
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

        fmt.Println("Finished")
}
