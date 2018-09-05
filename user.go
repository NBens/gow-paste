package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

// APILogin: Url for the login api (which uses username and password to generate the user key)
// APIPost: Url for the pastes api, it is responsible for creating, listing and deleting pastes
const (
	APILogin = "https://pastebin.com/api/api_login.php"
	APIPost  = "https://pastebin.com/api/api_post.php"
)

// User structure contains information about the pastebin user(username, password) and methods to generate the api user key
type User struct {
	apiKey   string
	username string
	password string
	userKey  string
}

// GenerateKey method generates the user api key, which is used to create private pastes and to list user's pastes
func (u *User) GenerateKey() (string, error) {

	dataValues := url.Values{
		"api_dev_key":       {u.apiKey},
		"api_user_name":     {u.username},
		"api_user_password": {u.password},
	}
	data, httpResponseError := http.PostForm(APILogin, dataValues)

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

// List method, lists the pastes of a user
func (u *User) List(limit int) (string, error) {

	apiOption := "list"
	dataValues := url.Values{
		"api_dev_key":       {u.apiKey},
		"api_user_key":      {u.userKey},
		"api_option":        {apiOption},
		"api_results_limit": {strconv.Itoa(limit)},
	}
	data, httpResponseError := http.PostForm(APIPost, dataValues)

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
