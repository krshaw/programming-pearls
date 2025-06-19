package main

import (
	"log"
	"math/rand/v2"
	"os"
	"strconv"

	"github.com/kelindar/bitmap"
)

func generateRandomNumbers(n, k int) {
	f, err := os.OpenFile("./c1/p4/output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	var bits bitmap.Bitmap
	// need to keep some memory because we need to keep track of what nums have already been generated
	for i := 0; i < k; i++ {
		// generate a random integer between 0 and n-1 inclusive
		// keep doing this until we find one not in the bitmap
		var j int
		for j = rand.IntN(n); bits.Contains(uint32(j)); j = rand.IntN(n) {
		}
		// j is not in the bitmap, add it and write it out to the file
		bits.Set(uint32(j))
		if _, err := f.WriteString(strconv.Itoa(j) + "\n"); err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	generateRandomNumbers(10000000, 1000000)
}
