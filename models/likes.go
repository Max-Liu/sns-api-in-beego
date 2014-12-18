package models

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Likes struct {
	Id        int64     `orm:"column(id);pk"`
	CreatedAt time.Time `orm:"column(created_at);type(timestamp);null"`
	UpdatedAt time.Time `orm:"column(updated_at);type(timestamp);null"`
	Photo     *Photos   `orm:"column(target_id);rel(fk)" valid:"Required"`
	User      *Users    `orm:"column(user_id);rel(fk)"`
}

type LikesApi struct {
	CreatedAt int64
	Photo     *PhotosApi
}
type LikesUsersApi struct {
	CreatedAt int64
	UserName  string
	UserImage string
	UserId    int64
}

func ConverToLikedPhotoApiStruct(m *Likes) (data *LikesApi) {
	data = new(LikesApi)
	data.CreatedAt = m.CreatedAt.Unix()
	data.Photo = ConverToPhotoApiStruct(m.Photo)
	return data
}

func init() {
	orm.RegisterModel(new(Likes))
}

// AddLikes insert a new Likes into database and returns
// last inserted Id on success.
func AddLikes(m *Likes) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

func GetLikesCount(userId int64) (count int64) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Likes))
	count, _ = qs.Filter("user_id", userId).Count()
	return count
}

// GetLikesById retrieves Likes by Id. Returns error if
// Id doesn't exist
func GetLikesById(id int64) (v *Likes, err error) {
	o := orm.NewOrm()
	v = &Likes{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllLikes retrieves all Likes matches certain condition. Returns empty list if
// no records exist
func GetAllLikes(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []orm.Params, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Likes))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []orm.Params
	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(limit, offset).Values(&l, fields...); err == nil {
		return l, nil
	}
	return nil, err
}

// UpdateLikes updates Likes by Id and returns error if
// the record to be updated doesn't exist
func UpdateLikesById(m *Likes) (err error) {
	o := orm.NewOrm()
	v := Likes{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteLikes deletes Likes by Id and returns error if
// the record to be deleted doesn't exist
func DeleteLikes(id int64) (err error) {
	o := orm.NewOrm()
	v := Likes{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Likes{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func DeleteLikedPhoto(userId, photoId int64) (num int64, err error) {
	o := orm.NewOrm()
	target := Photos{Id: photoId}
	err = o.Read(&target)
	if err != nil {
		return 0, err
	}
	user := Users{Id: userId}
	err = o.Read(&user)
	if err != nil {
		return 0, err
	}
	v := Likes{Photo: &target, User: &user}

	if err = o.Read(&v, "target_id", "user_id"); err == nil {
		num, err = o.Delete(&v)
	}
	return num, err
}

func HasLikedPhoto(photoId, userId int64) (hasLiked bool) {
	o := orm.NewOrm()
	exist := o.QueryTable(new(Likes)).Filter("target_id", photoId).Filter("user_id", userId).Exist()

	return exist
}

func GetUsersByLikesPhoto(photoId int64) (usersList []*LikesUsersApi, err error) {

	var lists []orm.Params
	o := orm.NewOrm()
	o.QueryTable(new(Likes)).Filter("target_id", photoId).Values(&lists, "User", "CreatedAt")
	for _, v := range lists {
		userId := v["User__User"]
		CreatedAt := v["CreatedAt"].(time.Time)
		user, _ := GetUsersById(userId.(int64))
		data := new(LikesUsersApi)

		data.CreatedAt = CreatedAt.Unix()
		data.UserImage = user.Head
		data.UserName = user.Name
		data.UserId = userId.(int64)
		usersList = append(usersList, data)
	}
	return usersList, err
}
