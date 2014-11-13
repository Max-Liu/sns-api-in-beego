package helper

import (
	"log"
	"net/http"
	"net/url"
)

type User struct {
	http.Client
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

type myJar struct {
	jar map[string][]*http.Cookie
}

func (j *myJar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	j.jar[u.Host] = cookies
}
func (j *myJar) Cookies(u *url.URL) []*http.Cookie {
	return j.jar[u.Host]
}
