package main

import (
	"log"
	"fmt"
	"github.com/gramstrong/aoc2018/utils"
)

func main() {
	var input, err = utils.Read("input.txt")
	if err != nil {
    log.Fatal(err)
	} else {
		fmt.Printf("A: %d\n", solveA(input));
		fmt.Printf("B: %s", solveB(input));
	}
}

func solveA (input []string) (int) {
	var dubsSum int = 0
	var trisSum int = 0

	for i := range input {
		var count map[rune]int = make(map[rune]int)
		var dubs bool = false
		var tris bool = false;

		for _, val := range input[i] {
			count[val] += 1
		}

		for _, val := range count {
			if val == 2 { dubs = true }
			if val > 2 { tris = true }		
		}

		if(dubs) { dubsSum += 1 }
		if(tris) { trisSum += 1 }
	}

	return dubsSum * trisSum;
}

func solveB (input []string) (string) {
	for i := range input{
		for j := range input {
			var numDiff int = 0;
			var answer string; 

			for ch := range input[i]{
				if(input[i][ch] != input[j][ch]) {
					numDiff += 1;
					answer = input[i][:ch] + input[i][ch+1:]
				}
				if(numDiff > 1){
					break
				}
			}

			if numDiff == 1 {
				return answer
			}
		}
	}
	
	return "?";
}