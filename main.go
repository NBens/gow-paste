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


type Pastebin struct {
	apiKey string
	title string
	text string
	privacy int // 0=public 1=unlisted 2=private
	expirationDate string // N: Never, 5M: 5 Minutes, 5H: 5Hours, 2D: 2 Days
	api_option string
}


func (p *Pastebin) NewPaste() (string, error) {

	p.api_option = "paste"

	data_values := url.Values{
		"api_dev_key": {p.apiKey},
		"api_option": {p.api_option},
		"api_paste_code": {p.text},
		"api_paste_name": {p.title},
		"api_expire_date": {p.expirationDate},
		"api_paste_private": {strconv.Itoa(p.privacy)},
	}
	data, respErr := http.PostForm(API_URL, data_values)

	if strings.TrimSpace(p.title) == "" {
		err := errors.New("Empty title was given")
		return "", err
	} else if respErr != nil {
		return "", respErr
	} else {
		returnedData, err := ioutil.ReadAll(data.Body)
		if err != nil {
			return "", err
		}

		return string(returnedData), nil
	}

}

func main() {


	p := &Pastebin{
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