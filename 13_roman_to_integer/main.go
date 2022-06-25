package main

import (
	"fmt"
	"time"
)

// accidently made this function convert roman numerals in any order to integers assuming the pairs are in a set
func main() {
	start := time.Now()
	m := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000, "IV": 4, "IX": 9, "XL": 40, "XC": 90, "CD": 400, "CM": 900}
	s := "IILLCDXMIXLXLXCCDMCMI"
	sum := 0
	for i := 0; i < len(s); {
			x := string(s[i])
		if i < len(s)-1 {
			y := string(s[i+1])
			set := (x+y)
			if _, ok := m[set]; ok {
				sum += m[set]
				i += 2
			} else {
				sum += m[x]
				i++
				}
			} else {
			sum += m[x]
			break
		}
		

	}

	fmt.Println("input:", s,"sum:", sum, "time:", time.Since(start))


	

}

