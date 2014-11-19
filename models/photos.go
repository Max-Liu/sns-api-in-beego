package models

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/beego/redigo/redis"
)

type Photos struct {
	Id        int64     `orm:"column(id);pk"`
	Title     string    `orm:"column(title);size(45);null" form:"title" valid:"Required"`
	Path      string    `orm:"column(path);null"`
	CreatedAt time.Time `orm:"column(created_at);type(timestamp);null"`
	UpdatedAt time.Time `orm:"column(updated_at);type(timestamp);null"`
	User      *Users    `orm:"column(user_id);rel(fk)"`
	Likes     int64     `orm:"column(likes);null"`
}
type PhotosApi struct {
	Id        int64
	Title     string
	Path      string
	CreatedAt int64
	//UpdatedAt int64
	User  *UsersApi
	Likes int64
}

func init() {
	orm.RegisterModel(new(Photos))
}

func ConverToPhotoApiStruct(m *Photos) (data *PhotosApi) {
	data = new(PhotosApi)
	data.Id = m.Id
	data.Title = m.Title
	data.CreatedAt = m.CreatedAt.Unix()
	//data.UpdatedAt = m.UpdatedAt.Unix()
	data.User = ConverToUserApiStruct(m.User)
	data.Likes = m.Likes
	data.Path = m.Path

	return data
}

// AddPhotos insert a new Photos into database and returns
// last inserted Id on success.
func AddPhotos(m *Photos) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPhotosById retrieves Photos by Id. Returns error if
// Id doesn't exist
func GetPhotosById(id int64, fields ...string) (v *Photos, err error) {
	o := orm.NewOrm()
	v = &Photos{Id: id}
	if err = o.Read(v, fields...); err == nil {
		if v.User != nil {
			o.Read(v.User)
		}
		return v, nil
	}
	return nil, err
}

// GetAllPhotos retrieves all Photos matches certain condition. Returns empty list if
// no records exist
func GetAllPhotos(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []orm.Params, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Photos))
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

	//var l []Photos
	var l []orm.Params
	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(limit, offset).Values(&l, fields...); err == nil {
		return l, nil
	}
	return nil, err
}

// UpdatePhotos updates Photos by Id and returns error if
// the record to be updated doesn't exist
func UpdatePhotosById(m *Photos) (err error) {
	o := orm.NewOrm()
	v := Photos{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeletePhotos deletes Photos by Id and returns error if
// the record to be deleted doesn't exist
func DeletePhotos(id int64) (err error) {
	o := orm.NewOrm()
	v := Photos{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Photos{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
func GetMyPhotos(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (l []orm.Params, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Photos))
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

	if _, err := qs.Limit(limit, offset).Values(&l, fields...); err == nil {
		// trim unused fields
		for _, v := range l {
			idStr := strconv.Itoa(int(v["Id"].(int64)))

			query = make(map[string]string)
			query["photo_id"] = idStr
			fields = []string{"user__name", "content", "CreatedAt"}
			sortby = []string{"id"}
			comments, _ := GetAllPhotoComments(query, fields, sortby, order, offset, limit)
			v["comment"] = comments
			if v["comment"].([]orm.Params) == nil {
				v["comment"] = ""
			}
		}
		return l, nil
	}
	return nil, err
}
func GetFollowingPhotos(userId int64, offset int64, limit int64) (interface{}, error) {
	redisAddress, _ := beego.Config("String", "redisServer", "")
	c, err := redis.Dial("tcp", redisAddress.(string))
	defer c.Close()
	if err != nil {
		beego.Error(err.Error())
	}
	userIdStr := strconv.FormatInt(userId, 10)
	result, err := c.Do("LRANGE", "ptm:"+userIdStr, offset, offset+limit)

	if err != nil {
		beego.Error(err.Error())
	}

	if reflect.TypeOf(result).String() == "[]interface {}" {
		if reflect.ValueOf(result).Len() == 0 {
			return result, nil
		}
	}

	var photoIdList []string
	for _, photoId := range result.([]interface{}) {
		photoIdList = append(photoIdList, string(photoId.([]uint8)))
	}
	o := orm.NewOrm()
	qs := o.QueryTable("photos")
	var lists []orm.Params
	qs.Filter("id__in", photoIdList).Values(&lists)
	var photoApiDatas []*PhotosApi
	var photo Photos
	for _, v := range lists {
		photo.CreatedAt = v["CreatedAt"].(time.Time)
		photo.Id = v["Id"].(int64)
		photo.Likes = v["Likes"].(int64)
		photo.Path = v["Path"].(string)
		photo.Title = v["Title"].(string)
		photo.User, _ = GetUsersById(v["User"].(int64))
		photoApiData := ConverToPhotoApiStruct(&photo)
		photoApiDatas = append(photoApiDatas, photoApiData)

	}
	return photoApiDatas, err
}
