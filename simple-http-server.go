package main

import(
"fmt"
"net/http"
)

func hello(responseWriter http.ResponseWriter, request *http.Request){
  
  
  fmt.Fprintf(responseWriter,"hello/n")  //this line writes "hello" onto to the /hello endpoint's page
}

func main(){
  http.HandleFunc("/hello",hello)//hello function is called as handler function when the /hello endpoint is used
  
  http.ListenAndServe(":8090",nil)//It listens on the port 8090 for incoming connections
}
