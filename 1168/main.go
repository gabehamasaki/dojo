package main

import (
	"fmt"
	"strconv"
)

// https://judge.beecrowd.com/pt/problems/view/1168
func main() {
	lines := 0

	leds := map[int]int{
		1: 2,
		2: 5,
		3: 5,
		4: 4,
		5: 5,
		6: 6,
		7: 3,
		8: 7,
		9: 6,
		0: 6,
	}

	fmt.Scanln(&lines)

	for i := 0; i < lines; i++ {
		var n string
		fmt.Scanln(&n)

		total := 0
		for _, r := range n {
			v := string(r)
			index, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}

			total += leds[index]
		}
		fmt.Println(total, "leds")
	}
}
