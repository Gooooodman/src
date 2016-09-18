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
	Client   string `json:"client"`
	Fun      string `json:"fun"`
}

const (
	username = "saltapi"
	password = "saltapi"
	eauth    = "pam"
	//url      = "https://192.168.30.129:8888/"
	url = "https://139.196.231.187:8888/"
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

// ssl  忽略认证
func Client() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &http.Client{Transport: tr}

}

func init() {
	S.Username = username
	S.Password = password
	S.Eauth = eauth
	client := Client()
	b, err := json.Marshal(S)
	if err != nil {
		fmt.Println("json err:", err)
	}
	body := bytes.NewBuffer([]byte(b))
	res, err := client.Post(url+"/login", "application/json;charset=utf-8", body)
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
}

func (o Obj) ListALLKey() (string, error) {
	client := Client()
	req, err := http.NewRequest("GET", url+"/keys", nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("X-Auth-Token", obj.Return[0].Token)
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	if res.StatusCode != 200 {
		fmt.Println(res.Status)
		return "", err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func main() {
	ret, _ := obj.ListALLKey()
	fmt.Println(ret)
}
