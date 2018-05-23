package main

import (
	"flag"
	"fmt"
	"strconv"
)

func main() {
	numbers := map[int]bool{}
	max := flag.Int("max", 10, "an int")
	sum := 0
	flag.Parse()

	input := flag.Args()

	for _, i := range input {
		x, err := strconv.ParseInt(i, 10, 0)
		if err != nil {
			break
		}
		cnt := 0

		for u := 0; u < *max; u = int(x) * cnt {
			cnt++
			if numbers[u] != true {
				// fmt.Println("-> ", u)
				sum += u
				numbers[u] = true
			}
		}
	}

	fmt.Println("total sum was ", sum)
}
