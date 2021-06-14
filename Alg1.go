package main

import (
  "fmt"
 // "strings"
//"strconv"
//  "math"
  )

func SameSock(n int32, arr []int32)int32 {
  
	 CountSame := make(map[int32]int)
  	//fmt.Print(CountSame)

    for i := 0; i < int(n); i++ {
                    // fmt.Print(CountSame[arr[i]])

        CountSame[arr[i]] =  CountSame[arr[i]] + 1

    }
       // fmt.Print(CountSame)

    Result:= 0
    for _, value := range CountSame {
       // fmt.Println(value)
        
       Result = Result + value / 2
           //fmt.Print(sum)

    }
     fmt.Print(Result)


    return int32(Result)

		
  
}

func main() {
	var val = []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}
	SameSock(9,val)
	fmt.Println("Hello, playground")
}
