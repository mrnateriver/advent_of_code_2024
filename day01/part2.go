package day01

import (
	"fmt"
	"log"
	"strconv"
)

type pair struct {
	score int
	found bool
}

func CalculateSimilarityScore() int {
	scores := make(map[string]pair)

	ch := readInput()
	for {
		v, ok := <-ch
		if !ok {
			break
		}

		v0, v1 := v[0], v[1]

		parsedV1, err := strconv.Atoi(v1)
		if err != nil {
			panic(fmt.Errorf("error parsing int: %v", err))
		}

		if p, ok := scores[v1]; ok {
			scores[v1] = pair{score: p.score + parsedV1, found: p.found}
		} else {
			scores[v1] = pair{score: parsedV1, found: false}
		}

		if p, ok := scores[v0]; ok {
			scores[v0] = pair{score: p.score, found: true}
		} else {
			scores[v0] = pair{score: 0, found: true}
		}
	}

	log.Println("Scores map size: ", len(scores))

	similarityScore := 0
	for _, p := range scores {
		if p.found {
			similarityScore += p.score
		}
	}

	return similarityScore
}
