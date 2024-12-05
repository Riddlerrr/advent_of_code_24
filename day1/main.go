package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type List struct {
	items []int
}

func (l *List) AddItem(item int) {
	l.items = append(l.items, item)
}

func (l *List) Sort() {
	sort.Ints(l.items)
}

func (l *List) Count(item int) int {
	count := 0
	for _, i := range l.items {
		if i == item {
			count++
		}
	}
	return count
}

func ReadFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}

	return lines
}

func parseLine(line string) (int, int) {
	if len(line) == 0 {
		return 0, 0
	}

	parts := strings.Split(line, "   ")
	if len(parts) != 2 {
		return 0, 0
	}

	left, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0
	}

	right, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0
	}

	return left, right
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	content := ReadFile("challenge.txt")
	left_list := List{}
	right_list := List{}

	for _, line := range content {
		l, r := parseLine(line)
		left_list.AddItem(l)
		right_list.AddItem(r)
	}

	left_list.Sort()
	right_list.Sort()

	part_one_sum := 0
	part_two_sum := 0
	for i := 0; i < len(left_list.items); i++ {
		l_item := left_list.items[i]
		r_item := right_list.items[i]

		part_one_sum += abs(l_item - r_item)
		part_two_sum += l_item * right_list.Count(l_item)
	}

	fmt.Println("Answer for part one:", part_one_sum)
	fmt.Println("Answer for part two:", part_two_sum)
}
