package day07

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

var operators = []func(a, b int) int{
	func(a, b int) int {
		return a * b
	},
	func(a, b int) int {
		return a + b
	},
	func(a, b int) int {
		newNum, err := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
		if err != nil {
			fmt.Printf("Error converting to integer: %v\n", err)
			return 0
		}
		return newNum
	},
}

func helperRecursive(input string, isPart2 bool) int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	var sum int
	ch := make(chan int)
	var wg sync.WaitGroup

	for scanner.Scan() {
		equationString := strings.Split(scanner.Text(), ":")
		rhsString := strings.Fields(equationString[1])
		lhsInt, err := strconv.Atoi(equationString[0])
		if err != nil {
			fmt.Printf("Error converting '%s' to integer: %v\n", equationString[0], err)
		}
		rhsInt, _ := convertStringToInt(rhsString)
		wg.Add(1)
		go solve(lhsInt, rhsInt, ch, &wg, isPart2)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	for e := range ch {
		sum += e
	}
	return sum
}

func solve(result int, values []int, ch chan int, wg *sync.WaitGroup, isPart2 bool) {
	defer wg.Done()
	if num := process(values[1:], values[0], result, isPart2); num != 0 {
		ch <- num
	}
}

func process(values []int, current, result int, isPart2 bool) int {
	if current == result && len(values) == 0 {
		return result
	}
	if current > result || len(values) == 0 {
		return 0
	}

	for i, operator := range operators {
		if process(values[1:], operator(current, values[0]), result, isPart2) != 0 {
			return result
		}
		if !isPart2 && i == 1 {
			break
		}
	}
	return 0
}
