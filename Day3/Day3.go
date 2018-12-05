package main

import (
	"strconv"
	"log"
	"fmt"
	"regexp"
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
	var fabric = map[int]map[int]int{}

	for row := range input {
		var _, left, top, x, y = getCut(input[row])
		for i := top; i < top+y; i++ {
			for j := left; j < left+x; j++ {
				if _, ok := fabric[i]; !ok {
					fabric[i] = make(map[int]int)
				}

				fabric[i][j] += 1;
			}	
		} 
	}

	return processOverlap(fabric)
}

func getCut (line string) (int, int, int, int, int){
	r := regexp.MustCompile(`#([0-9]+) [@] ([0-9]+),([0-9]+): ([0-9]+)x([0-9]+)`)
	var groups []string = r.FindStringSubmatch(line);
	return cutAtoi(groups[1]), cutAtoi(groups[2]), cutAtoi(groups[3]), cutAtoi(groups[4]), cutAtoi(groups[5])
}

func cutAtoi(cut string) (int) {
	var cutI, err = strconv.Atoi(cut)
	if(err != nil){
		log.Fatal(err)
	} 
	
	return cutI
}

func processOverlap(fabric map[int]map[int]int) int {
	var count int = 0;
	for _, row := range fabric{
		for _, sqInch := range row {
			if sqInch > 1 { count += 1 }
		}
	}

	return count;
}

func solveB (input []string, init int) (int) {
	var fabric = map[int]map[int]int{}

	for row := range input {
		var _, left, top, x, y = getCut(input[row])
		for i := top; i < top+y; i++ {
			for j := left; j < left+x; j++ {
				if _, ok := fabric[i]; !ok {
					fabric[i] = make(map[int]int)
				}

				fabric[i][j] += 1;
			}	
		} 
	}

	return processUnderlap(fabric, input)
}

func processUnderlap(fabric map[int]map[int]int, input []string) int {

	var cleanID int = -1;

	for row := range input {
		var cleanCut bool = true

		var id, left, top, x, y = getCut(input[row])
		for i := top; i < top+y; i++ {
			for j := left; j < left+x; j++ {
				if(fabric[i][j] > 1) {
					cleanCut = false;
					continue
				}
			}
			if !cleanCut{ continue }
		}
		if (cleanCut) { cleanID = id}
	}

	return cleanID;
}