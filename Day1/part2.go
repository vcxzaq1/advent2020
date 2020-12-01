package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
    "strconv"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

    nums := make([]int, 0)
    for scanner.Scan() {
       num, _ := strconv.Atoi(scanner.Text())
       nums = append(nums, num)
	}

    exit:
    for _, num := range nums {
        for _, innerNum := range nums {
            for _, innerInnerNum := range nums {
                if num + innerNum + innerInnerNum == 2020 {
                    fmt.Println(num * innerNum * innerInnerNum)
                    break exit
                }
            }
        }
    }
}
