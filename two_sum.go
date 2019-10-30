package main

import (
	"fmt"
)


func twoSum(nums []int, target int) []int { //returns the array index
    
    numMap := make(map[int]int, len(nums))
    
    
    for i := 0; i < len(nums); i++ { 
        
        if compliment, ok := numMap[target - nums[i]]; ok { 
            
            return []int {compliment, i}
 
        }
        numMap[nums[i]] = i
    }
    
    return []int{0,0}
}



func main() { 

	arr := []int{2,8,9,0,10,-6,18} 
	solution := twoSum(arr, 4)
	fmt.Println("answer:", solution)
}
