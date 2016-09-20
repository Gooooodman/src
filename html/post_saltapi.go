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

type Ret struct {
	// Perms  []string
	// Start  float64
	Token string
	// Expire float64
	// User   string
	// Eauth  string
}

type Obj struct {
	Return []Ret `json:"return"` // json  解码可以不用
}

func main() {
	S := &Server{Username: username, Password: password, Eauth: eauth}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	http := &http.Client{Transport: tr}
	b, err := json.Marshal(S)
	if err != nil {
		fmt.Println("json err:", err)
	}
	body := bytes.NewBuffer([]byte(b))
	res, err := http.Post(url+"login", "application/json;charset=utf-8", body)
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
	//fmt.Printf("%s\n", result)

	//******************************************************************//

	var obj Obj
	fmt.Printf("%s\n", result)
	json.Unmarshal([]byte(result), &obj)
	fmt.Println(obj.Return[0].Token)
	// o := obj["return"]
	//******************************************************************//
	// // m := obj.(map[string]interface{})
	// // fmt.Println(m["return"])
	// // json.Unmarshal([]byte(m["return"]), &obj)
	//******************************************************************//
	// //断言
	// m = obj.(map[string]interface{})
	// fmt.Println(m)
	// // for i, v := range m["return"] {
	// // 	fmt.Println(i, v)
	// // }
	// var s interface{}
	// json.Unmarshal([]byte(o), &s)
	// s = obj.([]map[string]string)
	// fmt.Println(s)

}
