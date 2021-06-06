package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	UserId   int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func main() {
	//send request and get response
	url := "https://jsonplaceholder.typicode.com/users"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return
	}
	//response body convert to bytes
	jsonData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	// bytes from response convert to needed struct
	var UsersData []User
	err = json.Unmarshal(jsonData, &UsersData)
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, v := range UsersData {
		fmt.Println(v)
		//v.ShowInfo()
	}

}

func (u *User) ShowInfo() {
	message := fmt.Sprintf("Id:%d, Name:%s", u.UserId, u.Name)
	fmt.Println(message)
}
