package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	// Part 1
	// step := 0
	// trees := 0
	// scanner.Scan()
	// for scanner.Scan() {
	// 	step = step + 3
	// 	length := len(scanner.Text())
	// 	road := scanner.Text()
	// 	position := step % length
	// 	character := string(road[position])
	// 	if character == "#" {
	// 		trees = trees + 1
	// 	}

	// }

	var roads []string
	for scanner.Scan() {
		roads = append(roads, scanner.Text())
	}

	// Part 2
	total := a(roads, 1, 1) * a(roads, 3, 1) * a(roads, 5, 1) * a(roads, 7, 1) * a(roads, 1, 2)
	fmt.Println(total)
}

func a(roads []string, right int, down int) int {

	trees := 0
	step := 0
	for i := down; i <= len(roads)-1; i += down {
		step += right
		line := roads[i]
		length := len(line)
		position := step % length
		character := string(line[position])
		if character == "#" {
			trees = trees + 1
		}
	}
	return trees

}
