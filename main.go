package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/get", handleGet)
	http.HandleFunc("/get_json", handleGetJson)
	http.HandleFunc("/post", handlePost)
	http.HandleFunc("/tugas", handleTugas)

	fmt.Println("SERVING ... ")
	http.ListenAndServe(":8181", nil)
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

func handleTugas(w http.ResponseWriter, r *http.Request) {
	jsonData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	var treq TugasReq
	err = json.Unmarshal(jsonData, &treq)
	if err != nil {
		log.Fatal(err)
		return
	}

	var result string
	k := 1
	for i := 0; i < treq.MaxHeight; i++ {
		for j := 0; j <= i; j++ {
			result = fmt.Sprintf("%s%d", result, k)
			k++
			if k > treq.MaxNum {
				k = 1
			}
		}
		result = fmt.Sprintf("%s\n", result)
	}

	resp := TugasResp{
		Response: result,
	}

	data, err := json.Marshal(resp)
	if err != nil {
		log.Fatal(err)
		return
	}

	w.Write(data)
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DOR"))
}

func handleGetJson(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()

	name := queryValues.Get("name")
	idStr := queryValues.Get("id")
	id, _ := strconv.Atoi(idStr)
	alias := queryValues.Get("alias")

	p := People{
		ID:        id,
		Name:      name,
		AliasName: alias,
	}

	data, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
		return
	}

	w.Write(data)
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	jsonData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	var p People
	err = json.Unmarshal(jsonData, &p)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("%+v", p)

	newP := People{
		ID:        p.ID + 1000,
		Name:      fmt.Sprintf("NEW-%s", p.Name),
		AliasName: fmt.Sprintf("NEW-%s", p.AliasName),
	}

	data, err := json.Marshal(newP)
	if err != nil {
		log.Fatal(err)
		return
	}

	w.Write(data)
}
