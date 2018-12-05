package main

import (
	"strconv"
	"log"
	"fmt"
	"sort"
	"regexp"
	"github.com/gramstrong/aoc2018/utils"
)

func main() {
	var input, err = utils.Read("input.txt")
	if err != nil {
    log.Fatal(err)
	} else {
		fmt.Printf("A: %d\n", solveA(input));
		fmt.Printf("B: %d", solveB(input, 0));
	}
}

func solveA (input []string) (int) {
	var guards = processGuards(input)
	var max int = 0;
	var maxID int = 0;

	for i := range guards {
		var curr int = 0;
		for j := range guards[i] {
			curr += guards[i][j]
		}
		if curr > max {
			max = curr
			maxID = i
		}
	}

	var maxMinute int = 0;
	var maxMinuteIdx int = 0;

	for i := range guards[maxID] {
		if guards[maxID][i] >= maxMinute {
			maxMinuteIdx = i
			maxMinute = guards[maxID][i]
		}
	}

	return maxID * maxMinuteIdx;
}

func solveB (input []string, init int) (int) {
	var guards map[int][60]int = processGuards(input) 
	var maxID int = 0;
	var maxMinute int = 0;
	var maxMinuteIdx int = 0;

	for i := range guards {
		for j := range guards[i] {
			if guards[i][j] >= maxMinute {
				maxMinute = guards[i][j]
				maxID = i
				maxMinuteIdx = j
			}
		}
	}

	return maxID * maxMinuteIdx;
}

func parseBeginShift (line string, id int) int {
	bs := regexp.MustCompile(`begins shift`)
	var newId int = id;

	if bs.MatchString(line) {
		var match []string = regexp.MustCompile(`Guard \#([0-9]+)`).FindStringSubmatch(line)
		newId, _ = strconv.Atoi(match[1])
	}
	return newId
}

func parseFallsAsleep (line string, sleepIdx int) int {
	fa := regexp.MustCompile(`falls asleep`)
	var newSleepIdx int = sleepIdx;

	if fa.MatchString(line) {
		var newSleepIdxStr string
		_,_,_,_,newSleepIdxStr = getTime(line)
		newSleepIdx, _ = strconv.Atoi(newSleepIdxStr)
	}
	return newSleepIdx
}

func parseWakesUp (line string, sleepIdx int, sched [60]int) [60]int {
	wu := regexp.MustCompile(`wakes up`)
	var wakeIdx int

	if wu.MatchString(line) {
		var wakeIdxStr string
		_,_,_,_,wakeIdxStr = getTime(line)
		wakeIdx, _ = strconv.Atoi(wakeIdxStr)

		for i := sleepIdx; i < wakeIdx; i ++ {
			sched[i] += 1;
		}
	}

	return sched
}


func getTime (line string) (string, string, string, string, string) {
	r := regexp.MustCompile(`\[([0-9]{4})-([0-9]{2})-([0-9]{2}) ([0-9]{2}):([0-9]{2})\]`)
	var groups []string = r.FindStringSubmatch(line);
	return groups[1], groups[2], groups[3], groups[4], groups[5]
}

func concatTime(y string, m string, d string, h string, i string) int {
	var time string = y+m+d+h+i
	var timeInt, err = strconv.Atoi(time)

	if err != nil {
		log.Fatal(err)
	}
	return timeInt
}

func processGuards (input []string) map[int][60]int {
	sort.Slice(input, func(i, j int) bool {
		var timeI int = concatTime(getTime(input[i]))
		var timeJ int = concatTime(getTime(input[j]))
		return timeI < timeJ
	})
	
	var id int = 0;
	var sleepIdx = 0;
	var guards map[int][60]int = make(map[int][60]int)

	for i := range input {
		id = parseBeginShift(input[i], id)
		sleepIdx = parseFallsAsleep(input[i], sleepIdx)
		guards[id] = parseWakesUp(input[i], sleepIdx, guards[id])
	}

	return guards;

}
