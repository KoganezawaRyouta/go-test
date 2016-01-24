package main

import "fmt"
import "math"

/*
   *
  ***
 *****
*******
 *****
  ***
   *
*/

var outputStr = "*"
var space = " "
var max = 7

func main() {
	middleLine := calcMiddleLine()

	for x := 1; x <= max; x++ {
		for y := 1; y <= max; y++ {
			if middleLine == x {
				fmt.Printf(outputStr)
			}

			if (middleLine - x) > 0 {
				for y := 1; y <= (middleLine - x); y++ {
					fmt.Printf(space)
				}
				fmt.Printf(outputStr)
			}
		}
		fmt.Println()
	}
}

func calcMiddleLine() int {
	middle := (math.Trunc(float64(max) / 2))
	if (int(middle) % 2) != 0 {
		middle++
	}
	return int(middle)
}
