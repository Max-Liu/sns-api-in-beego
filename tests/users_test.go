package test

import (
	"net/http"
	_ "pet/routers"
	"pet/seed"
	"pet/utils"
	"testing"

	_ "github.com/astaxie/beego/session/mysql"
	_ "github.com/astaxie/beego/session/redis"
)

var client *helper.User

var testUrl string = "http://localhost:8080"
var jsonData helper.Response
var request *http.Request

func init() {
	client = helper.NewLogin(testUrl + "/v1/users/login?info=forevervmax@gmail.com&password=123")
}

func TestGetUsersId(t *testing.T) {
	request, _ = http.NewRequest("GET", testUrl+"/v1/users/29", nil)
	baseTest(t)
}

func TestRegister(t *testing.T) {
	user := seed.GenerateUser()

	query := make(map[string]string)
	query["email"] = user.Email
	query["phone"] = user.Phone
	query["name"] = user.Name
	query["password"] = user.Password
	request = helper.MakeRequest(query, testUrl+"/v1/users/register", "POST")
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
