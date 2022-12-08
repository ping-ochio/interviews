package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan string)

	var wg sync.WaitGroup
	wg.Add(2)
	go generaData(ch, &wg)
	go usaData(ch, &wg)
	wg.Wait()

}
func generaData(ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- "Hi"
}
func usaData(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	text := <-ch
	fmt.Println(text)
}
