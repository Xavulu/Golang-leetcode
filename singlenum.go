package main

import (
	"fmt"
)


func singleNumber(nums []int) int { 
    
    result := 0  
    
    for i := range nums { result ^= nums[i]} //XOR out all the paired numbers with ^=  
    
    return result
}// has a O(n) runtime (linear) as per task request

func main() {

	test := []int{1,1,2,2,3,3,4,5,5} 
	single := singleNumber(test)
	fmt.Println("The single number was: ", single)

}

