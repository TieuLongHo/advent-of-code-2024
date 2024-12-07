package day07

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
)

func helper(input string, isPart2 bool) int {
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
		go fixEquation(lhsInt, rhsInt, ch, &wg, isPart2)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()
	for e := range ch {
		sum += e
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}

	return sum
}

func fixEquation(lhs int, rhs []int, ch chan<- int, wg *sync.WaitGroup, isPart2 bool) {
	defer wg.Done()
	for i := int(math.Pow(2, float64(len(rhs)-1))); i >= 0; i-- {
		binary := fmt.Sprintf("%0*b", len(rhs), i)
		var result int = rhs[0]
		for j := 1; j < len(rhs); j++ {
			switch binary[j] {
			case '0':
				result = mul(result, rhs[j])
			case '1':
				result = add(result, rhs[j])
			}
			if result > lhs {
				break
			}
			if result == lhs && j == len(rhs)-1 {
				ch <- lhs
				return
			}
		}
	}

	if isPart2 {
		wg.Add(1)
		tryConcat(lhs, rhs, ch, wg)
	}

}

func tryConcat(lhs int, rhs []int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := int(math.Pow(3, float64(len(rhs)-1))); i >= 0; i-- {
		base3 := strconv.FormatInt(int64(i), 3)
		if len(base3) < len(rhs) {
			base3 = strings.Repeat("0", len(rhs)-len(base3)) + base3
		}
		var result int = rhs[0]
		for j := 1; j < len(rhs); j++ {
			switch base3[j] {
			case '0':
				result = concat(result, rhs[j])
			case '1':
				result = mul(result, rhs[j])
			case '2':
				result = add(result, rhs[j])
			}

			if result > lhs {
				break

			}
			if result == lhs && j == len(rhs)-1 {
				ch <- lhs
				return
			}
		}
	}
}

func add(a, b int) int {
	return a + b
}
func mul(a, b int) int {
	return a * b
}
func concat(a, b int) int {
	newNum, err := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
	if err != nil {
		fmt.Printf("Error converting to integer: %v\n", err)
		return 0
	}
	return newNum
}

func convertStringToInt(strSlice []string) (newIntSlice []int, err error) {
	intSlice := make([]int, len(strSlice))
	for i, str := range strSlice {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Printf("Error converting '%s' to integer: %v\n", str, err)
			return []int{}, fmt.Errorf("error converting")
		}
		intSlice[i] = num
	}
	return intSlice, nil
}
