package main

import (
	"fmt"
	"strings"
	//"strconv"
	//  "math"
)

func Zero(n string){
  
	Convert := strings.Split(n,".")
	//fmt.Println(Convert)
	x := strings.Join(Convert,"")
//	fmt.Println(x)
	y:=len(x)-1
//	fmt.Println(y)

	for _, Set:=range x{
	// fmt.Println(Set)
	 StrAppend := []string{}
		for i:=0; i < y;i++{
			GetZero := "0"
			StrAppend = append(StrAppend,GetZero)
		}
		y--
		
	   Result := string(Set) + strings.Join(StrAppend,"")
		
		fmt.Println(Result)		
				
		
	}

		
  
}

func main() {
	Zero("1.345.679")
	fmt.Println("Hello, playground")
}
