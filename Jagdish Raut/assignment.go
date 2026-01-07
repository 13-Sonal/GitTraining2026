package main

import(
	"fmt"
)

func reverse(str string) string{

	var  reverse_str string
	
	for i:=len(str)-1 ;i >= 0; i--{
		
		reverse_str += string(str[i])
	}

return reverse_str
	

}



func main(){
	

	var str string
	fmt.Print("Enter a string :")

	fmt.Scan(&str)

	fmt.Println("reverse of string is :", reverse(str))
	
}