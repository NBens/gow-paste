package main

import (
	"net/http"
	"net/url"
	"errors"
	"io/ioutil"
	"strconv"
)

const (
	API_LOGIN = "https://pastebin.com/api/api_login.php"
	API_POST = "https://pastebin.com/api/api_post.php"
)

type User struct {
	apiKey string
	username string
	password string
	userKey string
	api_option string
}

func (u *User) GenerateKey() (string, error) {

	data_values := url.Values{
		"api_dev_key": {u.apiKey},
		"api_user_name":  {u.username},
		"api_user_password": {u.password},
	}
	data, httpResponseError := http.PostForm(API_LOGIN, data_values)

	if httpResponseError != nil {
		return "", httpResponseError
	}
	
	if data.StatusCode != 200 {
		return "", errors.New("Couldn't generate the key: " + strconv.Itoa(data.StatusCode))
	}

	defer data.Body.Close()
	
	returnedData, err := ioutil.ReadAll(data.Body)
	if err != nil {
		return "", err
	}

	u.userKey = string(returnedData)
	return u.userKey, nil	
}


func (u *User) List(limit int) (string, error) {

	apiOption := "list"
	data_values := url.Values{
		"api_dev_key": {u.apiKey},
		"api_user_key": {u.userKey},
		"api_option": {apiOption},
		"api_results_limit": {strconv.Itoa(limit)},
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