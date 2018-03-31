package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Dsvcb20c struct {
	Id                       int       `orm:"column(ID);pk" description:"ID"`
	NDEALERID                int       `orm:"column(NDEALERID)" description:"经销商ID"`
	NBRANCHID                int       `orm:"column(NBRANCHID)" description:"网点代码"`
	CSERVICEREQUISITIONNO    string    `orm:"column(CSERVICEREQUISITIONNO);size(13)" description:"委托书，是客户与维修站的一种维修合同，委托单号格式，类别（1）+-+年（4）+月（2）+顺序号（5） 13位"`
	NVEHICLEID               int       `orm:"column(NVEHICLEID)" description:"车辆档案的ID号"`
	CREPAIRSTATUS            string    `orm:"column(CREPAIRSTATUS);size(1)" description:"关于汽车的修理委托书的状态和种类的说明cXLZT='0' 在修，1，结清，2，欠帐"`
	DENTERDATE               time.Time `orm:"column(DENTERDATE);type(date)" description:"汽车进厂维修并开委托单的日期"`
	DPROMISEFINISHDATE       time.Time `orm:"column(DPROMISEFINISHDATE);type(date);null" description:"约定盗车时间，经销商与用户约定的维修车辆的预计完工时间"`
	CFASTREPAIRFLAG          string    `orm:"column(CFASTREPAIRFLAG);size(1);null" description:"快修理"`
	NOPERATORID              int       `orm:"column(NOPERATORID)" description:"系统操作员"`
	NSERVICECONSULTANTID     int       `orm:"column(NSERVICECONSULTANTID)" description:"服务顾问，受理客户预约，负责维修车辆客户的接待，负责车辆的故障诊断，并与用户达成协议（任务委托书），负责车辆维修后的电话服务跟踪。负责向用户"`
	CCONTRACTWORKMODE        string    `orm:"column(CCONTRACTWORKMODE);size(1)" description:"'CBGFS=‘0’，不包，CBGFS=‘1’，包工料，CBGFS=‘2’，包工，若包工方式为''包工料''or''包工''or''此用户工时优惠率''为0 ';"`
	CPARKINGPOSITION         string    `orm:"column(CPARKINGPOSITION);size(10);null" description:"停车车位"`
	NMILEAGE                 int       `orm:"column(NMILEAGE);null" description:"'汽车的行驶里程  单位 万公里"`
	DFIRSTCHECKOUTDATE       string    `orm:"column(DFIRSTCHECKOUTDATE);size(20)" description:"出厂日期"`
	DCHECKDATE               string    `orm:"column(DCHECKDATE);size(20)" description:"车辆维修后，进行完工审核的日期"`
	NEXAMINEID               int       `orm:"column(NEXAMINEID);null" description:"'质量检查员：汽车维修质量的检验，不合格品返修质量的监督，检查。"`
	NCHECKERID               int       `orm:"column(NCHECKERID);null" description:" '进行完工审核的人 他有工时优惠和材料优惠的权利 来源于DXT004';"`
	NCHECKBILLID             int       `orm:"column(NCHECKBILLID);null" description:"'维修的竣工单号 ，可以自动生成单号"`
	NCONTRACTID              int       `orm:"column(NCONTRACTID);null" description:" '修车的合同号 用于防止与客户的纠纷。自动生成';"`
	CCREDENTIALS             string    `orm:"column(CCREDENTIALS);size(20);null" description:"  修车人的证件号';"`
	COWEFLAG                 string    `orm:"column(COWEFLAG);size(1);null" description:" '表示该委托单存在欠款。月结时，对CXLZT=2（欠款）的委托单，打上此标识"`
	NMANHAUR                 float64   `orm:"column(NMANHAUR);null" description:"'修理项目的合计工时。';"`
	NSTANDARDRATIONID        int       `orm:"column(NSTANDARDRATIONID);null" description:"'维修结算使用的工时定额标准，按照车型和客户种类的不同而不同';"`
	NMATECOST                float64   `orm:"column(NMATECOST);null" description:"'一般结算时的材料成本。进价与数量的乘积，包含税。';"`
	NEFFECTIVEMANHAUR        float64   `orm:"column(NEFFECTIVEMANHAUR);null" description:"'总计的加工工时 单位：小时';"`
	CACROSSCURRENTPERIODFLAG string    `orm:"column(CACROSSCURRENTPERIODFLAG);size(1);null" description:" '委托书是否为本期产生的标识 ';"`
	NRESCUEEXPEN             float64   `orm:"column(NRESCUEEXPEN);null" description:"'施救费 指的是外出救援等发生的费用(含税)';"`
	NCONTRACTEXPEN           float64   `orm:"column(NCONTRACTEXPEN);null" description:"'经销商与客户共同确定的维修工时费用，(含税)';"`
	CPROBLEMSETTLEFLAG       string    `orm:"column(CPROBLEMSETTLEFLAG);size(1);null" description:"'问题解决标识';"`
	CMALFUNCTIONDESCRIPTION  string    `orm:"column(CMALFUNCTIONDESCRIPTION);size(100);null" description:"'故障描述';"`
	CREMARK                  string    `orm:"column(CREMARK);size(100);null" description:"没有可录入项，但需要特别说明的信息"`
	DTSTAMP                  time.Time `orm:"column(DTSTAMP);type(timestamp)"`
	NLINKMANID               int       `orm:"column(NLINKMANID);null"`
	NRVTIMES                 int       `orm:"column(NRVTIMES);null" description:"回访次数"`
	CCUSTOMERCODE            string    `orm:"column(CCUSTOMERCODE);size(20);null" description:"客户代码"`
	CCUSTOMERNAME            string    `orm:"column(CCUSTOMERNAME);size(100);null" description:"'客户名称';"`
	CTEL                     string    `orm:"column(CTEL);size(20);null" description:"电话"`
	CMOBILE                  string    `orm:"column(CMOBILE);size(20);null" description:"移动电话"`
	CLINKMANNAME             string    `orm:"column(CLINKMANNAME);size(100);null" description:"联系人名称"`
	CLINKMANTEL              string    `orm:"column(CLINKMANTEL);size(20);null" description:"联系人电话"`
	CLINKMANMOBILE           string    `orm:"column(CLINKMANMOBILE);size(20);null" description:"联系人移动电话"`
	CVIN                     string    `orm:"column(CVIN);size(17);null" description:"底盘号"`
	CLICENCETAG              string    `orm:"column(CLICENCETAG);size(20);null" description:"牌照号"`
	CPOSTNO                  string    `orm:"column(CPOSTNO);size(6);null" description:"邮编"`
	CENGINENO                string    `orm:"column(CENGINENO);size(30);null" description:"发动机号"`
	CRADIOPWD                string    `orm:"column(CRADIOPWD);size(20);null" description:"收音机密码"`
	DLICENCEDATE             string    `orm:"column(DLICENCEDATE);size(20);null" description:"发牌日期"`
	CPRODUCTCODE             string    `orm:"column(CPRODUCTCODE);size(18);null" description:"产品代码"`
	CLINKMANADDRESS          string    `orm:"column(CLINKMANADDRESS);size(100);null" description:"联系人地址"`
	CLINKMANPOSTNO           string    `orm:"column(CLINKMANPOSTNO);size(6);null" description:"联系人邮编"`
	NLINKMANPROVINCEID       int       `orm:"column(NLINKMANPROVINCEID);null" description:"联系人省区ID"`
	NLINKMANCITYID           int       `orm:"column(NLINKMANCITYID);null" description:"联系人市县ID"`
	CVEHICLETYPECODE         string    `orm:"column(CVEHICLETYPECODE);size(30);null" description:"车型代码"`
	CVEHICLENAME             string    `orm:"column(CVEHICLENAME);size(50);null" description:"车型名称"`
	CCUSTOMERADDRESS         string    `orm:"column(CCUSTOMERADDRESS);size(100);null" description:"客户地址"`
	CWASHCARCODE             string    `orm:"column(CWASHCARCODE);size(1);null" description:"洗车标识，是否洗车"`
	DSALEDATE                time.Time `orm:"column(DSALEDATE);type(date);null" description:"车辆销售日期"`
	DNEXTBYDATE              time.Time `orm:"column(DNEXTBYDATE);type(date);null" description:"下次保养时间"`
	NNEXTBYMILES             float64   `orm:"column(NNEXTBYMILES);null" description:"下次保养里程"`
	CPDIS                    string    `orm:"column(CPDIS);size(1);null" description:"PDIS标识"`
	COUTSERVICEFLAG          string    `orm:"column(COUTSERVICEFLAG);size(1);null" description:"外出服务标识"`
	CREVOKEFLAG              string    `orm:"column(CREVOKEFLAG);size(1);null" description:"作废标识"`
	CONTIMEFLAG              string    `orm:"column(CONTIMEFLAG);size(1);null" description:"按时交车状态“是”为按时交车，“否”为未按时交车，“空白”如果车辆未结算或者未做服务顾问打印结算单则显示空白';"`
	CFEEFORESEEFLAG          string    `orm:"column(CFEEFORESEEFLAG);size(1);null" description:"费用预估状态是”为预估准确，“否”为预估不准确，“空白”如果车辆未结算';"`
	DFIRSTPRINTDATE          string    `orm:"column(DFIRSTPRINTDATE);size(20);null" description:"'服务顾问第一次打印结算单时间';"`
	NPRINTPERSONID           int       `orm:"column(NPRINTPERSONID);null" description:"第一次打印人"`
	DWTSFIRSTPRINTDATE       string    `orm:"column(DWTSFIRSTPRINTDATE);size(20);null" description:"'服务顾问第一次打印委托书时间';"`
	NWTSPRINTPERSONID        int       `orm:"column(NWTSPRINTPERSONID);null" description:"委托书第一次打印人';"`
	CDIAGNOSENOPASSFLAG      string    `orm:"column(CDIAGNOSENOPASSFLAG);size(1);null" description:"'一次性问诊未通过标识(0通过一次性问诊,1未通过一次性问诊)';"`
	CITEMADDFLAG             string    `orm:"column(CITEMADDFLAG);size(1);null" description:"'项目新增标识';"`
	CITEMDELFLAG             string    `orm:"column(CITEMDELFLAG);size(1);null" description:"'项目删除标识';"`
	CITEMUPDFLAG             string    `orm:"column(CITEMUPDFLAG);size(1);null" description:"'项目修改标识';"`
	CBJADDFLAG               string    `orm:"column(CBJADDFLAG);size(1);null" description:"备件预估新增标识';"`
	CBJDELFLAG               string    `orm:"column(CBJDELFLAG);size(1);null" description:"'备件预估删除标识';"`
	CBJUPDFLAG               string    `orm:"column(CBJUPDFLAG);size(1);null" description:"'备件预估修改标识';"`
	CWTSSECONDPRINTFLAG      string    `orm:"column(CWTSSECONDPRINTFLAG);size(1);null" description:"委托书变为一次性问诊未通过后是否再次打印';"`
	NTRAFFICSUBSIDY          float64   `orm:"column(NTRAFFICSUBSIDY);null" description:"交通补助"`
	CTRAFFICCARFLAG          string    `orm:"column(CTRAFFICCARFLAG);size(1);null" description:"是否备用车　1是0否';"`
	CROADTRIDE               string    `orm:"column(CROADTRIDE);size(1);null" description:"是否路试"`
	CSPBJNOCLM               string    `orm:"column(CSPBJNOCLM);size(1);null" description:"备件维修委托书标识（默认0, 为1时不参与三包保修记录统计）';"`
	CCHANGEORSCRAPFLAG       string    `orm:"column(CCHANGEORSCRAPFLAG);size(1);null" description:"退换车或报废车标识';"`
	NADVISORID               int       `orm:"column(NADVISORID);null" description:"'专属服务顾问id';"`
	CUPLOADFLAG              string    `orm:"column(CUPLOADFLAG);size(1);null" description:"健康档案上传标识';"`
	CUPLOADDATE              string    `orm:"column(CUPLOADDATE);size(20)" description:"健康档案上传时间';"`
	COTOCODE                 string    `orm:"column(COTOCODE);size(15);null" description:"o2o代码';"`
}

