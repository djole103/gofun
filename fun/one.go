package main
import "fmt"

func work(num int)(x,y int){
	x = num +2
	y = num +4
	return
}

func main(){
	fmt.Println(work(5))
}
