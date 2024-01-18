package main 

import (
    "fmt" 
    "math"
)

func main() { 
    var intVal = math.MaxInt64 
    var floatVal = math.MaxFloat64 
    fmt.Println(intVal * 2) 
    fmt.Println(floatVal * 2) 
    fmt.Println(math.IsInf((floatVal * 2), 0))
}

