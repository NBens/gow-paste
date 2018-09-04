package main

import (
	"errors"
	"strings"
	"net/http"
	"net/url"
	"io/ioutil"
	"strconv"
)


/*
* Paste Structure
* Paste.privacy: 0=public 1=unlisted 2=private
* Paste.expirationDate:  // N: Never, 5M: 5 Minutes, 5H: 5Hours, 2D: 2 Days
* Paste.format: Syntax highlight format
* check https://pastebin.com/api/ for more details about privacy and expirationDate.
*/


type Paste struct {
	apiKey string         // Required Parameter
	title string          // Optional Parameter
	text string           // Required Parameter
	privacy int           // Optional Parameter
	format string         // Optional Parameter
	expirationDate string // Optional Parameter
	pasteID string        // Optional Parameter (Required to delete a paste)
}

/*
* Method responsible for creating a new paste, returns the paste's url (string)
*/


func (p *Paste) NewPaste() (string, error) {

	apiOption := "paste"

	data_values := url.Values{
		"api_dev_key": {p.apiKey},
		"api_option": {apiOption},
		"api_paste_code": {p.text},
		"api_paste_name": {p.title},
		"api_expire_date": {p.expirationDate},
		"api_paste_format": {p.format},
		"api_paste_private": {strconv.Itoa(p.privacy)},
	}
	data, httpResponseError := http.PostForm(API_POST, data_values)

	if strings.TrimSpace(p.text) == "" {
		err := errors.New("Empty text was given")
		return "", err
	} 

	if httpResponseError != nil {
		return "", httpResponseError
	}
	
	if data.StatusCode != 200 {
		return "", errors.New("Couldn't create the paste: " + strconv.Itoa(data.StatusCode))
	}

	defer data.Body.Close()
	
	returnedData, err := ioutil.ReadAll(data.Body)
	if err != nil {
		return "", err
	}
	return string(returnedData), nil

}


func (p *Paste) PrivatePaste(userKey string) (string, error) {

	apiOption := "paste"
	data_values := url.Values{
		"api_dev_key": {p.apiKey},
		"api_user_key": {userKey},
		"api_option": {apiOption},
		"api_paste_code": {p.text},
		"api_paste_name": {p.title},
		"api_expire_date": {p.expirationDate},
		"api_paste_format": {p.format},
		"api_paste_private": {"2"},
	}
	data, httpResponseError := http.PostForm(API_POST, data_values)

	if httpResponseError != nil {
		return "", httpResponseError
	}
	
	if data.StatusCode != 200 {
		return "", errors.New("Couldn't list the pastes: " + strconv.Itoa(data.StatusCode))
	}

	defer data.Body.Close()
	
	returnedData, err := ioutil.ReadAll(data.Body)
	if err != nil {
		return "", err
	}

	return string(returnedData), nil

}


func (p *Paste) Delete(userKey string) (string, error) {

	apiOption := "delete"
	data_values := url.Values{
		"api_dev_key": {p.apiKey},
		"api_user_key": {userKey},
		"api_option": {apiOption},
		"api_paste_key": {p.pasteID},
	}
	data, httpResponseError := http.PostForm(API_POST, data_values)

	if httpResponseError != nil {
		return "", httpResponseError
	}
	
	if data.StatusCode != 200 {
		return "", errors.New("Couldn't delete the paste: " + strconv.Itoa(data.StatusCode))
	}

	defer data.Body.Close()
	
	returnedData, err := ioutil.ReadAll(data.Body)
	if err != nil {
		return "", err
	}

	return string(returnedData), nil

}
