package models

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Photos struct {
	Id        int       `orm:"column(id);pk"`
	Title     string    `orm:"column(title);size(45);null" form:"title" valid:"Required"`
	Path      string    `orm:"column(path);null"`
	CreatedAt time.Time `orm:"column(created_at);type(timestamp);null"`
	UpdatedAt time.Time `orm:"column(updated_at);type(timestamp);null"`
	User      *Users    `orm:"column(user_id);rel(fk)"`
	Likes     int       `orm:"column(likes);null"`
}

func init() {
	orm.RegisterModel(new(Photos))
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
func GetPhotosById(id int, fields ...string) (v *Photos, err error) {
	o := orm.NewOrm()
	v = &Photos{Id: id}
	if err = o.Read(v, fields...); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllPhotos retrieves all Photos matches certain condition. Returns empty list if
// no records exist
func GetAllPhotos(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
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

	var l []Photos
	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
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
func DeletePhotos(id int) (err error) {
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
