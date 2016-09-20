package main

import (
	//"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	//"log"
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
	Client   *http.Client
}

const (
	username  = "saltapi"
	password  = "saltapi"
	eauth     = "pam"
	serverUrl = "https://192.168.30.129:8888/"
	//serverUrl = "https://139.196.231.187:8888/"
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

//var S *Server = new(Server)

// ssl  忽略认证
func (S *Server) GetToken() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	S.Client = &http.Client{Transport: tr}
	values := url.Values{}
	values.Add("username", username)
	values.Add("password", password)
	values.Add("eauth", eauth)
	result := S.RequestPost("POST", serverUrl+"/login", values)
	json.Unmarshal([]byte(result), &obj)
	S.Token = obj.Return[0].Token
}

func (S *Server) ListALLKey() (string, error) {
	req, err := http.NewRequest("GET", serverUrl+"/keys", nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("X-Auth-Token", S.Token)
	res, err := S.Client.Do(req)
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

// cmd.run 需要arg  ,test.ping 则不需要
func (S *Server) Execution(tgt, fun, arg string) string {
	values := url.Values{}
	values.Add("client", "local")
	values.Add("fun", fun)
	if arg != "" {
		values.Add("arg", arg)
	}
	values.Add("tgt", tgt)
	//fmt.Println(strings.NewReader(values.Encode()))
	//Encode方法将v编码为url编码格式("bar=baz&foo=quux")，编码时会以键进行排序。
	return S.RequestPost("POST", serverUrl, values)
}

// 配置模板
func (S *Server) DeployMouldel(tgt, arg string) string {
	values := url.Values{}
	values.Add("client", "local")
	values.Add("fun", "state.sls")
	values.Add("tgt", tgt)
	values.Add("arg", arg)
	return S.RequestPost("POST", serverUrl, values)
}

func (S *Server) RequestPost(method, serverUrl string, values url.Values) string {
	//fmt.Println(strings.NewReader(values.Encode()))
	//Encode方法将v编码为url编码格式("bar=baz&foo=quux")，编码时会以键进行排序。
	req, err := http.NewRequest(method, serverUrl, strings.NewReader(values.Encode()))
	if err != nil {
		fmt.Println("new...", err)
		return ""
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-Auth-Token", S.Token)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded") //没有出现：406 Not Acceptable <nil>
	res, err := S.Client.Do(req)
	if err != nil {
		fmt.Println("do ...", err)
		return ""
	}
	if res.StatusCode != 200 {
		fmt.Println(res.Status, err)
		return ""
	}
	defer res.Body.Close()
	context, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("read:  ", err)
		return ""
	}
	return string(context)
}

func (S *Server) AsyncDeployMouldel(tgt, arg string) string {
	values := url.Values{}
	values.Add("client", "local_async")
	values.Add("fun", "state.sls")
	values.Add("tgt", tgt)
	values.Add("arg", arg)
	return S.RequestPost("POST", serverUrl, values)
}

//通过jid 获取内容,没有jid 则获取所有jid
func (S *Server) GetAsyncContent(job_id string) string {
	if job_id == "" {
		return S.RequestPost("GET", serverUrl+"/jobs/", nil)
	}
	return S.RequestPost("GET", serverUrl+"/jobs/"+job_id, nil)
}

func (S *Server) Events() string {
	return S.RequestPost("GET", serverUrl+"/events/", nil)
}

//获取minion 的基本信息
func (S *Server) Minions(minion string) string {
	if minion != "" && minion != "*" {
		return S.RequestPost("GET", serverUrl+"/minions/"+minion, nil)
	}
	return S.RequestPost("GET", serverUrl+"/minions/", nil)
}

//管理key  way 方式
func (S *Server) ManageKey(way, id string) string {
	values := url.Values{}
	values.Add("client", "wheel")
	if way == "accept" {
		values.Add("fun", "key.accept")
	} else {
		values.Add("fun", "key.delete")
	}
	values.Add("match", id)
	return S.RequestPost("POST", serverUrl, values)
}

func main() {
	var S *Server = new(Server)
	S.GetToken()
	// ret, _ := obj.ListALLKey()
	// fmt.Println(ret)

	//tgt := "data-2"
	//fun := "test.ping"
	//arg := "nginx"
	// fmt.Println(S.Execution(tgt, fun, arg))
	//fmt.Println("go")
	// Login()
	// fmt.Println(S)
	//S.DeployMouldel(tgt, arg)
	// values := url.Values{}
	// values.Add("client", "local")
	// values.Add("fun", "state.sls")
	// values.Add("tgt", tgt)
	// values.Add("arg", arg)
	// RequestPost("POST", serverUrl, values)
	//fmt.Println(S.AsyncDeployMouldel(tgt, arg))
	//fmt.Println(S.GetAsyncContent(""))
	//fmt.Println(S.Minions(""))
	//ret, _ := S.ListALLKey()
	//fmt.Println(ret)
	// S.AcceptKey("accept", "data-2")
	ret, _ := S.ListALLKey()
	fmt.Println(ret)

}
