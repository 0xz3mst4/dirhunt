package main

import (
	// "io/ioutil"
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func makeRequest(){
	file, err := os.Open("/home/lists/big.txt")
	defer file.Close()
	args := os.Args

	r := bufio.NewReader(file)

	for{
		line, _, err := r.ReadLine()
		if len(line) > 0{
			response, err := http.Get(args[1] + line);

		}
	}

	if err != nil{
		log.Fatal(err)
		return
	}
	statusCode := response.StatusCode
	fmt.Printf("Status Code: %d\n", statusCode)
	defer response.Body.Close();
}

func main(){
	makeRequest();
}
