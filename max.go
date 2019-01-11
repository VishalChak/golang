package main
import "fmt"

func max (num1, num2 int) int{
	var res int
	if (num1 > num2){
	res = num1
	} else {
	res = num2
	}
	
	return res
}

func main(){
	var a int = 100
	var b int = 200
	var ret  int = max(a,b)
	
	fmt.Println(ret)
}