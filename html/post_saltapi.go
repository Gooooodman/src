package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Server struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Eauth    string `json:"eauth"`
}

const (
	username = "saltapi"
	password = "saltapi"
	eauth    = "pam"
	url      = "https://139.196.231.187:8888/"
)

func main() {
	S := &Server{Username: username, Password: password, Eauth: eauth}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	http := &http.Client{Transport: tr}
	b, err := json.Marshal(S)
	//fmt.Println(b)
	if err != nil {
		fmt.Println("json err:", err)
	}

	body := bytes.NewBuffer([]byte(b))
	//fmt.Println(body)
	res, err := http.Post(url+"login", "application/x-yaml;charset=utf-8", body)
	if err != nil {
		log.Fatal(err)
		return
	}
	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("%s", result)
}
