package main

import (
	"fmt"
	"sync"
)

var msg string

func updateMessage(s string) {
	msg = s
	printMesage()
}

func printMesage() {
	fmt.Println(msg)
}

func main() {
	var wg sync.WaitGroup

	//msg = "Hello, world!"

	/*		wg.Add(1)

			go updateMessage("Hello, universe!", &wg)
			printMesage()

			wg.Wait()

			wg.Add(1)

			go updateMessage("Hello, cosmos!", &wg)
			printMesage()

			wg.Wait()

			wg.Add(1)

			go updateMessage("Hello, world!", &wg)
			printMesage()

			wg.Wait()*/

	greet := []string{"Hello, universe!", "Hello, cosmos!", "Hello, world!"}

	wg.Add(len(greet))

	for _, val := range greet {

		go func(val string) {
			updateMessage(fmt.Sprintf("%s", val))
			//	printMesage()
			wg.Done()

		}(val)
	}
	wg.Wait()
}
