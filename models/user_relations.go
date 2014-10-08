package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type UserRelations struct {
	Id        int       `orm:"column(id);pk"`
	Following *Users    `orm:"column(following);rel(fk)" form:"following" valid:"Required"`
	Follower  *Users    `orm:"column(follower);size(45);rel(fk)" form:"follower" valid:"Required"`
	CreatedAt time.Time `orm:"column(created_at);type(timestamp);null"`
	UpdatedAt time.Time `orm:"column(updated_at);type(timestamp);null"`
}

func init() {
	orm.RegisterModel(new(UserRelations))
}

// AddUserRelations insert a new UserRelations into database and returns
// last inserted Id on success.
func AddUserRelations(m *UserRelations) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetUserRelationsById retrieves UserRelations by Id. Returns error if
// Id doesn't exist
func GetUserRelationsById(id int) (v *UserRelations, err error) {
	o := orm.NewOrm()
	v = &UserRelations{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllUserRelations retrieves all UserRelations matches certain condition. Returns empty list if
// no records exist
func GetAllUserRelations(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(UserRelations))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v).RelatedSel()
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

	var l []UserRelations
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

// UpdateUserRelations updates UserRelations by Id and returns error if
// the record to be updated doesn't exist
func UpdateUserRelationsById(m *UserRelations) (err error) {
	o := orm.NewOrm()
	v := UserRelations{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteUserRelations deletes UserRelations by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUserRelations(id int) (err error) {
	o := orm.NewOrm()
	v := UserRelations{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&UserRelations{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
func DeleteUserRelationsByUsers(follower, following int) (num int64, err error) {

	o := orm.NewOrm()
	followerUser := Users{Id: follower}
	err = o.Read(&followerUser)
	if err != nil {
		return 0, err
	}

	followingUser := Users{Id: following}
	err = o.Read(&followingUser)
	if err != nil {
		return 0, err
	}
	v := UserRelations{Follower: &followerUser, Following: &followingUser}
	// ascertain id exists in the database
	if err = o.Read(&v, "follower", "following"); err == nil {
		num, err = o.Delete(&v)
	}
	return num, err
}
