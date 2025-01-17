package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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
	/*f2, err2 := os.OpenFile("input/input.txt", os.O_RDONLY, os.ModePerm)
	if err2 != nil {
		log.Fatalf("open file error: %v", err2)
		return
	}
	defer f2.Close()
	sc2 := bufio.NewScanner(f2)
	part02(sc2)*/
}

func part01(sc *bufio.Scanner) {
	mulSum := 0
	for sc.Scan() {
		line := sc.Text()
		reg, _ := regexp.Compile(`mul\([1-9][0-9]{0,2}\,[1-9][0-9]{0,2}\)`)
		for _, match := range reg.FindAllString(line, -1) {
			commaIndex := strings.Index(match, ",")
			val1 := match[4:commaIndex]
			val2 := match[commaIndex+1 : len(match)-1]
			v1, _ := strconv.Atoi(val1)
			v2, _ := strconv.Atoi(val2)
			mulSum += v1 * v2
		}
	}
	fmt.Println(mulSum)
}

func part02(sc *bufio.Scanner) {

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
