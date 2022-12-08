/*In this case we remove the race condition using "sync.Mutex, Lock() and Unlock".
Therefore, no goroutine overlaps the execution time of another goroutine.
*/

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var l sync.Mutex

func main() {
	incremento := 0
	wg.Add(100)
	for i := 0; i < 100; i++ {

		go func() {
			l.Lock()
			p := incremento
			p++
			incremento = p
			fmt.Println(incremento)
			l.Unlock()
			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println("El valor final es:", incremento)

}
