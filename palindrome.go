package main

import (
	"fmt"
) 

func isPalindrome(x int) bool {
    
    clone := x
    reverse := 0
    
    if x < 0 {return false} //this wants a -number to be reversed from lets say -101 to 101- therefore not beinga palindrome, casting to string  wouldve been ok too
    if x == 0 {return true}
    
    for x != 0{
        
        remainder := x % 10  
        reverse = reverse * 10 + remainder 
        x /= 10
        
    }
    
    if(reverse == clone){return true} else {return false}
    
}


func main() {

	x := -10101 
	y := 10101 
	z := x + y 

	vibecheck1 := isPalindrome(x)  
	vibecheck2 :=  isPalindrome(y)
	vibecheck3 := isPalindrome(z) 


	fmt.Println(x, " is a palindrome: ", vibecheck1)
	fmt.Println(y, " is a palindrome: ", vibecheck2) 
	fmt.Println(z, " is a palindrome: ", vibecheck3) 

}
