package day01

import (
	"fmt"
	"log"
	"strconv"

	"mrnateriver.io/advent_of_code_2024/shared"
)

func CalculateTotalDistance() int {
	var left, right *shared.Tree

	leftChan := make(chan string)
	rightChan := make(chan string)

	go fillTree(&left, leftChan)
	go fillTree(&right, rightChan)

	ch := readInput()
	for {
		v, ok := <-ch
		if !ok {
			break
		}

		leftChan <- v[0]
		rightChan <- v[1]
	}

	close(leftChan)
	close(rightChan)

	log.Println("Left tree size: ", left.Size())
	log.Println("Right tree size: ", right.Size())

	log.Println("Calculating distance...")

	leftWalker := left.Walker()
	rightWalker := right.Walker()

	lineCount := 0
	totalDistance := 0
	for {
		v1, ok1 := <-leftWalker
		v2, ok2 := <-rightWalker
		if !ok1 || !ok2 {
			break
		}

		dist := v1 - v2
		if dist < 0 {
			dist = -dist
		}

		totalDistance += dist

		if lineCount%100 == 0 {
			log.Println("Processed ", lineCount, " lines")
		}

		lineCount++
	}

	log.Println("Processed ", lineCount, " lines")

	return totalDistance
}

func fillTree(t **shared.Tree, ch chan string) {
	for {
		v, ok := <-ch
		if !ok {
			break
		}

		parsed, err := strconv.Atoi(v)
		if err != nil {
			panic(fmt.Errorf("invalid number: %s", v))
		}

		*t = (*t).Insert(parsed)
	}
}
