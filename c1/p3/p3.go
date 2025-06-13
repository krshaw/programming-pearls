package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/kelindar/bitmap"
)

const MAX_SIZE = 10000000

func fastSort(file *os.File) {
	var bits bitmap.Bitmap
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		bits.Set(uint32(n))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	f, err := os.OpenFile("./c1/p3/output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for i := 0; i < MAX_SIZE; i++ {
		if bits.Contains(uint32(i)) {
			_, err = f.WriteString(strconv.Itoa(i) + "\n")
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func slowSort(file *os.File) {
	nums := make([]int, 1000000)
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		nums[i] = n
		i++
	}
	// sort nums
	sort.Ints(nums)
	f, err := os.OpenFile("./c1/p3/output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for _, n := range nums {
		_, err = f.WriteString(strconv.Itoa(n) + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	file, err := os.Open("./c1/p3/nums.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fastSort(file)
	slowSort(file)
}
