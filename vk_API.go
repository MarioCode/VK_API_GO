package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
)

const API_VK_URL = "https://api.vk.com/method/"
const API_METHOD = "users.get"
const ACESS_TOKEN = "YOUR TOKEN"

type UserData struct {
	Id      int    `json:"uid"`
	Name    string `json:"first_name"`
	Surname string `json:"last_name"`
	City    int    `json:"city"`
}

var vk_Resp struct {
	Users []UserData `json:"response"`
}

func main() {
	start_id := 1
	step_parse := 5
	var wg sync.WaitGroup

	for j := 0; j < 1; j++ {
		for i := 0; i < 3; i++ {
			wg.Add(1)

			go func(id1, id2 int) {
				defer wg.Done()
				vk_get_users(id1, id2)
			}(start_id, start_id+step_parse)
			start_id = start_id + step_parse
		}
		wg.Wait()
	}
}

func vk_get_users(start_ID, end_ID int) {

	values := make(url.Values)
	postUrl := API_VK_URL + API_METHOD
	get_users := []string{}

	for i := start_ID; i < end_ID; i++ {
		get_users = append(get_users, fmt.Sprintf("%d", i))
	}

	fmt.Println(start_ID, end_ID)
	values.Set("user_ids", strings.Join(get_users, ","))
	values.Set("fields", "city")
	values.Set("access_token", ACESS_TOKEN)

	resp, err := http.PostForm(postUrl, values)
	defer resp.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(body, &vk_Resp)

	//users := vk_Resp.Users
	// for _, u := range users {
	// 	fmt.Printf("ID: %d, %s %s\n", u.Id, u.Name, u.Surname)
	// }
}
