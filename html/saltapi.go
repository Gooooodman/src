package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// 参数
type Server struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Eauth    string `json:"eauth"`
	Token    string `json:"token"`
	client   *http.Client
}

const (
	username = "saltapi"
	password = "saltapi"
	eauth    = "pam"
	//serverUrl      = "https://192.168.30.129:8888/"
	serverUrl = "https://139.196.231.187:8888/"
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

var S *Server = new(Server)

// ssl  忽略认证
func Client() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	S.client = &http.Client{Transport: tr}
	// }

	// func Login() {
	S.Username = username
	S.Password = password
	S.Eauth = eauth
	//client := Client()
	b, err := json.Marshal(S)
	if err != nil {
		fmt.Println("json err:", err)
	}
	body := bytes.NewBuffer([]byte(b))
	res, err := S.client.Post(serverUrl+"/login", "application/json;charset=utf-8", body)
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
	S.Token = obj.Return[0].Token
}

func (S *Server) ListALLKey() (string, error) {
	req, err := http.NewRequest("GET", serverUrl+"/keys", nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("X-Auth-Token", S.Token)
	res, err := S.client.Do(req)
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

func (S *Server) Execution(tgt, fun, arg string) {
	values := url.Values{}
	values.Add("client", "local")
	values.Add("fun", fun)
	values.Add("arg", arg)
	values.Add("tgt", "*")
	//fmt.Println(strings.NewReader(values.Encode()))
	//Encode方法将v编码为url编码格式("bar=baz&foo=quux")，编码时会以键进行排序。
	req, err := http.NewRequest("POST", serverUrl, strings.NewReader(values.Encode()))
	if err != nil {
		fmt.Println("new...", err)
		return
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-Auth-Token", S.Token)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded") //没有出现：406 Not Acceptable <nil>
	res, err := S.client.Do(req)
	if err != nil {
		fmt.Println("do ...", err)
		return
	}
	if res.StatusCode != 200 {
		fmt.Println(res.Status, err)
		return
	}
	defer res.Body.Close()
	context, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("read:  ", err)
	}
	fmt.Printf("%s\n", context)
}

func main() {
	Client()
	// ret, _ := obj.ListALLKey()
	// fmt.Println(ret)

	tgt := "data-2"
	fun := "cmd.run"
	arg := "df -h"
	S.Execution(tgt, fun, arg)
	//fmt.Println("go")
	// Login()
	// fmt.Println(S)
}
