package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
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

type instruction struct {
	index  int
	enable bool
}

func indexAt(s, substring string, n int) int {
	idx := strings.Index(s[n:], substring)
	if idx > -1 {
		idx += n
	}
	return idx
}

func part02(sc *bufio.Scanner) {
	mulSum := 0
	line := ""
	for sc.Scan() {
		line += sc.Text()
	}
	instructions := []instruction{}

	currentIndex := 0
	for {
		currentIndex = indexAt(line, `do()`, currentIndex)
		if currentIndex < 0 {
			break
		}
		instructions = append(instructions, instruction{currentIndex, true})
		currentIndex++
	}
	currentIndex = 0
	for {
		currentIndex = indexAt(line, `don't()`, currentIndex)
		if currentIndex < 0 {
			break
		}
		instructions = append(instructions, instruction{currentIndex, false})
		currentIndex++
	}
	sort.Slice(instructions[:], func(i, j int) bool {
		return instructions[i].index < instructions[j].index
	})

	instructionIndex := 0
	enabled := true
	reg, _ := regexp.Compile(`mul\([1-9][0-9]{0,2}\,[1-9][0-9]{0,2}\)`)
	indexes := reg.FindAllStringIndex(line, -1)
	for index, match := range reg.FindAllString(line, -1) {
		for instructionIndex < len(instructions) && instructions[instructionIndex].index <= indexes[index][0] {
			enabled = instructions[instructionIndex].enable
			instructionIndex++
		}
		if enabled {
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
