package day09

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func Part1(input string) int {
	return helper(input)
}

func helper(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	var newLine string

	for scanner.Scan() {
		newLine = (scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
	return solve(newLine)
}

func solve(input string) int {
	compactFile := compressFile(mapToFile(input))

	var checksum int
	for i, e := range compactFile {
		if e != -1 {
			checksum += i * e
		}
	}

	return checksum
}

func compressFile(fileBlock []int) []int {
	prevIndex := len(fileBlock) - 1

OuterLoop:
	for i, e := range fileBlock {
		if e != -1 {
			continue
		}
		for j := prevIndex; j >= 0; j-- {
			if j <= i {
				break OuterLoop
			}
			if fileBlock[j] == -1 {
				continue
			}
			prevIndex = j
			fileBlock[i] = fileBlock[j]
			fileBlock[j] = -1
			break
		}
	}

	return fileBlock
}

func mapToFile(input string) []int {
	var file []int
	for i, e := range input {
		fileVal, _ := strconv.Atoi(string(e))
		if i%2 == 0 {
			blockId := i / 2
			for j := 0; j < fileVal; j++ {
				file = append(file, blockId)
			}
			continue
		}
		for j := 0; j < fileVal; j++ {
			file = append(file, -1)
		}
	}
	return file
}
