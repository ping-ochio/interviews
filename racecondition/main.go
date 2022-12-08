/*Here we generate an inconsistent condition by using goroutines
that access the same value at almost the same time.
To check for this condition, we have the "-race" flag,
which we can use as follows "go run -race main.go"
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	incremento := 0
	wg.Add(100)
	for i := 0; i < 100; i++ {

		go func() {
			p := incremento
			//"runtime.Gosched" gives the processor to another process
			runtime.Gosched()
			p++
			incremento = p
			fmt.Println(incremento)
			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println("El valor final es:", incremento)

}
