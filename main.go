package main

import (
	"fmt"
	"math/big"
	"time"
)

func main() {
	currentTime := time.Now()
	unixTimestamp := currentTime.Unix()

	numbersToCheck := []*big.Int{
		big.NewInt(unixTimestamp),
	}

	k := 23 // Number of iterations (witnesses) - higher k increases certainty

	fmt.Printf("Running Miller-Rabin test with %d iterations:\n", k)
	for _, num := range numbersToCheck {
		isPrime := MillerRabin(num, k)
		fmt.Printf("Is %s likely prime? %t\n", num.String(), isPrime)
	}
}