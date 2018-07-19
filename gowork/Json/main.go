package main

import (
	"encoding/json"
	"strconv"
	//"encoding/json"
	"fmt"
	//	"reflect"
)

//当出现大量数据，内存不足以缓存那么多的数据时，将其序列化，使数据持久化，存储在电脑磁盘里
//结构体的字段加tag `json:"usename"`
//`json:"age_int,omitempty"` 后面加上omitempty,可以忽略nil值或者0值
//`json:"-"` 后面加上-,直接忽略这一字段
//只序列化能够导出的字段，也就是字段的首字母大写，并且结构体的首字幕也要大写。
//主要原理是通过reflect包，也即是反射来解决的这一问题
//对于结构体中的匿名字段，也要加上变量名，不让会出错
//通过包"encoding/json"的 func Marshal(v interface{}) ([]byte, error) 来完成数据序列化
//通过包"encoding/json"的 func Unmarshal(data []byte, v interface{}) error 来完成数据反序列化
//还可以通过重写json包的  MarshalJSON() (date []byte, err error) 方法来自定义序列化
//还可以通过重写json包的  UnmarshalJSON(date []byte) error 方法来自定义反序列化
//反序列化的时候记得接受者必须为指针类型。
type (
	PersonInfo struct {
		Name string      `json:"usename"`
		Age  int         `json:"age_int,omitempty"`
		Pho  Phone_info  `json:"phone"`
		Sch  School_info `json:"level"`
	}
	Phone_info struct {
		Phone string
	}
	School_info struct {
		Level int
	}
)

func (p Phone_info) MarshalJSON() (date []byte, err error) {

	if len(p.Phone) != 11 {
		err = fmt.Errorf("请输入11位手机号")
		return

	} else {
		date = []byte(p.Phone)
		return

	}

}

func (p School_info) MarshalJSON() (date []byte, err error) {

	if p.Level < 10 {
		err = fmt.Errorf("请选择难度")
		return

	} else {
		date = []byte(fmt.Sprintf("%d", p.Level))
		return

	}

}

func (p *Phone_info) UnmarshalJSON(date []byte) error {

	//	if len(p.Phone) != 11 {

	//		return fmt.Errorf("not full")

	//	}

	p.Phone = string(date)
	return nil

}

func (p *School_info) UnmarshalJSON(date []byte) error {

	//	if len(p.Phone) != 11 {

	//		return fmt.Errorf("not full")

	//	}

	p.Level, _ = strconv.Atoi(string(date))
	return nil

}

func main() {

	p := PersonInfo{Name: "tom", Age: 20, Pho: Phone_info{"12345678910"}, Sch: School_info{Level: 11}}

	date, err := json.Marshal(p)
	if err != nil {
		fmt.Println("err:", err)

	} else {

		fmt.Println(string(date))

	}

	var p1 PersonInfo
	err1 := json.Unmarshal(date, &p1)
	if err1 != nil {

		fmt.Println("err1:", err1)

	}
	fmt.Println(p1)

}
