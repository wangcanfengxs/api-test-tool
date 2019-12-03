package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/testGet", testGet)
	http.HandleFunc("/testPost", testPost)
	http.HandleFunc("/testPut", testPut)
	http.HandleFunc("/testDelete", testDelete)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
		return
	}
}

func testDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		w.WriteHeader(405)
		return
	}
	fmt.Println(formatHttpRequest(r))
	w.WriteHeader(200)
	w.Write([]byte(`
        {
            "code": 200,
            "message": "SUCCESS"
        }
    `))
}

func testPut(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		w.WriteHeader(405)
		return
	}
	fmt.Println(formatHttpRequest(r))
	w.WriteHeader(200)
	w.Write([]byte(`
        {
            "code": 200,
            "message": "SUCCESS"
        }
    `))
}

func testPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(405)
		return
	}
	fmt.Println(formatHttpRequest(r))
	w.WriteHeader(200)
	w.Write([]byte(`{
        "code": 200,
        "message": "SUCCESS"
    }`))
}

func testGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(405)
		return
	}
	fmt.Println(formatHttpRequest(r))
	jsonData, err := json.Marshal(r.URL.Query())
	if err != nil {
		w.WriteHeader(200)
		w.Write([]byte("{}"))
		return
	}
	w.WriteHeader(200)
	w.Write(jsonData)
}

func formatHttpRequest(r *http.Request) string {
	first := fmt.Sprintf("requestLine: %s %s HTTP/1.1\n", r.Method, r.Host)
	header, _ := json.Marshal(r.Header)
	body, _ := ioutil.ReadAll(r.Body)
	return fmt.Sprintf("%sheader: %s\nbody: %s\n", first, header, body)
}
