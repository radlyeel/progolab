package main 

import ( 
    "fmt" 
    //"math/rand" 
) 

func main() { 
    // Understanding iota
    const  (
        Watersports = iota 
        Soccer 
        Chess 
    )
    fmt.Println( Watersports, Soccer, Chess )
}

