package test

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	_ "pet/routers"
	"pet/utils"
	"testing"

	_ "github.com/astaxie/beego/session/mysql"
	_ "github.com/astaxie/beego/session/redis"
)

var dbAddress string = "root:38143195@tcp(192.168.33.11:3306)/pet?charset=utf8"
var client http.Client

var testUrl string = "http://localhost:8080"
var jsonData helper.Response
var request *http.Request

func init() {
	request, _ = http.NewRequest("GET", testUrl+"/v1/users/login?info=forevervmax@gmail.com&password=123", nil)
	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	jar := new(testJar)
	jar.jar = make(map[string][]*http.Cookie)
	jar.SetCookies(request.URL, resp.Cookies())
	client.Jar = jar
}

type testJar struct {
	jar map[string][]*http.Cookie
}

func (j *testJar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	j.jar[u.Host] = cookies
}
func (j *testJar) Cookies(u *url.URL) []*http.Cookie {
	return j.jar[u.Host]
}

func Login(t *testing.T) {
	request, _ = http.NewRequest("GET", testUrl+"/v1/users/login?info=forevervmax@gmail.com&password=123", nil)
	resp, err := client.Do(request)
	if err != nil {
		t.Error(err.Error())
	}
	if resp.StatusCode != 200 {
		t.Error("登陆失败")
		t.FailNow()
	}
	t.Log("status 200")

	bodyString, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err.Error())
	}
	err = json.Unmarshal(bodyString, &jsonData)
	if err != nil {
		t.Error(err.Error())
	}
	if jsonData.Err != 0 {
		t.Error(jsonData.Msg)
		t.FailNow()
	}

	jar := new(testJar)
	jar.jar = make(map[string][]*http.Cookie)
	jar.SetCookies(request.URL, resp.Cookies())

	client.Jar = jar
	t.Log("登陆成功")
}

func TestGetUsersId(t *testing.T) {
	request, _ = http.NewRequest("GET", testUrl+"/v1/users/29", nil)
	baseTest(t)
}

func TestUpdateUser(t *testing.T) {
	query := make(map[string]string)
	query["id"] = "29"
	query["gender"] = "1"
	query["birthday"] = "1989-09-13"
	request = helper.MakeUploadRequest(query, testUrl+"/v1/users/29", "PUT", "head", "head.jpg")
	baseTest(t)
}
