package main

import (
	"strconv"
	"log"
	"fmt"
	"github.com/gramstrong/aoc2018/utils"
)

func main() {
	var input, err = utils.Read("input.txt")
	if err != nil {
    log.Fatal(err)
	} else {
		fmt.Printf("A: %d\n", solveA(input, 0));
		fmt.Printf("B: %d", solveB(input, 0));

	}
}

func solveA (input []string, init int) (int) {
	for i := range input {
		var num, err = strconv.Atoi(input[i])
		
		if err != nil {
			log.Fatal(err)
		} else {
			init += num;
		}

	}
	return init;
}

func solveB (input []string, init int) (int) {
	var dups map[int]int
	var found bool = false

	dups = make(map[int]int)
	dups[init] = 1;

	for !found {
		for i := range input {
			var num, err = strconv.Atoi(input[i])

			if err != nil {
				log.Fatal(err)
			} else {
				init += num;
	
				if dups[init] == 1 {
					found = true;
					break
				}
	
				dups[init] = 1;
			}
	
		}
	}
	
	return init;
}