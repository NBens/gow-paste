package main

import (
	"fmt"
	"os"
	. "github.com/NBens/gow-paste"
)

func main() {

	developerKey := os.Getenv("PASTEBIN_DEVELOPER_KEY")
	

	p := &Paste{
		apiKey: developerKey,
		title: "Title",
		text: "echo 'Hello World';",
		expirationDate: "N", 
		format: "PHP",
		privacy: 0,
	}

	u := &User{
		apiKey: developerKey,
		username: "USERNAME",
		password: "PASSWORD",
	}

	u.GenerateKey()

	newPrivate, err := p.PrivatePaste(u.userKey)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(newPrivate)
	}

	list, err := u.List(10)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(list)
	}


}