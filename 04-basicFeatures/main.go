package main 

import ( 
    "fmt" 
    "math/rand" 
) 

func main() { 
    // Note that the current (2024-01-18) rand() defaults
    // To a random-seeded behavior
    fmt.Println("Random value: ", rand.Int()) 
    
}

