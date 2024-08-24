package main

import (
	"fmt"
)

func main() {
	var slice = []int64{7,90,213,90,1,-9}
	for {
		stop := true
		for i,val := range slice {
			if i >= len(slice)-1 {
				break
			}
			if val > slice[i+1]  {
				slice[i] = slice[i+1]
				slice[i+1] = val
				stop = false
			}
		}
		if stop {goto ComeToMe}	
	}
	ComeToMe:
	fmt.Println(slice)
}	