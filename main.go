package main

import (
	"fmt"
	"os"
	"errors"
	"strings"
	"net/http"
	"net/url"
	"io/ioutil"
	"strconv"
)

const (
	API_URL = "https://pastebin.com/api/api_post.php"
)

/*
* Paste Structure
* Paste.privacy: 0=public 1=unlisted 2=private
* Paste.expirationDate:  // N: Never, 5M: 5 Minutes, 5H: 5Hours, 2D: 2 Days
* check https://pastebin.com/api/ for more details about privacy and expirationDate.
*/


type Paste struct {
	apiKey string
	title string
	text string
	privacy int 
	expirationDate string
	api_option string
}

/*
* Method responsible for creating a new paste, returns the paste's url (string)
*/


func (p *Paste) NewPaste() (string, error) {

	p.api_option = "paste"

	data_values := url.Values{
		"api_dev_key": {p.apiKey},
		"api_option": {p.api_option},
		"api_paste_code": {p.text},
		"api_paste_name": {p.title},
		"api_expire_date": {p.expirationDate},
		"api_paste_private": {strconv.Itoa(p.privacy)},
	}
	data, httpResponseError := http.PostForm(API_URL, data_values)

	if strings.TrimSpace(p.text) == "" {
		err := errors.New("Empty text was given")
		return "", err
	} 

	if httpResponseError != nil {
		return "", httpResponseError
	}
	
	if data.StatusCode != 200 {
		return "", errors.New("Couldn't create the paste")
	}

	defer data.Body.Close()
	
	returnedData, err := ioutil.ReadAll(data.Body)
	if err != nil {
		return "", err
	}
	return string(returnedData), nil

	
}




func main() {


	p := &Paste{
		apiKey: os.Getenv("PASTEBIN_DEVELOPER_KEY"),
		title: "Title",
		text: "Text",
		expirationDate: "N", 
		privacy: 0,
	}

	newPaste, err := p.NewPaste()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(newPaste)
	}

}