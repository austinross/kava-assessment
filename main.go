package main

import (
	"log"
)

func main() {
	lines, err := readLines("transactions.txt")

	if err != nil {
		log.Fatalf("Could not read in file: %s", err)
	}

	mp := convertLinesToMempool(lines, 5000)

	if err := writeLines("prioritized-transactions.txt", mp.ToStrings()); err != nil {
		log.Fatalf("Could not write to file: %s", err)
	}
}
