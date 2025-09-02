package main

import (
	"encoding/json"
	"fmt"
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

	_ = json.Unmarshal([]byte(data), &blocks)

	var bestBlockIndex int
	bestMaxDist := len(blocks)

	for i := 0; i < len(blocks); i++ {
		currGymMinDist := distanceToNearest(blocks, i, func(b Blocks) bool { return b.Gym })
		currSchoolMinDist := distanceToNearest(blocks, i, func(b Blocks) bool { return b.School })
		currStoreMinDist := distanceToNearest(blocks, i, func(b Blocks) bool { return b.Store })

		currMaxDist := max(currGymMinDist, currSchoolMinDist, currStoreMinDist)

		if currMaxDist < bestMaxDist {
			bestMaxDist = currMaxDist
			bestBlockIndex = i
		}
	}

	fmt.Println("Best block index:", bestBlockIndex)
}

func distanceToNearest(blocks []Blocks, blockIndex int, service func(Blocks) bool) int {
	n := len(blocks)
	minDist := n

	for i := 0; i < n; i++ {
		dist := abs(i - blockIndex)
		if service(blocks[i]) {
			if dist < minDist {
				minDist = dist
			}
		}
	}
	return minDist
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
