package main

import (
	"fmt"
)

func main() {

	developerKey := "Your Pastebin Developer Key" // or use os.Getenv("PASTEBIN_DEVELOPER_KEY") if you set your env variables up
	

	p := &Paste{
		apiKey: developerKey,
		title: "Title", // Paste title
		text: "echo 'Hello World';", // Paste body
		expirationDate: "N", // Expiration date (N: Never)
		format: "PHP", // Syntax Highlighting
	}

	u := &User{
		apiKey: developerKey,
		username: "USERNAME", // Set your username here
		password: "PASSWORD", // Set your password here
	}

	u.GenerateKey() // Generate Key: Generates User API Key (Which is used to delete a paste, or to create a private paste)

	newPrivate, err := p.PrivatePaste(u.userKey) // Create a Private paste : PrivatePaste takes userKey(The Generated Key)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(newPrivate)
	}

	list, err := u.List(10) // List 10 Pastes Maximum

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(list)
	}

	p.pasteID = "xA5DG0" // Set paste id in the Paste struct

	p.Delete(u.userKey) // Delete a paste: Delete takes userKey(The generated key)


}