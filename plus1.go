package main

import (
	"fmt"
)

func plusOne(digits []int) []int { 
    
    num := len(digits) - 1
    
    for i := num; i >= 0; i-- {
        
        if digits[i] < 9 {digits[i] += 1; return digits} else{digits[i] = 0} //adds one to end of array if index content is less than resets otherwise 
        
    } 
    
    return append([]int{1}, digits...)
    
}

func main() {

	arr := []int{1,2,3,4,5,6}//think of it as 123456 
	appended := plusOne(arr)
	for i := range appended{ 

		fmt.Println("",appended[i])
	}
	//now its 123457, ta da!!! 
}
