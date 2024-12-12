package shared

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type LineProcessor[T any] func(string) (T, error)

func ReadInputLines(file string) <-chan string {
	return ReadInput(file, func(line string) (string, error) {
		return line, nil
	})
}

func ReadInput[T any](file string, processLine LineProcessor[T]) <-chan T {
	ch := make(chan T)

	go func(ch chan T) {
		file, err := os.Open(file)
		if err != nil {
			panic(fmt.Errorf("error opening file: %v", err))
		}
		defer file.Close()

		fileStats, err := file.Stat()
		log.Println("Opened file:", fileStats.Size(), "bytes")

		log.Println("Reading file...")

		lineCount := 0
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()

			res, err := processLine(line)
			if err != nil {
				panic(fmt.Errorf("invalid line: %s; %w", line, err))
			}

			ch <- res

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
