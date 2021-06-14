package main

import (
  "fmt"
  //"strings"
//"strconv"
//  "math"
  )

func Valley(steps int32, path string) int32 {
    // Write your code here
    
    var current,result int32 
    var Sea  = false 
    GetData:= []rune(path)
    
    for _, getValley := range GetData{
        if string(getValley) == "U" {
             fmt.Println(string(getValley))
            current++
            if current > 0{
                Sea = false
            }
        }else if string(getValley) == "D"{
            fmt.Println(string(getValley))
            current--
            if current < 0{
                Sea = true
            }
            
        }
         if current == 0 && Sea{
            result = result + 1
       

        }
    }
  
    fmt.Println(result)
    return result
}


func main() {
	Valley(8, "UDDDUDUU")
	fmt.Println("Hello, playground")
}