package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("please input:")
	a := bufio.NewReader(os.Stdin)
	input, _ := a.ReadString('\n')
	inputint, _ := strconv.Atoi(input)
	number := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, numbers := range number {
		if inputint%numbers != 0 {
			fmt.Println("不是素数")
			break
		} else {
			fmt.Println("是素数")
			break
		}
	}
	return
}
