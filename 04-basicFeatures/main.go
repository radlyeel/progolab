package main 

import ( 
    "fmt" 
    //"math/rand" 
) 

func main() { 
    const price, tax float32 = 275, 27.50 
    const quantity, inStock = 2, true 
    fmt.Println("Total:", 2 * quantity * (price + tax)) 
    fmt.Println("In stock: ", inStock)
}

