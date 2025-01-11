package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func main() {
	f, err := os.OpenFile("input/input.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	firstSet := []int{}
	secondSet := []int{}
	for sc.Scan() {
		parts := strings.Fields(sc.Text())
		val, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		firstSet = append(firstSet, val)

		val, err = strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		secondSet = append(secondSet, val)
	}
	part01(firstSet, secondSet)
	part02(firstSet, secondSet)
}

func part01(firstSet []int, secondSet []int) {
	slices.Sort(firstSet)
	slices.Sort(secondSet)
	diffSum := 0
	for i := range firstSet {
		diffSum += diff(firstSet[i], secondSet[i])
	}
	fmt.Println(diffSum)
}

func part02(firstSet []int, secondSet []int) {
	diffSum := 0
	for i := range firstSet {
		occurrence := 0
		for j := range secondSet {
			if firstSet[i] == secondSet[j] {
				occurrence++
			}
		}
		diffSum += firstSet[i] * occurrence
	}
	fmt.Println(diffSum)
}
