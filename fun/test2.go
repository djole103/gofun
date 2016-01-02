package main 

import "fmt"

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func main(){
a := make([]int,5,10) //type,length,capacity
a= append(a,1)
printSlice("a",a)
for i,_:= range a {
	a[i]++
	}
printSlice("a",a)
}