func (t *Dsvcb20c) TableName() string {
	return "dsvcb20c"
}

func init() {
	orm.RegisterModel(new(Dsvcb20c))
}

// AddDsvcb20c insert a new Dsvcb20c into database and returns
// last inserted Id on success.
func AddDsvcb20c(m *Dsvcb20c) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetDsvcb20cById retrieves Dsvcb20c by Id. Returns error if
// Id doesn't exist
func GetDsvcb20cById(id int) (v *Dsvcb20c, err error) {
	o := orm.NewOrm()
	v = &Dsvcb20c{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllDsvcb20c retrieves all Dsvcb20c matches certain condition. Returns empty list if
// no records exist
func GetAllDsvcb20c(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Dsvcb20c))
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

	var l []Dsvcb20c
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

// UpdateDsvcb20c updates Dsvcb20c by Id and returns error if
// the record to be updated doesn't exist
func UpdateDsvcb20cById(m *Dsvcb20c) (err error) {
	o := orm.NewOrm()
	v := Dsvcb20c{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteDsvcb20c deletes Dsvcb20c by Id and returns error if
// the record to be deleted doesn't exist
func DeleteDsvcb20c(id int) (err error) {
	o := orm.NewOrm()
	v := Dsvcb20c{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Dsvcb20c{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
