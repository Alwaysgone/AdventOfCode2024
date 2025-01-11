package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
func safeDiff(diff int) bool {
	return (diff >= 1 && diff <= 3) ||
		(diff >= -3 && diff <= -1)
}

func main() {
	f, err := os.OpenFile("input/input.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return
	}
	defer f.Close()

	sc := bufio.NewScanner(f)

	part01(sc)
	//part02(firstSet, secondSet)
}

func part01(sc *bufio.Scanner) {
	safeReports := 0
	for sc.Scan() {
		reportString := strings.Fields(sc.Text())
		report := make([]int, len(reportString))
		for v := range reportString {
			val, err := strconv.Atoi(reportString[v])
			if err != nil {
				panic(err)
			}
			report[v] = val
		}
		//fmt.Println(report)
		safeReport := 1
		lastDiffValue := report[0] - report[1]
		if !safeDiff(lastDiffValue) {
			safeReport = 0
		}
		for i := 2; i < len(report); i++ {
			currentDiffValue := report[i-1] - report[i]
			sameSign := (lastDiffValue * currentDiffValue) > 0
			if !sameSign || !safeDiff(currentDiffValue) {
				safeReport = 0
				break
			}
		}
		safeReports += safeReport
		//fmt.Println(safeReports)
	}
	fmt.Println(safeReports)
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
