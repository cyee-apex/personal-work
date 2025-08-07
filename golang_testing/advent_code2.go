package main

import (
	"os"
	"fmt"
	"strings"
	"bufio"
)

func main() {
	moves := map[string] string {
		"A": "X", //rock
		"B": "Y", //paper
		"C": "Z", //scissor
	}

	points:= map[string] int {
		"X": 1, 
		"Y": 2, 
		"Z": 3,
	}

	file, _ := os.Open("input2.txt")
	scanner := bufio.NewScanner(file)
	var count = 0

	for scanner.Scan() {

		var line = scanner.Text()
		var pair = strings.Fields(line)

		var opp = pair[0]
		opp = moves[opp]
		var mine = pair[1]
		
		if opp == mine{
			count += 3
			count += points[mine]
		} else {
			if mine == "X" && opp == "Y" || mine == "Y" && opp == "Z" || mine == "Z" && opp == "X" {
				count += points[mine]
			}
			if mine == "X" && opp == "Z" || mine == "Y" && opp == "X" || mine == "Z" && opp == "Y" {
				count += 6
				count += points[mine]
			}
		}
	}
	fmt.Println(count)
}

//part1 ans = 12855