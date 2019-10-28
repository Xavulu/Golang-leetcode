package main 

import ( 
	"fmt" 
) 

func mySqrt(x int) int {
	//square root via binary search, O(log(n)) run time, results will be rounded off            
    low := 0 
    max := x   
    sqrt := 0 
    if x == 1 || x == 0 {return x} 
    for low <= max{ 
    
        middle := (low + max) / 2 
        if middle * middle == x{ return middle } 
        if middle * middle < x {low = middle + 1; sqrt = middle} else { max = middle - 1} 
        
    } 
    
    return sqrt
}

func main() {

	x := 8 
	square := mySqrt(x) 

	//the main is just here to test and i didnt actually include comments in the leetcode submission 
	fmt.Println("The square root of ", x, " is ", square)


}
