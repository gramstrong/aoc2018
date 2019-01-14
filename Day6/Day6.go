package main

import (
	"strconv"
	"log"
	"fmt"
	"sort"
	"regexp"
	"math"
	"github.com/gramstrong/aoc2018/utils"
)

func main() {
	var input, err = utils.Read("input.txt")
	if err != nil {
    log.Fatal(err)
	} else {
		fmt.Printf("A: %d\n", solveA(input));
		//fmt.Printf("B: %d", solveB(input, 0));
	}
}

func solveA (input []string) int {
	var sizes map[string]int = make(map[string]int)

	minx, miny, maxx, maxy := getMinMax(input)

	print("minX: ", minx, "\n")
	print("minY: ", miny, "\n")
	print("maxX: ", maxx, "\n")
	print("maxY: ", maxy, "\n")

	var grid = make([][]string, maxx)
	for i := range grid {
		grid[i] = make([]string, maxy)
	}

	for i := range grid {
		for j := range grid[i] {
			minDist := 10000
			for k := range input {
				x, y := getXY(input[k])
				dist := euclDist(i,j,x,y)
				if dist < minDist {
					grid[i][j] = fmt.Sprintf("%d, %d", x, y)
					minDist = dist
				} else if dist == 0 {
					grid[i][j] = "0000"
				} else if dist == minDist {
					grid[i][j] = "0000"					
				}
			}
		}
	}

	for key, _ := range sizes {
		x, y := getXY(key)
		grid[x][y] = "0000"
	}

	for i := range grid {
		for j := range grid {
			err := utils.Write("./output.txt", grid[i][j] + " ")

			if err != nil {
				log.Fatal(err)
			}

			sizes[grid[i][j]] += 1
		}
		err := utils.Write("/output.txt", "\n")

		if err != nil {
			log.Fatal(err)
		}

	}



	for i := range grid[0] {
		sizes[grid[0][i]] = 0
		sizes[grid[len(grid)-1][i]] = 0
	}

	for i:= range grid {
		sizes[grid[i][0]] = 0
		sizes[grid[i][len(grid[i])-1]] = 0
	}

	var max string = "0000";
	print("size: ", len(grid), "x", len(grid[0]), "\n")
	for key, val := range sizes {
		print(key, ": ", val, "\n")
		if val > sizes[max] {
			max = key
		}
	}
	
	return sizes[max]
}

func getXY(input string) (int, int) {
	r := regexp.MustCompile(`([0-9]+), ([0-9]+)`)
	var match []string = r.FindStringSubmatch(input)
	var x, _ = strconv.Atoi(match[1])
	var y, _ = strconv.Atoi(match[2])
	return x, y
}

func euclDist(xi int, yi int, xj int, yj int) int {
	return int(math.Sqrt((math.Pow(float64(xi) - float64(xj), 2)) + (math.Pow(float64(yi) - float64(yj), 2))))
}

func solveB (input []string, init int) (int) {
	return 0
}

func getMinMax (input []string) (int, int, int, int) {
	var minx, miny, maxx, maxy int = 100000, 100000, 0, 0
	for i := range input {
		r := regexp.MustCompile(`([0-9]+), ([0-9]+)`)
		var match []string = r.FindStringSubmatch(input[i])
		var x, _ = strconv.Atoi(match[1])
		var y, _ = strconv.Atoi(match[2])

		if x < minx { minx = x }
		if y < miny { miny = y }
		
		if x > maxx { maxx = x }
		if y > maxy { maxy = y }
	}

	return minx, miny, maxx, maxy
}

func sortCords (input []string) []string {
	sort.Slice(input, func(i, j int) bool {
		r := regexp.MustCompile(`([0-9]+), ([0-9]+)`)
		var matchI []string = r.FindStringSubmatch(input[i])
		var matchJ []string = r.FindStringSubmatch(input[j])
		var xi, _ = strconv.Atoi(matchI[1])
		var yi, _ = strconv.Atoi(matchI[2])
		var xj, _ = strconv.Atoi(matchJ[1])
		var yj, _ = strconv.Atoi(matchJ[2])
		
		if(xi == xj){
			return yi < yj
		}

		return xi < xj
	})

	return input
}