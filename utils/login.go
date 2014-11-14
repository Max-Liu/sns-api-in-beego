package helper

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"testing"
)

type User struct {
	Info string
	Pwd  string
	http.Client
	Request  *http.Request
	Response *http.Response
	Resp     DataResponse
}

func NewLogin(loginUrl string) (user *User) {

	user = new(User)
	request, _ := http.NewRequest("GET", loginUrl, nil)
	resp, err := user.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	jar := new(myJar)
	jar.jar = make(map[string][]*http.Cookie)
	jar.SetCookies(request.URL, resp.Cookies())

	user.Jar = jar
	return user
}

func (client *User) Login() {
	jar := new(myJar)
	jar.jar = make(map[string][]*http.Cookie)
	client.Jar = jar
	client.Request, _ = http.NewRequest("GET", "http://localhost:8080"+"/v1/users/login?info="+client.Info+"&password="+client.Pwd, nil)
	client.DoRequest()
	jar.SetCookies(client.Request.URL, client.Response.Cookies())

	client.Jar = jar
}

func (client *User) DoRequest() {

	var err error
	client.Response, err = client.Do(client.Request)
	if err != nil {
		log.Fatal(err.Error())
	}
	bodyByte, err := ioutil.ReadAll(client.Response.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	if client.Response.StatusCode != 200 {
		log.Println("Printing error html file in current Path:output.html")
		ioutil.WriteFile("clientErr.html", bodyByte, 0644)
		log.Fatalln("resp code is not 200", client.Response.StatusCode)
	}
	var jsonData DataResponse
	err = json.Unmarshal(bodyByte, &jsonData)
	if err != nil {
		log.Fatal(err.Error())

	}
	if jsonData.Err != 0 {
		log.Println(jsonData.Msg)
	}
	client.Resp = jsonData
}

func (client *User) BaseTest(t *testing.T) {
	resp, err := client.Do(client.Request)
	if err != nil {
		t.Error(err.Error())
	}
	bodyByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err.Error())
	}
	if resp.StatusCode != 200 {
		//spew.Dump(body.Bytes())
		log.Println("Printing error html file in current Path:output.html")
		ioutil.WriteFile("output.html", bodyByte, 0644)
		t.Error(resp.StatusCode)
		t.FailNow()
	} else {
		t.Log("status 200")
	}

	var jsonData DataResponse
	err = json.Unmarshal(bodyByte, &jsonData)
	if err != nil {
		t.Error(err.Error())
	}
	if jsonData.Err != 0 {
		t.Error(jsonData.Msg)
	}
	client.Resp = jsonData
}

type myJar struct {
	jar map[string][]*http.Cookie
}

func (j *myJar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	j.jar[u.Host] = cookies
}
func (j *myJar) Cookies(u *url.URL) []*http.Cookie {
	return j.jar[u.Host]
}
