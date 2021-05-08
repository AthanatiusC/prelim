package main

import (
	"fmt"
	"strings"
)

func reverse_int(n int) int {
	new_int := 0
	for n > 0 {
		remainder := n % 10
		new_int *= 10
		new_int += remainder
		n /= 10
	}
	return new_int
}

func socks() {
	inputs := []int{10, 20, 20, 10, 10, 30, 50, 10, 20}
	sets := make(map[int]int)
	for _, val := range inputs {
		sets[val] = sets[val] + 1
	}
	for _, val := range sets {
		sets[val] = sets[val] + 1
	}
	for _, val := range inputs {
		delete(sets, val)
	}
	fmt.Println(len(sets))
}

func valley() {
	inputs := "U D D D U D U U"
	s := strings.Split(inputs, " ")
	valleys := 0
	for idx, val := range s {
		if idx == 0 {
			continue
		}
		if val == s[idx-1] {
			if val == "U" {
				valleys += 1
			}
		}
	}
	fmt.Println(valleys)
}

func pyramid() {
	inputs := []int{1, 3, 4, 5, 6, 7, 9}
	for idx, input := range inputs {
		fmt.Printf("%d%0*d\n", input, len(inputs)-idx-1, 0)
	}
}

func light() {
	total := 0
	num := 0
	for i := 0; i < 100; i++ {
		if i == 0 {
			num = 100
		} else {
			num = 100 / i
		}
		total += num
	}
	fmt.Println("Total lamp turned on : ", total)
}

func main() {
	socks()
	pyramid()
	valley()
	light()
}
