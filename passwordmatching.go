package main    
import (  
    "fmt"  
    "regexp"  
)  
  
func main() {  
    re := regexp.MustCompile("f([a-z]+)ing")  
    fmt.Println(re.FindStringSubmatch("rejcating"))  
    fmt.Println(re.FindStringSubmatch("fering"))  
    fmt.Println(re.FindStringSubmatch("comforting"))  

    fmt.Println(re.FindStringSubmatch("abcfloatingxyz"))  
}  