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

	gymDistances := make([]int, len(blocks))
	schoolDistances := make([]int, len(blocks))
	storeDistances := make([]int, len(blocks))

	prevGymIndex := -len(blocks)
	nextGymIndex := 2 * len(blocks)
	for i := range blocks {
		if hasService(blocks, i, func(b Blocks) bool { return b.Gym }) {
			prevGymIndex = i
		}
		gymDistances[i] = i - prevGymIndex
	}

	for i := len(blocks) - 1; i >= 0; i-- {
		if hasService(blocks, i, func(b Blocks) bool { return b.Gym }) {
			nextGymIndex = i
		}
		gymDistances[i] = min(gymDistances[i], nextGymIndex-i)
	}

	prevSchoolIndex := -len(blocks)
	nextSchoolIndex := 2 * len(blocks)
	for i := range blocks {
		if hasService(blocks, i, func(b Blocks) bool { return b.School }) {
			prevSchoolIndex = i
		}
		schoolDistances[i] = i - prevSchoolIndex
	}

	for i := len(blocks) - 1; i >= 0; i-- {
		if hasService(blocks, i, func(b Blocks) bool { return b.School }) {
			nextSchoolIndex = i
		}
		schoolDistances[i] = min(schoolDistances[i], nextSchoolIndex-i)
	}

	prevStoreIndex := -len(blocks)
	nextStoreIndex := 2 * len(blocks)
	for i := range blocks {
		if hasService(blocks, i, func(b Blocks) bool { return b.Store }) {
			prevStoreIndex = i
		}
		storeDistances[i] = i - prevStoreIndex
	}

	for i := len(blocks) - 1; i >= 0; i-- {
		if hasService(blocks, i, func(b Blocks) bool { return b.Store }) {
			nextStoreIndex = i
		}
		storeDistances[i] = min(storeDistances[i], nextStoreIndex-i)
	}

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
