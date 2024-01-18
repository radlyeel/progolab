package main 

import (
    "fmt" 
    "math"
)

func main() { 
    posResult := 3 % 2 
    negResult := -3 % 2 
    absResult := math.Abs(float64(negResult)) 

    fmt.Println(posResult) 
    fmt.Println(negResult) 
    fmt.Println(absResult)

}

