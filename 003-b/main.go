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
	r, _ := regexp.Compile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)
	matches := r.FindAllString(string(data), -1)

	enabled := true
	total := 0

	for _, match := range matches {
		if match == "do()" {
			enabled = true
			continue
		} else if match == "don't()" {
			enabled = false
			continue
		}

		if enabled {
			r, _ := regexp.Compile(`(\d+),(\d+)`)
			m := r.FindAllString(match, -1)
			parts := strings.Split(m[0], ",")

			a, _ := strconv.Atoi(parts[0])
			b, _ := strconv.Atoi(parts[1])

			total += (a * b)
		}
	}

	fmt.Println(total)
}
