package main 

import ( 
    "fmt" 
    //"math/rand" 
) 

func main() { 
    var price float32 = 275.00 
    var tax float32 = 27.50 
    fmt.Println(price + tax) 
    price = 300 
    fmt.Println(price + tax)
}

