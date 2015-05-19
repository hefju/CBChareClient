package models

import (
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	_ "github.com/lunny/godbc"
	"log"
	"time"
)

var engine *xorm.Engine

func init() {
	var err error
	//	engine, err = xorm.NewEngine("odbc", "driver={SQL Server};Server=192.168.1.200; Database=charge; uid=sa; pwd=123;")
	engine, err = xorm.NewEngine("odbc", "driver={SQL Server};server=.;database=charge;integrated security=SSPI;")

	if err != nil {
		log.Fatalln("xorm create error", err)
	}
	//engine.ShowSQL = true
	engine.SetMapper(core.SameMapper{})
	// engine.CreateTables(new(tp_charge_billing))
	err = engine.Sync2(new(Tp_charge_billing)) //, new(Group))
	if err != nil {
		log.Fatalln("xorm sync error", err)
	}
}

//传入的时间已经处理过, 是前一天的时间.
func GetChargeListByDate(date time.Time) []Tp_charge_billing {
//	date1 := date.Format("2006-01-02 00:00:00")
//	date = date.AddDate(0, 0, 1)
//	date2 := date.Format("2006-01-02 00:00:00")
    	date1 := date.Format("2006-01-02") + " 18:00:00"//昨晚6点
    	date = date.AddDate(0, 0, 1)
    	date2 := date.Format("2006-01-02") + " 06:00:00"//今日早上6点


    bills := make([]Tp_charge_billing, 0)
	//  engine.Where("Crt_date>='?' and Crt_date<'?'",date1,date2).Find(&bills)
	engine.Where("Crt_date>=? and Crt_date<?", date1, date2).Find(&bills)
	//log.Println("bills length:",len(bills))
	return bills
}

type Tp_charge_billing struct {
	Sn             int `xorm:"pk"`
	Chg_id         int
	Area_id        int
	Chg_port       int
	Chg_sn         string
	Opt_model      int
	Card_type      int
	Card_id        string
	Card_money     int
	Card_money_end int
	Chg_model      int
	Chg_para       int
	Charge         int
	Chg_pw         int
	Chg_time       int
	Pw_total       int
	Pw_total_end   int
	Soc_st         int
	Soc_ed         int
	Ch_st_date     string
	Ch_ed_date     string
	St_date        string
	Ed_date        string
	Ed_code        int
	Crt_date       time.Time
}

//报告类
type StatusReport struct  {
    Id int64
    From string      //发送人
    FromTime time.Time  //发送的时间
    Title string //标题(分类: 健康,统计的,)
    Content string //详细内容
}
