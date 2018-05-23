package main

import (
	"fmt"
)

/**	*/
func fibGenerator(quit <-chan struct{}) <-chan int {
	rtnNumber := make(chan int)

	go func() {
		defer close(rtnNumber)
		for i, j := 0, 1; ; i, j = i+j, i {
			select {
			case rtnNumber <- i:
			case <-quit:
				return
			}
		}
	}()

	return rtnNumber
}

func main() {
	sum := 0
	quit := make(chan struct{})
	defer close(quit)
	for i := range fibGenerator(quit) {
		if i >= 4000000 {
			quit <- struct{}{}
		} else if i%2 == 0 {
			sum += i
		}
	}
	fmt.Println("total sum was ", sum)
}
