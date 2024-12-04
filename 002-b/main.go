package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var nums [][]int

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")

		var num []int

		for _, s := range parts {
			n, _ := strconv.Atoi(s)
			num = append(num, n)
		}

		nums = append(nums, num)
	}

	total := 0

	for _, n := range nums {
		if isSafe(n) || canBeSafe((n)) {
			total += 1
		}
	}

	fmt.Println(total)
}

func isSafe(nums []int) bool {
	increasing := false

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			if i == 1 {
				increasing = true
			} else if !increasing {
				return false
			}

			diff := nums[i] - nums[i-1]

			if diff >= 1 && diff <= 3 {
				continue
			}

			return false
		} else if nums[i] < nums[i-1] {
			if i == 1 {
				increasing = false
			} else if increasing {
				return false
			}

			diff := nums[i-1] - nums[i]

			if diff >= 1 && diff <= 3 {
				continue
			}
			return false
		} else {
			return false
		}
	}

	return true
}

func canBeSafe(nums []int) bool {
	for i := range nums {
		if isSafe(remove(nums, i)) {
			return true
		}
	}

	return false
}

func remove(slice []int, s int) []int {
	n := make([]int, len(slice))
	copy(n, slice)
	return append(n[:s], n[s+1:]...)
}
