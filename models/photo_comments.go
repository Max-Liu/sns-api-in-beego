package models

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type PhotoComments struct {
	Id        int       `orm:"column(id);pk"`
	Photo     *Photos   `orm:"column(photo_id);rel(fk)"`
	Content   string    `orm:"column(content);null" form:"content" valid:"Required"`
	CreatedAt time.Time `orm:"column(created_at);type(timestamp);null"`
	UpdatedAt time.Time `orm:"column(updated_at);type(timestamp);null"`
	User      *Users    `orm:"column(user_id);null;rel(fk)"`
}
type PhotoCommentsApi struct {
	Content   string
	CreatedAt int64
	UserName  string
}

func ConverToCommentsApirStruct(m *PhotoComments) (data *PhotoCommentsApi) {
	data = new(PhotoCommentsApi)
	data.Content = m.Content
	data.CreatedAt = m.CreatedAt.Unix()
	data.UserName = m.User.Name
	return data
}

func init() {
	orm.RegisterModel(new(PhotoComments))
}

// AddPhotoComments insert a new PhotoComments into database and returns
// last inserted Id on success.
func AddPhotoComments(m *PhotoComments) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPhotoCommentsById retrieves PhotoComments by Id. Returns error if
// Id doesn't exist
func GetPhotoCommentsById(id int) (v *PhotoComments, err error) {
	o := orm.NewOrm()
	v = &PhotoComments{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllPhotoComments retrieves all PhotoComments matches certain condition. Returns empty list if
// no records exist
func GetAllPhotoComments(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []orm.Params, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(PhotoComments))
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

// UpdatePhotoComments updates PhotoComments by Id and returns error if
// the record to be updated doesn't exist
func UpdatePhotoCommentsById(m *PhotoComments) (err error) {
	o := orm.NewOrm()
	v := PhotoComments{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeletePhotoComments deletes PhotoComments by Id and returns error if
// the record to be deleted doesn't exist
func DeletePhotoComments(id int) (err error) {
	o := orm.NewOrm()
	v := PhotoComments{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&PhotoComments{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
