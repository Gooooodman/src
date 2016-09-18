package main

import (
	"bytes"
	"crypto/tls"
	//"encoding/base64"
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
	Client   string `json:"client"`
	Fun      string `json:"fun"`
}

//type Server1 struct {
// 	Client string `json:"client"`
// 	Fun    string `json:"fun"`
// }

const (
	username = "saltapi"
	password = "saltapi"
	eauth    = "pam"
	url      = "https://192.168.30.129:8888/"
	//url = "https://139.196.231.187:8888"
)

type Ret struct {
	Perms  []string
	Start  float64
	Token  string
	Expire float64
	User   string
	Eauth  string
}

type Obj struct {
	Return []Ret `json:"return"` // json  解码可以不用
}

var obj Obj

var S Server

func init() {
	//S = &Server{Username: username, Password: password, Eauth: eauth}
	S.Username = username
	S.Password = password
	S.Eauth = eauth
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	http := &http.Client{Transport: tr}
	b, err := json.Marshal(S)
	if err != nil {
		fmt.Println("json err:", err)
	}
	body := bytes.NewBuffer([]byte(b))
	res, err := http.Post(url+"/login", "application/json;charset=utf-8", body)
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

	json.Unmarshal([]byte(result), &obj)
	//fmt.Println(obj)
}

func (o Obj) ListALLKey() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("GET", url+"/keys", nil)
	if err != nil {
		fmt.Println("request  -- >", err)
		return
	}
	req.Header.Add("X-Auth-Token", obj.Return[0].Token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("clint: -- > ", err)
		return
	}
	if res.StatusCode != 200 {
		fmt.Println(res.Status)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("read:   -- >", err)
	}
	fmt.Println(string(body))
}

func main() {
	fmt.Println(obj.Return[0].Token)
	obj.ListALLKey()
}
