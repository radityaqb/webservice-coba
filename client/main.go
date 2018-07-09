package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var c *http.Client

func main() {
	client := http.Client{
		Timeout: 1 * time.Second,
	}
	c = &client

	// get()
	// post()
	posttugas()
}

type People struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	AliasName string `json:"alias_name"`
	age       int    `json:"age"`
}

type TugasReq struct {
	MaxNum    int `json:"max_num"`
	MaxHeight int `json:"max_height"`
}

type TugasResp struct {
	Response string `json:"response"`
}

func posttugas() {
	treq := TugasReq{
		MaxNum:    5,
		MaxHeight: 8,
	}

	fmt.Printf("POST REQUEST : %+v\n\n", treq)

	jsondata, err := json.Marshal(treq)
	if err != nil {
		log.Fatal(err)
		return
	}

	req, err := http.NewRequest("POST", "http://localhost:8181/tugas",
		bytes.NewBuffer(jsondata))
	if err != nil {
		log.Fatal(err)
		return
	}

	resp, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	var pResp TugasResp
	err = json.Unmarshal(respData, &pResp)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("POST RESPONSE : \n%+v\n\n", pResp)

}

func get() {
	url := "http://localhost:8181/get_json"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return
	}

	q := req.URL.Query()
	q.Add("id", "1234")
	q.Add("name", "mira")
	q.Add("alias", "tri")
	req.URL.RawQuery = q.Encode()

	resp, err := c.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	var pResp People
	err = json.Unmarshal(respData, &pResp)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("GET RESPONSE : %+v\n\n", pResp)
}

func post() {
	p := People{
		ID:        9,
		Name:      "Radit",
		AliasName: "Raditya",
	}

	fmt.Printf("POST REQUEST : %+v\n\n", p)

	jsondata, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
		return
	}

	req, err := http.NewRequest("POST", "http://localhost:8181/post",
		bytes.NewBuffer(jsondata))
	if err != nil {
		log.Fatal(err)
		return
	}

	resp, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	var pResp People
	err = json.Unmarshal(respData, &pResp)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("POST RESPONSE : %+v\n\n", pResp)

}
