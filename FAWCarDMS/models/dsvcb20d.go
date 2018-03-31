package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Dsvcb20d struct {
	Id                       int       `orm:"column(ID);pk" description:"ID"`
	NMAINID                  string    `orm:"column(NMAINID);size(20)" description:"委托书ID"`
	CITEMCODE                string    `orm:"column(CITEMCODE);size(8)" description:"修理项目的代码"`
	CITEMNAME                string    `orm:"column(CITEMNAME);size(30)" description:"修理项目的名称"`
	CKIND                    string    `orm:"column(CKIND);size(1);null" description:"'修理项目的性质决定该修理项目是否是收费的项目。性质 0 正常,1 首保,2 索赔,3 免费,4 另项';"`
	NMANHAUREXPEN            float64   `orm:"column(NMANHAUREXPEN);null" description:"'修理项目的工时费等于工时乘以工时单价（使用标准工时时）';"`
	NMANHAUR                 float64   `orm:"column(NMANHAUR);null" description:"'修理项目的工需要工作时间数';"`
	NMANHAURDISCOUNTRATE     float64   `orm:"column(NMANHAURDISCOUNTRATE);null" description:"'单个项目的工时优惠率';"`
	NWORKGROUPID             float64   `orm:"column(NWORKGROUPID);null" description:"维修车辆的班组"`
	NPRIMARYREPAIRMANID      int       `orm:"column(NPRIMARYREPAIRMANID);null" description:"修车人"`
	DTASKBEGINDATE           time.Time `orm:"column(DTASKBEGINDATE);type(date);null" description:"分配修理项目的时间"`
	DBEGINDATE               time.Time `orm:"column(DBEGINDATE);type(date);null" description:"修车开始的时间"`
	DTASKFINISHDATE          time.Time `orm:"column(DTASKFINISHDATE);type(date);null" description:"完成该修理项目的时间"`
	NEFFECTIVEMANHAUR        float64   `orm:"column(NEFFECTIVEMANHAUR);null" description:"实际有效的工作时间，应等于工时'"`
	CPURSUEITEMFLAG          string    `orm:"column(CPURSUEITEMFLAG);size(1);null" description:"Lzjxm=1 新项目 追加项目"`
	CREPAIRAGAINFLAG         string    `orm:"column(CREPAIRAGAINFLAG);size(1);null" description:"lfg=1返工的项目 lfg=1（返工）的项目不允许修改'"`
	CACROSSCURRENTPERIODFLAG string    `orm:"column(CACROSSCURRENTPERIODFLAG);size(1);null" description:"跨期标识"`
	CBEGINFLAG               string    `orm:"column(CBEGINFLAG);size(1);null" description:"开始标识"`
	CFINISHFLAG              string    `orm:"column(CFINISHFLAG);size(1);null" description:"完工标识"`
	CREMARK                  string    `orm:"column(CREMARK);size(50);null" description:"备注"`
	DTSTAMP                  time.Time `orm:"column(DTSTAMP);type(timestamp)"`
	CKINDBACK                string    `orm:"column(CKINDBACK);size(1);null" description:"原维修类别数据备份"`
	CRETURNFLAG              string    `orm:"column(CRETURNFLAG);size(1);null" description:"预警标识"`
	CRETURNFLAG1             string    `orm:"column(CRETURNFLAG1);size(1);null" description:"预警标识是否允许修改0允许1不允许"`
	CRETURNFLAG2             string    `orm:"column(CRETURNFLAG2);size(1);null" description:"返修标识"`
	CREMINDFLAG              string    `orm:"column(CREMINDFLAG);size(1);null" description:"是否提醒标识"`
}

func (t *Dsvcb20d) TableName() string {
	return "dsvcb20d"
}

func init() {
	orm.RegisterModel(new(Dsvcb20d))
}

// AddDsvcb20d insert a new Dsvcb20d into database and returns
// last inserted Id on success.
func AddDsvcb20d(m *Dsvcb20d) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetDsvcb20dById retrieves Dsvcb20d by Id. Returns error if
// Id doesn't exist
func GetDsvcb20dById(id int) (v *Dsvcb20d, err error) {
	o := orm.NewOrm()
	v = &Dsvcb20d{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllDsvcb20d retrieves all Dsvcb20d matches certain condition. Returns empty list if
// no records exist
func GetAllDsvcb20d(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Dsvcb20d))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
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

	var l []Dsvcb20d
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
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

// UpdateDsvcb20d updates Dsvcb20d by Id and returns error if
// the record to be updated doesn't exist
func UpdateDsvcb20dById(m *Dsvcb20d) (err error) {
	o := orm.NewOrm()
	v := Dsvcb20d{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteDsvcb20d deletes Dsvcb20d by Id and returns error if
// the record to be deleted doesn't exist
func DeleteDsvcb20d(id int) (err error) {
	o := orm.NewOrm()
	v := Dsvcb20d{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Dsvcb20d{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
