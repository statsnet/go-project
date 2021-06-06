package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  struct {
		Street  string `json:"street"`
		Suite   string `json:"suite"`
		City    string `json:"city"`
		Zipcode string `json:"zipcode"`
		Geo     struct {
			Lat string `json:"lat"`
			Lng string `json:"lng"`
		} `json:"geo"`
	} `json:"address"`
	Phone   string `json:"phone"`
	Website string `json:"website"`
	Company struct {
		Name        string `json:"name"`
		CatchPhrase string `json:"catchPhrase"`
		Bs          string `json:"bs"`
	} `json:"company"`
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
