package main

import (
	"fmt"
	"time"
)


func main() {
	start := time.Now()
	x := 123412412341234
	// catch single digits and negatives
	if x < 0 {
		fmt.Println(false)
} else if x < 10 {
		fmt.Println(true)
} else  {

	// set bits to 1 temp to x

	var reverse int 
	// saves foward number
	foward := x 

	//generation of reverse number
	for x > 0 {
		// takes last digit via mod, adds to y which is shifted by 10 every iteration
		reverse = reverse*10 + x%10
		// removes last digit by dividing by 10 
		x /= 10
	}
	// checks if foward and reverse are the same
	fmt.Println( foward == reverse )
}
	fmt.Println(time.Since(start))


	

}

