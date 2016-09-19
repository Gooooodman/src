package salt

import (
	"crypto/tls"
	"encoding/json"
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

var obj Obj

// ssl  忽略认证  保存token
func (S *Server) GetToken() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	S.Client = &http.Client{Transport: tr}
	values := url.Values{}
	values.Add("username", username)
	values.Add("password", password)
	values.Add("eauth", eauth)
	result, _ := S.RequestPost("POST", serverUrl+"/login", values)
	json.Unmarshal([]byte(result), &obj)
	S.Token = obj.Return[0].Token
}

//列出所有key
func (S *Server) ListALLKey() (ret string) {
	ret, _ = S.RequestPost("GET", serverUrl+"/keys", nil)
	return
}

// cmd.run 需要arg  ,test.ping 则不需要
func (S *Server) Execution(tgt, fun, arg string) (ret string) {
	values := url.Values{}
	values.Add("client", "local")
	values.Add("fun", fun)
	if arg != "" {
		values.Add("arg", arg)
	}
	values.Add("tgt", tgt)
	//fmt.Println(strings.NewReader(values.Encode()))
	//Encode方法将v编码为url编码格式("bar=baz&foo=quux")，编码时会以键进行排序。
	ret, _ = S.RequestPost("POST", serverUrl, values)
	return
}

// 配置模板
func (S *Server) DeployMouldel(tgt, arg string) (ret string) {
	values := url.Values{}
	values.Add("client", "local")
	values.Add("fun", "state.sls")
	values.Add("tgt", tgt)
	values.Add("arg", arg)
	ret, _ = S.RequestPost("POST", serverUrl, values)
	return

}

//发送请求返回结果
func (S *Server) RequestPost(method, serverUrl string, values url.Values) (string, error) {
	//fmt.Println(strings.NewReader(values.Encode()))
	//Encode方法将v编码为url编码格式("bar=baz&foo=quux")，编码时会以键进行排序。
	req, err := http.NewRequest(method, serverUrl, strings.NewReader(values.Encode()))
	if err != nil {
		log.Println("Request.. ", err)
		return "", err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-Auth-Token", S.Token)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded") //没有出现：406 Not Acceptable <nil>
	res, err := S.Client.Do(req)
	defer res.Body.Close()
	if err != nil {
		log.Println("Do.. ", err)
		return "", err
	}
	if res.StatusCode != 200 {
		log.Println("Status.. ", res.StatusCode)
		return "", err
	}
	context, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("ReadALL.. ", err)
		return "", err
	}
	return string(context), nil
}

//异步处理 返回 jid 与 id
func (S *Server) AsyncDeployMouldel(tgt, arg string) (ret string) {
	values := url.Values{}
	values.Add("client", "local_async")
	values.Add("fun", "state.sls")
	values.Add("tgt", tgt)
	values.Add("arg", arg)
	ret, _ = S.RequestPost("POST", serverUrl, values)
	return
}

//通过jid 获取内容,没有jid 则获取所有jid
func (S *Server) GetAsyncContent(job_id string) (ret string) {
	// if job_id == "" {
	// 	return S.RequestPost("GET", serverUrl+"/jobs/", nil)
	// }
	ret, _ = S.RequestPost("GET", serverUrl+"/jobs/"+job_id, nil)
	return
}

//时间太长不建议使用
func (S *Server) Events() (ret string) {
	ret, _ = S.RequestPost("GET", serverUrl+"/events/", nil)
	return

}

//获取minion 的基本信息
func (S *Server) Minions(minion string) (ret string) {
	// if minion != "" && minion != "*" {
	ret, _ = S.RequestPost("GET", serverUrl+"/minions/"+minion, nil)
	return
	// }
	// ret, _ =  S.RequestPost("GET", serverUrl+"/minions/", nil)
}

//管理key  way 方式
func (S *Server) ManageKey(way, id string) (ret string) {
	values := url.Values{}
	values.Add("client", "wheel")
	if way == "accept" {
		values.Add("fun", "key.accept")
	} else {
		values.Add("fun", "key.delete")
	}
	values.Add("match", id)
	ret, _ = S.RequestPost("POST", serverUrl, values)
	return
}
