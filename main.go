package main

import (
        "fmt"
	"sync"
)

func sayHello(wg *sync.WaitGroup) {
	wg.Done()
        for i := 0; i < 5; i++ {
                fmt.Println("Hello")
        }
}
func sayWorld(wg *sync.WaitGroup) {
	wg.Done()
        for i := 0; i < 5; i++ {
                fmt.Println("World")
        }
}
func main() {
	var wg sync.WaitGroup
	
	wg.Add(1)
	go sayHello(&wg)
	wg.Wait()

	wg.Add(1)
        go sayWorld(&wg)
        wg.Wait()

        fmt.Println("Finished")
}
