package day01

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readInput() <-chan [2]string {
	ch := make(chan [2]string)

	go func(ch chan [2]string) {
		file, err := os.Open("day01/input")
		if err != nil {
			panic(fmt.Errorf("error opening file: %v", err))
		}
		defer file.Close()

		fileStats, err := file.Stat()
		log.Println("Read file: ", fileStats.Size(), "bytes")

		log.Println("Reading file...")

		lineCount := 0
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			cols := strings.Split(line, "   ")

			if len(cols) != 2 {
				panic(fmt.Errorf("invalid line: %s", line))
			}

			ch <- [2]string{cols[0], cols[1]}

			lineCount++

			if lineCount%100 == 0 {
				log.Println("Read ", lineCount, " lines")
			}
		}
		close(ch)

		if err := scanner.Err(); err != nil {
			panic(fmt.Errorf("error reading file: %v", err))
		}

		log.Println("Read ", lineCount, " lines")
	}(ch)

	return ch
}
