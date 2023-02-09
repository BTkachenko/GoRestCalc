package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Operation struct {
	Number1 int64 `json:"number1"`

	Number2 int64 `json:"number2"`
}

type Result struct {
	Result int64 `json:"result"`
}

func main() {
	var number1, number2 int
	var op string
	fmt.Print("Enter number1: ")
	fmt.Scanf("%d", &number1)

	fmt.Print("Enter number1: ")
	fmt.Scanf("%d", &number2)

	fmt.Print("Enter operator (+,/,-,*): ")
	fmt.Scanf("%s", &op)

	data := Operation{
		Number1: int64(number1),

		Number2: int64(number2),
	}

	var link string
	switch op {
	case "+":
		link = "http://localhost:8080/math/add"
	case "-":
		link = "http://localhost:8080/math/sub"
	case "/":
		link = "http://localhost:8080/math/div"
	case "*":
		link = "http://localhost:8080/math/multi"
	}

	// Encode the data
	postBody, _ := json.Marshal(data)
	responseBody := bytes.NewBuffer(postBody)
	// Leverage Go's HTTP Post function to make request
	resp, err := http.Post(link, "application/json", responseBody)
	// Handle Error
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	// defer resp.Body.Close()
	// Read the response body

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Println(sb)

	var res Result
	json.Unmarshal(body, &res)

	println("Wynik w struct odtworzonego z jsona ", res.Result)
}
