package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	var left, right []int

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "   ")
		p1, _ := strconv.Atoi(parts[0])
		p2, _ := strconv.Atoi(parts[1])
		left = append(left, p1)
		right = append(right, p2)
	}

	sort.Ints(left)
	sort.Ints(right)

	total := 0

	for i := 0; i < len(left); i++ {
		total += abs(left[i] - right[i])
	}

	fmt.Println(total)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
