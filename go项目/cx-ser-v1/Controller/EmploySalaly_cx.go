package Controller

import (
	"fmt"
	"time"
)
type statement struct {
	name string
	workyear string
	station   string    //岗位
	classtime  float64 //课时量
	stay string //住宿
	pmumber  string// 党员
	details string//完整的打印信息
	wage float64//工资
	t1 time.Time
	t2 time.Time
	month int

}

func Createstament() *statement{
	return &statement{
		name:"",
		workyear: "",
		station :"",
		classtime  :0.0,
		stay :"",
		pmumber  :"",
		wage :5000.0,
		details : "\t姓名\t 薪资",
	}
}
//录入信息
func (m_cx *statement)showIn(){
	fmt.Print("姓名：")
	fmt.Scan(&m_cx.name)
	m_cx.contMonth() //调用contMonth方法
	fmt.Print("岗位")
	fmt.Scan(&m_cx.station)
	fmt.Print("课时量")
	fmt.Scan(&m_cx.classtime)
	fmt.Print("是否住宿")
	fmt.Scanln(&m_cx.stay)
	fmt.Print("是否是党员")
	fmt.Scanln(&m_cx.pmumber)
	m_cx.EmploySalayCount()
}

func (m_cx *statement) contMonth() {
	fmt.Println("入职时间：")
	y1 := m_cx.t1.Year()
	m1 := int(m_cx.t1.Month())
	d1 := m_cx.t1.Day()
	fmt.Scan(&y1,&m1,&d1)
	fmt.Println("当前时间")
	y2 := m_cx.t2.Year()
	m2 := int(m_cx.t2.Month())
	d2 := m_cx.t2.Day()
	fmt.Scan(&y2,&m2,&d2)
	//time.Parse(t1, "2021-03-15")
	yearInterval := y2 - y1
	if m1 < m2 || m1 == m2 && d1 < d2 {
		yearInterval--
	}
	monthInterval := (m1 + 12) - m2
	if d1 < d2 {
		monthInterval--
	}
	monthInterval %= 12
	m_cx.month = yearInterval*12 + monthInterval
}
func (m_cx *statement) EmploySalayCount (){
	switch {
	case  m_cx.station =="教师" :
		m_cx.wage+=(40.0)*m_cx.classtime*0.85
	case m_cx.station =="教研室主任" :
		m_cx.wage+=float64(m_cx.month*300)
	case m_cx.station =="行政岗位":  //行政岗位
		m_cx.wage-=500
	}
	if m_cx.month>=12{
		m_cx.wage+=500
	}

	if	m_cx.pmumber=="是" {//判断是否是党员
		switch  {
		case m_cx.wage<=3000  :
			m_cx.wage=m_cx.wage*(1-0.005)
		case m_cx.wage>3000&& m_cx.wage<5000  :
			m_cx.wage=m_cx.wage*(1-0.01)
		case m_cx.wage>=5000&&m_cx.wage<=10000  :
			m_cx.wage=m_cx.wage*(1-0.15)
		case m_cx.wage>10000  :
			m_cx.wage=m_cx.wage*(1-0.2)
		}
	}
	if m_cx.stay =="是" {
		m_cx.wage-=300
	}
	m_cx.details += fmt.Sprintf("\n\t%v\t%v",m_cx.name,m_cx.wage)
	fmt.Print(m_cx.details)
}
//打印信息
//func (m_cx *statement)showDetails(){
//	fmt.Println("------------------------打印姓名 工资-----------------------------------")
//	fmt.Println(m_cx.details)
//
//}

//程序退出
//func (m_cx *statement)exit() bool {
//	flag := ""
//	loop := true
//	for {
//		fmt.Print("你确定要退出吗？y/n\t")
//		fmt.Scanln(&flag)
//		if flag == "y"{
//			break
//		}else if flag == "n"{
//			loop = false
//		} else {
//			continue
//		}
//		break
//	}
//	return loop
//}
//主窗口
//func (m_cx *statement) MainMeau() {
//	var mode int64
//	loop := false
//	for  {
//		//打印展示台信息
//		fmt.Println("---------------------------财务报表---------------------------")
//		fmt.Println("                           1、打印信息（姓名 -工资")
//		fmt.Println("                           2、录入")
//		fmt.Println("                           3、退    出")
//		fmt.Print("请输入【1-3】：")
//		fmt.Scanln(&mode)
//		switch mode{
//		case 1:
//			m_cx.showDetails()//打印所有信息
//		case 2:
//			m_cx.showIn()//录入
//		case 3:
//			loop = m_cx.exit()
//		}
//		if loop {
//			return
//		}
//		fmt.Println("你退出了使用！")
//	}
//}

