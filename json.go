package main    
import ( 
    "fmt"
    "encoding/json"
)   
type Human struct{         
   Name string 
    Age int
    Address string 
}   
func main() { 
       
    // defining a struct instance 
    human1 := Human{"Rajasekhar", 24, "Rayachoti"} 
       
    // encoding human1 struct 
    // into json format 
    human_enc, err := json.Marshal(human1) 
       
    if err != nil { 
           
        // if error is not nil 
        // print error 
        fmt.Println(err) 
    } 
       
    // as human_enc is in a byte array 
    // format, it needs to be  
    // converted into a string  
    fmt.Println(string(human_enc)) 
       
    // converting slices from 
    // golang to JSON fomat 
       
    // defining an array 
    // of struct instance 
    human2 := []Human{ 
        {Name: "pavan", Age: 23, Address: "hyb"},
        {Name: "nagadhri", Age: 25, Address: "kadapa"}, 
        {Name: "chanti", Age: 27, Address: "Bangalore"}, 
    } 
       
    // encoding into JSON format 
    human2_enc, err := json.Marshal(human2) 
        
        if err != nil { 
       
        // if error is not nil 
        // print error 
            fmt.Println(err) 
        } 
           
    // printing encoded array 
    fmt.Println() 
        fmt.Println(string(human2_enc)) 
} 