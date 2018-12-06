package main

import (
	"log"
	"fmt"
	"strings"
	"github.com/gramstrong/aoc2018/utils"
)

func main() {
	var input, err = utils.Read("input.txt")
	if err != nil {
    log.Fatal(err)
	} else {
		fmt.Printf("A: %d\n", solveA(input));
		fmt.Printf("B: %d", solveB(input));
	}
}

func solveA (input []string) int {

	var line string = input[0];
	return removePolymer(line)
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func minInt(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func removePolymer(line string) int {
	var i int = 0;

	for i < len(line)-1 {
		var val int = int(line[i]) - int(line[i+1])
		if absInt(val) == 32 {
			line = utils.RemoveCharAtIndex(line, i)
			line = utils.RemoveCharAtIndex(line, i)
			i = 0;
		} else {
			i+=1
		}
	}

	return len(line)
}

func removePolymer2(minLen int, line string) int {
	var i int = 0;

	for i < len(line)-1 {
		var val int = int(line[i]) - int(line[i+1])
		if absInt(val) == 32 {
			line = utils.RemoveCharAtIndex(line, i)
			line = utils.RemoveCharAtIndex(line, i)
			i = 0;
		} else {
			i+=1
		}
	}
	return minInt(minLen, len(line))
}

func solveB (input []string) (int) {
	var line string = input[0]
	var minLen int = len(line);

	minLen = removePolymer2(minLen, strings.Replace(strings.Replace(line, "a", "", -1), "A", "", -1))
	minLen = removePolymer2(minLen, strings.Replace(strings.Replace(line, "b", "", -1), "B", "", -1)) 
	minLen = removePolymer2(minLen, strings.Replace(strings.Replace(line, "c", "", -1), "C", "", -1)) 
	minLen = removePolymer2(minLen, strings.Replace(strings.Replace(line, "d", "", -1), "D", "", -1)) 
	minLen = removePolymer2(minLen, strings.Replace(strings.Replace(line, "e", "", -1), "E", "", -1)) 
	minLen = removePolymer2(minLen, strings.Replace(strings.Replace(line, "f", "", -1), "F", "", -1)) 
	minLen = removePolymer2(minLen, strings.Replace(strings.Replace(line, "g", "", -1), "G", "", -1)) 
	minLen = removePolymer2(minLen, strings.Replace(strings.Replace(line, "h", "", -1), "H", "", -1)) 
	minLen = removePolymer2(minLen, strings.Replace(strings.Replace(line, "i", "", -1), "I", "", -1)) 
	minLen = removePolymer2(minLen, strings.Replace(strings.Replace(line, "j", "", -1), "J", "", -1)) 
	minLen = removePolymer2(minLen, strings.Replace(strings.Replace(line, "k", "", -1), "K", "", -1)) 
	minLen = removePolymer2(minLen, strings.Replace(strings.Replace(line, "l", "", -1), "L", "", -1)) 
	minLen = removePolymer2(minLen, strings.Replace(strings.Replace(line, "m", "", -1), "M", "", -1)) 
	minLen = removePolymer2(minLen, strings.Replace(strings.Replace(line, "n", "", -1), "N", "", -1)) 
	minLen = removePolymer2(minLen, strings.Replace(strings.Replace(line, "o", "", -1), "O", "", -1)) 
	minLen = removePolymer2(minLen, strings.Replace(strings.Replace(line, "p", "", -1), "P", "", -1)) 
	minLen = removePolymer2(minLen, strings.Replace(strings.Replace(line, "q", "", -1), "Q", "", -1)) 
	minLen = removePolymer2(minLen, strings.Replace(strings.Replace(line, "r", "", -1), "R", "", -1)) 
	minLen = removePolymer2(minLen, strings.Replace(strings.Replace(line, "s", "", -1), "S", "", -1)) 
	minLen = removePolymer2(minLen, strings.Replace(strings.Replace(line, "t", "", -1), "T", "", -1)) 
	minLen = removePolymer2(minLen, strings.Replace(strings.Replace(line, "u", "", -1), "U", "", -1)) 
	minLen = removePolymer2(minLen, strings.Replace(strings.Replace(line, "v", "", -1), "V", "", -1)) 
	minLen = removePolymer2(minLen, strings.Replace(strings.Replace(line, "w", "", -1), "W", "", -1)) 
	minLen = removePolymer2(minLen, strings.Replace(strings.Replace(line, "x", "", -1), "X", "", -1)) 
	minLen = removePolymer2(minLen, strings.Replace(strings.Replace(line, "y", "", -1), "Y", "", -1)) 
	minLen = removePolymer2(minLen, strings.Replace(strings.Replace(line, "z", "", -1), "Z", "", -1))

	return minLen
}