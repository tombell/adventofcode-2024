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

	var left, right []int

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "   ")
		p1, _ := strconv.Atoi(parts[0])
		p2, _ := strconv.Atoi(parts[1])
		left = append(left, p1)
		right = append(right, p2)
	}

	scores := make(map[int]int)

	for _, i := range left {
		for _, j := range right {
			if i == j {
				scores[i] += j
			}
		}
	}

	total := 0

	for _, val := range scores {
		total += val
	}

	fmt.Println(total)
}
