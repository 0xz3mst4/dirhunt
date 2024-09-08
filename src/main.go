package main

import (
	// "io/ioutil"
	"fmt"
	"os"
	"log"
	"net/http"
)

func makeRequest(){
	args := os.Args

	response, err := http.Get(args[1]);
	if err != nil{
		log.Fatal(err)
		return
	}
	defer response.Body.Close();
	statusCode := response.StatusCode
	fmt.Printf("Status Code: %d\n", statusCode)
}

func main(){
	makeRequest();
}
