package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func add(x, y int) int{
  return x+y
}
func sub(x, y int) int{
  return x-y
}
func divide(x, y int) int{
  return x/y
}
func multiply(x, y int) int{
  return x*y
}
func main(){
  var x, y int
  fmt. Println("Enter First Number: ")
  fmt.Scanln(&x)
  fmt.Println("Enter Second Number: ")
  fmt.Scanln(&y)
  
  reader := bufio.NewReader(os.Stdin)
  fmt.Println("What operation do you want to perform: ")
  operation, _ := reader.ReadString('\n')
  operation = strings.TrimSpace(operation)
  if operation == "add"{
    fmt.Println("Result is ", add(x,y))
  }else if operation == "sub"{
    fmt.Println("Result is",sub(x,y))
  }else if operation == "divide"{
    fmt.Println("Result is ", divide(x,y))
  }else if operation == "multiply"{
    fmt.Println("Result is ", multiply(x,y))
  }else{
    fmt.Println("Invalid Operation")
  }
}