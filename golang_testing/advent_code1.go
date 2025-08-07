package main

import (
    "os"
    "log"
	"fmt"
	"bufio"
	"strconv"
	"sort"
)



func main() {
	file, err := os.Open("input1.txt")
	scanner := bufio.NewScanner(file)

	if err != nil{
	 	log.Fatal(err)
	}

	s := make([]int, 0)
	var sum = 0

	for scanner.Scan() {
		var line = scanner.Text()
		//fmt.Println(line)
		if line == "" {
			s= append(s, sum)
			sum = 0
		} else {
			value, _ := strconv.Atoi(line)
			sum += value
		}
	}

	// var ret = 0
	// for _, v := range s {
	// 	if v > ret {
	// 		ret = v
	// 	}
    // }
	// fmt.Println(ret)
	// if err != nil{
	// 	log.Fatal(err) //for part1
	// }
	
	sort.Sort(sort.Reverse(sort.IntSlice(s)))

	fmt.Println(s[0] + s[1] + s[2]) //for part2

}

//ans day1 68292

