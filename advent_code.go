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
	file, err := os.Open("input.txt")
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
	// 	log.Fatal(err)
	// }
	sort.Ints(s)
	sort.Reverse(s)
	fmt.Println(s[0])
	fmt.Println(s[1])
	fmt.Println(s[2])

}

//ans day1 68292

