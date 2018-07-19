//go语言mysql数据库操作
package main

import (
	"fmt"

	"github.com/astaxie/beego/orm"     //导入beego包orm
	_ "github.com/go-sql-driver/mysql" //导入mysql数据库驱动
)

//初始化函数
func init() {
	//对mysql数据库操作步骤：
	//首先注册数据库，必须有一个default，否则会报错。
	//使用beego包orm.RegisterDataBase（）函数注册数据库
	//可以同时注册多个数据库

	//注册模型 使用beego包orm.RegisterModel(e ...interface{})函数，可以同时注册多个对象
	//生成表	使用beego包orm.RunSyncdb("hx", false, true)选择哪一个数据库，报错调到下一个指令否，是否打印sql执行结果

	// set default database
	orm.RegisterDataBase("default", "mysql", "root:HX141656@tcp(127.0.0.1:3306)/hx?charset=utf8", 30)
	orm.RegisterDataBase("hx", "mysql", "root:HX141656@tcp(127.0.0.1:3306)/hx?charset=utf8", 30)
	//orm.RegisterDataBase("tmj", "mysql", "root:HX141656@tcp(127.0.0.1:3306)/tmj?charset=utf8", 30)
	//regist model
	orm.RegisterModel(new(User))

	// create table
	orm.RunSyncdb("hx", false, true)
}

func main() {

	o := orm.NewOrm()    //获取一个orm对象
	err := o.Using("hx") //选择已经注册的数据库

	//	一下就可以用orm对象完成对数据库的操作

	//	插入操作
	//	u := User{Name: "jack"}
	//	id, err := o.Insert(&u)
	//	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	//事务处理
	//	o.Begin()
	//	o.Rollback()
	//	o.Commit()

	//	可以直接发送SQL语句，下面是获取多条查询的数据
	//	var ids []int
	//	var names []string
	//	o.Raw("select id,name from User").QueryRows(&ids, &names)
	//	fmt.Println(ids, names)

	//下面是获取多条查询的数据
	//		var m []orm.Params
	//	id, err := o.Raw("SELECT * FROM User").Values(&m)
	//	fmt.Printf("ID: %d, ERR: %v\n", id, err)
	//	for _, term := range m {
	//		fmt.Println(term["id"], ":", term["name"])
	//	}

	//	下面是获取单挑查询的数据
	//	var id int
	//	var name string
	//	o.Raw("select id,name from User where id = ?", 1).QueryRow(&id, &name)
	//	fmt.Println(id, name)

	//	还可以执行删除和修改操作

}
