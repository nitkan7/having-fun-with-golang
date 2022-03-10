package main

import(
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){
	response,error:=http.Get("https://animechan.vercel.app/api/random")	//simple get request

	if error != nil {	//simple error handling
		panic(error)
	}
	defer response.Body.Close()
	body,error:=ioutil.ReadAll(response.Body) //returns bytes from the response body
	if error != nil {
		panic(error)
	}
	output:=string(body) //bytes are converted to string
	fmt.Println(output)
}
