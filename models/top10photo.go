package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Top10photo struct {
	Id    int64   `orm:"auto"`
	Photo *Photos `orm:"column(photo_id);rel(fk)" valid:"Required"`
}

func init() {
	orm.RegisterModel(new(Top10photo))
}
func ConverToTop10ApiStruct(m *Photos) (data *PhotosApi) {
	data = ConverToPhotoApiStruct(m)
	return data
}

func GetTop10() (data []*PhotosApi, err error) {
	var lists []orm.Params
	var photoList []*PhotosApi
	o := orm.NewOrm()
	_, err = o.QueryTable(new(Top10photo)).Values(&lists)

	for _, v := range lists {
		photo, _ := GetPhotosById(v["Photo"].(int64))
		photoApi := ConverToPhotoApiStruct(photo)
		photoList = append(photoList, photoApi)
	}
	return photoList, err
}

// AddTop10photo insert a new Top10photo into database and returns
// last inserted Id on success.
func AddTop10photo(m *Top10photo) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTop10photoById retrieves Top10photo by Id. Returns error if
// Id doesn't exist
func GetTop10photoById(id int64) (v *Top10photo, err error) {
	o := orm.NewOrm()
	v = &Top10photo{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTop10photo retrieves all Top10photo matches certain condition. Returns empty list if
// no records exist
func GetAllTop10photo(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Top10photo))
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

	var l []Top10photo
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

// UpdateTop10photo updates Top10photo by Id and returns error if
// the record to be updated doesn't exist
func UpdateTop10photoById(m *Top10photo) (err error) {
	o := orm.NewOrm()
	v := Top10photo{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTop10photo deletes Top10photo by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTop10photo(id int64) (err error) {
	o := orm.NewOrm()
	v := Top10photo{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Top10photo{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
