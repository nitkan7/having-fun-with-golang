//though this is a simple program...having the hello world program as the first one in the repo, I thought would be very nice

package main

import ("fmt")

func main(){
  //Ever wondered why the P in Println is in capitals ?
  //It is because the Println function is exported from the fmt package and all functions and variables that are to be exported are in capitals in golang.
  fmt.Println("Hello, people!")
}
