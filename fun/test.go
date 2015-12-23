package main

import "fmt"

func sum(x int) int {
	total:=0
	for i:=0; i<10; i++{
		total+=i
	}
	return total;
}	

func add(x int, y int) int {
	return x+ y
}

func swap(x,y string) (string, string) {
	return y, x
}

func pow(x,n,lim float64) float64 {
	if v:= math.Pow(x,n); v < lim {
		return v
	}
	return lim
}

func main(){
	fmt.Println(add(10,5))
	fmt.Println(sum(10))

}
