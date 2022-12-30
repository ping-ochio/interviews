package main

import "fmt"

func main() {
	// ---- Simple channel - bidrectional
	c := make(chan int)

	// send
	go func() {

		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()

	// recibe
	for v := range c {

		fmt.Println(v)
	}
	fmt.Println("The end")

	//------- Channels receive only, send only and select(works like switch)

	odd := make(chan int)
	even := make(chan int)
	exit := make(chan bool)

	// send
	go send(odd, even, exit)

	// recibe
	recibe(odd, even, exit)

	fmt.Println("End")

}

func send(o, e chan<- int, ex chan<- bool) {

	for j := 0; j < 100; j++ {
		if j%2 == 0 {
			o <- j

		} else {

			e <- j
		}
	}
	close(ex)
}

func recibe(o, e <-chan int, ex <-chan bool) {
	for {
		select {
		case v := <-o:
			fmt.Println("From odd channel:", v)
		case v := <-e:
			fmt.Println("From even channel:", v)
		case i, ok := <-ex:
			if !ok {
				fmt.Println("next:", i)
				return
			} else {
				fmt.Println("Exiting:", i)

			}
		}
	}

}
