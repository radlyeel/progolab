package main 

import ( 
    "fmt" 
) 

func main() { 
    first := 100 
    var second *int = &first 

    first++ 

    fmt.Println("First:", first) 
    fmt.Println("Second:", *second)

}

