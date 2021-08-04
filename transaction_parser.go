package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func writeLines(path string, lines []string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func convertLinesToMempool(lines []string, maxSize int) *Mempool {
	mp := &Mempool{}
	heap.Init(mp)

	for _, line := range lines {
		parts := strings.Fields(line)
		transaction := createTransaction(parts)
		heap.Push(mp, transaction)
		if mp.Len() > maxSize {
			heap.Pop(mp)
		}
	}

	return mp
}

func createTransaction(parts []string) Transaction {
	//Quick and easy way of reading in each part of the transaction since it is known
	hash := strings.SplitAfter(parts[0], "=")[1]
	gas, _ := strconv.ParseFloat(strings.SplitAfter(parts[1], "=")[1], 64)
	feePerGas, _ := strconv.ParseFloat(strings.SplitAfter(parts[2], "=")[1], 64)
	signature := strings.SplitAfter(parts[3], "=")[1]

	return Transaction{
		hash:      hash,
		gas:       gas,
		feePerGas: feePerGas,
		signature: signature,
	}
}
