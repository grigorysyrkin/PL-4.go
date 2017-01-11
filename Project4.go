package main

import (
	"fmt"
)

//declaring token struct
type Token struct {
	data      string
	recipient int
}

//create next thread or stop
func thread(i int, t Token, c chan string) {
	if t.recipient == i {
		r := "Token has reached the recipient"
		c <- r
	} else {
		go thread(i+1, t, c)
	}
}


func main() {
	c := make(chan string)
	token := Token{data: "keyword", recipient: 20}
//starting first thread here
	go thread(1, token, c)
	fmt.Println(<-c)
}