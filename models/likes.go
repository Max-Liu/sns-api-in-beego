package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Likes struct {
	Id        int       `orm:"column(id);pk"`
	CreatedAt time.Time `orm:"column(created_at);type(timestamp);null"`
	UpdatedAt time.Time `orm:"column(updated_at);type(timestamp);null"`
	TargetId  *Photos   `orm:"column(target_id);rel(fk)" valid:"Required"`
	UserId    *Users    `orm:"column(user_id);rel(fk)"`
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

// GetLikesById retrieves Likes by Id. Returns error if
// Id doesn't exist
func GetLikesById(id int) (v *Likes, err error) {
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
	offset int64, limit int64) (ml []interface{}, err error) {
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

	var l []Likes
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
func DeleteLikes(id int) (err error) {
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

func DeleteLikedPhoto(userId, photoId int) (num int64, err error) {
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
	v := Likes{TargetId: &target, UserId: &user}

	if err = o.Read(&v, "target_id", "user_id"); err == nil {
		num, err = o.Delete(&v)
	}
	return num, err
}
