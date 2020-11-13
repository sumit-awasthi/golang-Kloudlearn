// generating random numbers using channel and predefined rand func()
package main

import (
	"math/rand"
	"sync"
	"time"
)

func randomint(wg *sync.WaitGroup, c chan int) {
	defer wg.Done()

	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= 15; i++ { // Printing 15 values storing in buffer
		r := rand.Intn(100)
		c <- r
	}

}
func main() {
	var wg sync.WaitGroup

	c := make(chan int, 15) //channel with buffer
	wg.Add(1)

	go randomint(&wg, c)
	wg.Wait()
	close(c)

	for item := range c {
		println(item)
	}
}
