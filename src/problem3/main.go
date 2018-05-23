package main

import (
	"fmt"
	"os"
	"strconv"
)

/**	*/
func primeNumbers(quit <-chan struct{}) <-chan int {
	primeNumbers := []int{}
	rtnNumber := make(chan int)
	activeNumber := 3
	primeNumbers = append(primeNumbers, 2)

	go func() {
		defer close(rtnNumber)
		for {
			for {
				found := false
				for i := range primeNumbers {
					if activeNumber%primeNumbers[i] == 0 {
						found = true
						break
					}
				}
				if !found {
					primeNumbers = append(primeNumbers, activeNumber)
					break
				}
				activeNumber++

			}

			select {

			case rtnNumber <- activeNumber:
			case <-quit:
				return
			}
		}
	}()

	return rtnNumber
}

func main() {
	primeFactors := []int{}
	x, err := strconv.ParseInt(os.Args[1], 10, 0) //600851475143 //13195
	if err != nil {
		return
	}
	number := int(x)
	quit := make(chan struct{})
	defer close(quit)

	for i := range primeNumbers(quit) {

		for {
			if number%i != 0 {
				break
			}
			primeFactors = append(primeFactors, i)
			number = number / i
		}
		if number == 1 {
			quit <- struct{}{}
		}
	}

	fmt.Println(primeFactors)
	fmt.Println("Largest prime factor was", primeFactors[len(primeFactors)-1])
}
