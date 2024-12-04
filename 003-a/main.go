package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("./input.txt")
	r, _ := regexp.Compile(`mul\((\d+),(\d+)\)`)
	matches := r.FindAllString(string(data), -1)

	total := 0

	for _, match := range matches {
		r, _ := regexp.Compile(`(\d+),(\d+)`)
		m := r.FindAllString(match, -1)
		parts := strings.Split(m[0], ",")

		a, _ := strconv.Atoi(parts[0])
		b, _ := strconv.Atoi(parts[1])

		total += (a * b)
	}

	fmt.Println(total)
}
