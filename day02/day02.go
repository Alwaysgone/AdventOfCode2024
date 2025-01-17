package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func safeDiff(diff int) bool {
	safe := (diff >= 1 && diff <= 3) ||
		(diff >= -3 && diff <= -1)
	if !safe {
		//fmt.Printf("unsafe diff: %v\n", diff)
	}
	return safe
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
	f2, err2 := os.OpenFile("input/input.txt", os.O_RDONLY, os.ModePerm)
	if err2 != nil {
		log.Fatalf("open file error: %v", err2)
		return
	}
	defer f2.Close()
	sc2 := bufio.NewScanner(f2)
	part02(sc2)
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
		safe, _ := isSafeReport(report)
		if safe {
			safeReports++
		}
	}
	fmt.Println(safeReports)
}

func part02(sc *bufio.Scanner) {
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

		safe, faultOccurrence := isSafeReport(report)
		if safe {
			safeReports++
		} else {
			/*
				for i := range report {
					reportCopy := make([]int, len(report))
					copy(reportCopy, report)
					reportCopy = append(reportCopy[:i], reportCopy[i+1:]...)
					copySafe, _ := isSafeReport(reportCopy)
					if copySafe {
						safeReports++
						break
					}
				}
			*/

			reportCopy1 := make([]int, len(report))
			copy(reportCopy1, report)
			reportCopy1 = append(reportCopy1[:faultOccurrence], reportCopy1[faultOccurrence+1:]...)
			copy1Safe, _ := isSafeReport(reportCopy1)

			reportCopy2 := make([]int, len(report))
			copy(reportCopy2, report)
			//if faultOccurrence+2 == len(reportCopy2) {
			//	reportCopy2 = reportCopy2[:faultOccurrence+1]
			//} else {
			reportCopy2 = append(reportCopy2[:faultOccurrence+1], reportCopy2[faultOccurrence+2:]...)
			//}
			copy2Safe, _ := isSafeReport(reportCopy2)
			if copy1Safe || copy2Safe {
				safeReports++
			}
		}
	}
	fmt.Println(safeReports)
}

func isSafeReport(report []int) (bool, int) {
	lastDiffValue := report[0] - report[1]
	if !safeDiff(lastDiffValue) {
		return false, 0
	}
	for i := 2; i < len(report); i++ {
		currentDiffValue := report[i-1] - report[i]
		sameSign := (lastDiffValue * currentDiffValue) > 0
		if !sameSign || !safeDiff(currentDiffValue) {
			return false, i - 1
		}
		lastDiffValue = currentDiffValue
	}
	return true, -1
}

func isSafeReport2(report []int) bool {
	lastDiffValue := report[0] - report[1]
	badLevelCount := 0
	if !safeDiff(lastDiffValue) {
		badLevelCount++
	}
	for i := 2; i < len(report); i++ {
		currentDiffValue := report[i-1] - report[i]
		sameSign := (lastDiffValue * currentDiffValue) > 0
		if !sameSign || !safeDiff(currentDiffValue) {
			badLevelCount++
		}
		lastDiffValue = currentDiffValue
	}
	return badLevelCount <= 1
}
