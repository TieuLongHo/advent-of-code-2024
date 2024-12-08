package day08

import (
	"bufio"
	"fmt"
	"strings"
)

func Part1(input string) int {
	return helper(input, false)
}

func helper(input string, isPart2 bool) int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	antennaMap := make(map[byte][]Node)
	var antennaCount int
	var grid [][]byte
	var i int
	for scanner.Scan() {
		newLine := []byte(scanner.Text())
		for j := 0; j < len(newLine); j++ {
			if newLine[j] != '.' {
				antennaMap[newLine[j]] = append(antennaMap[newLine[j]], Node{j, i})
				antennaCount++
			}
		}
		grid = append(grid, newLine)
		i++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
	return solve(antennaMap, grid, isPart2)
}

func solve(antennaMap map[byte][]Node, grid [][]byte, isPart2 bool) int {
	antiNodeMap := make(map[Node]bool)
	var antennaCount int

	for frequency, antennas := range antennaMap {
		antennaCount += len(antennas)
		for _, antenna := range antennas {
			antennaMap[frequency] = antennaMap[frequency][1:]
			for i := 0; i < len(antennaMap[frequency]); i++ {
				antiNode, err := getAllAntiNodes(grid, antenna, antennaMap[frequency][i], isPart2)
				if err != nil {
					continue
				}
				for _, an := range antiNode {
					antiNodeMap[an] = true
				}
			}
		}
	}
	score := len(antiNodeMap)
	if isPart2 {
		score += antennaCount
	}
	return score
}

func getAllAntiNodes(grid [][]byte, pos1, pos2 Node, isPart2 bool) (antiNodes []Node, err error) {

	rotateVector := func(vector [2]int) [2]int {
		return [2]int{vector[0] * -1, vector[1] * -1}
	}
	vector := pos1.getVector(pos2)
	var potentialNodes []Node

	if !isPart2 {
		potentialNodes = []Node{
			{pos2.x + vector[1], pos2.y + vector[0]},
			{pos1.x + rotateVector(vector)[1], pos1.y + rotateVector(vector)[0]},
		}
		for _, node := range potentialNodes {
			if node.isInBound(grid) {
				antiNodes = append(antiNodes, node)
			}
		}
	} else {
		getAllResonantHarmonics(pos2, vector, &potentialNodes, grid)
		getAllResonantHarmonics(pos1, rotateVector(vector), &potentialNodes, grid)
		antiNodes = append(antiNodes, potentialNodes...)
	}
	if len(antiNodes) == 0 {
		return []Node{}, fmt.Errorf("all anti-nodes out of bounds")
	}
	return antiNodes, nil
}

func getAllResonantHarmonics(pos Node, vector [2]int, potentialNodes *[]Node, grid [][]byte) {
	potentialNode := Node{pos.x + vector[1], pos.y + vector[0]}
	if potentialNode.isInBound(grid) {
		if grid[potentialNode.y][potentialNode.x] == '.' {
			*potentialNodes = append(*potentialNodes, potentialNode)
		}
		getAllResonantHarmonics(potentialNode, vector, potentialNodes, grid)
	}
}
