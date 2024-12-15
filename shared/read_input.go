package shared

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type LineProcessor[T any] func(string) (T, error)

func ReadInputLines(file string) <-chan string {
	ch := make(chan string)

	go func(ch chan string) {
		file, err := os.Open(file)
		if err != nil {
			panic(fmt.Errorf("error opening file: %v", err))
		}
		defer func() {
			err := file.Close()
			if err != nil {
				panic(fmt.Errorf("error closing file: %v", err))
			}
		}()

		fileStats, err := file.Stat()
		log.Println("Opened file:", fileStats.Size(), "bytes")

		log.Println("Reading file...")

		lineCount := 0
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()

			ch <- line

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

func ReadInput[T any](file string, processLine LineProcessor[T]) <-chan T {
	ch := make(chan T)

	go func(ch chan T) {
		for line := range ReadInputLines(file) {
			res, err := processLine(line)
			if err != nil {
				panic(fmt.Errorf("invalid line: %s; %w", line, err))
			}

			ch <- res
		}
		close(ch)
	}(ch)

	return ch
}
