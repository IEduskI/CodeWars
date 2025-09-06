package main

import (
	"encoding/json"
	"log"
)

const data = `[
    { "gym": false, "school": true,  "store": false },
    { "gym": true,  "school": false, "store": false },
    { "gym": true,  "school": true,  "store": false },
    { "gym": false, "school": true,  "store": false },
    { "gym": false, "school": true,  "store": true }
  ]`

type Blocks struct {
	Gym    bool `json:"gym"`
	School bool `json:"school"`
	Store  bool `json:"store"`
}

func main() {
	var blocks []Blocks
	if err := json.Unmarshal([]byte(data), &blocks); err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	gymDistances := calcDistance(blocks, func(b Blocks) bool { return b.Gym })
	schoolDistances := calcDistance(blocks, func(b Blocks) bool { return b.School })
	storeDistances := calcDistance(blocks, func(b Blocks) bool { return b.Store })

	bestBlockIndex := len(blocks)
	bestMaxDist := len(blocks)

	for i := 0; i < len(blocks); i++ {
		currMaxDist := max(gymDistances[i], schoolDistances[i], storeDistances[i])
		if currMaxDist < bestMaxDist {
			bestMaxDist = currMaxDist
			bestBlockIndex = i
		}
	}

	// Output the best block index and its maximum distance
	log.Printf("Best block index: %d, Best maximum distance: %d\n", bestBlockIndex, bestMaxDist)
}

func hasService(block []Blocks, index int, service func(Blocks) bool) bool {
	if index < 0 || index >= len(block) {
		return false
	}
	return service(block[index])
}

func calcDistance(blocks []Blocks, service func(Blocks) bool) []int {
	distances := make([]int, len(blocks))
	prevIndex := -len(blocks)
	nextIndex := 2 * len(blocks)
	for i := range blocks {
		if hasService(blocks, i, service) {
			prevIndex = i
		}
		distances[i] = i - prevIndex
	}
	for i := len(blocks) - 1; i >= 0; i-- {
		if hasService(blocks, i, service) {
			nextIndex = i
		}
		distances[i] = min(distances[i], nextIndex-i)
	}
	return distances
}
