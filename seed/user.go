package seed

import (
	"math/rand"
	"pet/models"
	"strconv"
	"time"
	"github.com/astaxie/beego/orm"
	"github.com/manveru/faker"
)

import _ "github.com/go-sql-driver/mysql"
import _ "github.com/astaxie/beego/session/mysql"

var dbAddress string = "root:38143195@tcp(192.168.33.11:3306)/pet?charset=utf8"

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
	fakeUser.Password = fake.Name()
	fakeUser.UpdatedAt = time.Now()
	return fakeUser
}
func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}
