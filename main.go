package main

import (
	"fmt"
	"os"
)

func main() {

	developerKey := os.Getenv("PASTEBIN_DEVELOPER_KEY")
	/*

	p := &Paste{
		apiKey: developerKey,
		title: "Title",
		text: "echo 'Hello World';",
		expirationDate: "N", 
		format: "PHP",
		privacy: 0,
	}

	newPaste, err := p.NewPaste()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(newPaste)
	}
	*/

	u := &User{
		apiKey: developerKey,
		username: "nizarnizario",
		password: "nizarnizar",
	}

	u.GenerateKey()

	list, err := u.List(10)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(list)
	}

}