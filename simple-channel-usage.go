package main

import "fmt"

func main(){
  
  channel:=make(chan string)//this creates a channel which accepts/operates on values of type string
  
  go func(){//anonymous function as gorountine
  channel<-"data" //string is being put into the channel
  }()
  
  out:=<-channel
  
  fmt.Println(out)
  
}
