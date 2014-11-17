package seed

import (
	"math/rand"
	"pet/models"
	helper "pet/utils"
	"strconv"
	"time"
	"github.com/astaxie/beego/orm"
	"github.com/manveru/faker"
)

import _ "github.com/go-sql-driver/mysql"
import _ "github.com/astaxie/beego/session/mysql"

var dbAddress string = "root:38143195@tcp(192.168.33.11:3306)/pet?charset=utf8"
var host = "http://localhost:8080"

func init() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", dbAddress)
}

func GenerateUser() *models.Users {
	fake, _ := faker.New("en")
	fakeUser := new(models.Users)
	fakeUser.Birthday = "1989-09-13"
	fakeUser.Email = fake.Email()
	fakeUser.CreatedAt = time.Now()
	fakeUser.Gender = randInt(0, 1)
	fakeUser.Name = fake.Name()
	phoneInt := randInt(18600000000, 18619999999)
	phoneStr := strconv.Itoa(phoneInt)
	fakeUser.Phone = phoneStr
	fakeUser.Password = strconv.Itoa(randInt(18600000000, 18619999999))
	fakeUser.UpdatedAt = time.Now()
	return fakeUser
}

func UploadPhotos() {
	o := orm.NewOrm()
	user := new(models.Users)
	var lists []orm.Params
	o.QueryTable(user).OrderBy("-id").Limit(10).Values(&lists)
	for _, user := range lists {
		clientUser := new(helper.User)
		clientUser.Info = user["Email"].(string)
		clientUser.Pwd = user["Password"].(string)
		clientUser.Login()
		query := make(map[string]string)
		randStr := strconv.Itoa(time.Now().Nanosecond())
		query["title"] = "mydog" + randStr
		clientUser.Request = helper.MakeUploadRequest(query, host+"/v1/photos/", "POST", "photo", "dog.jpg")
		clientUser.DoRequest()
	}
}

func MakeFakeUserData() *helper.User {
	fakeUser := GenerateUser()
	query := make(map[string]string)
	query["email"] = fakeUser.Email + strconv.FormatInt(time.Now().UnixNano(), 10)
	query["phone"] = fakeUser.Phone
	query["name"] = fakeUser.Name
	query["password"] = fakeUser.Password

	user := new(helper.User)
	user.Request = helper.MakeRequest(query, host+"/v1/users/register", "POST")
	user.DoRequest()
	return user
}

func GenerateUserRelation() {
	var clientList []*helper.User

	for i := 0; i < 10; i++ {
		clientList = append(clientList, MakeFakeUserData())
	}
	for _, clientA := range clientList {
		for _, clientB := range clientList {
			tempResp := clientB.Resp
			userA := clientA.Resp.Data.(map[string]interface{})
			userB := clientB.Resp.Data.(map[string]interface{})
			if userA["Id"].(float64) == userB["Id"].(float64) {
				continue
			}

			clientB.Info = userB["Email"].(string)
			clientB.Pwd = userB["Pwd"].(string)
			clientB.Login()

			query := make(map[string]string)
			userAStrId := strconv.Itoa(int(userA["Id"].(float64)))
			query["following"] = userAStrId

			clientB.Request = helper.MakeRequest(query, host+"/v1/ul", "POST")
			clientB.DoRequest()
			clientB.Resp = tempResp
		}
	}
}

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}
