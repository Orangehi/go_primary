package main

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

const (
	_DB_NAME      = "root:HX141656@/hx_shiyan?charset=utf8"
	_MYSQL_DRIVER = "mysql"
)

type UserName struct {
	Id       int
	Name     string
	Password int
}

func init() {

	// 注册模型
	//orm.RegisterModel(new(Category), new(Topic), new(User))
	orm.RegisterModel(new(UserName))
	//注册sql驱动
	orm.RegisterDriver(_MYSQL_DRIVER, orm.DRMySQL)

	// 注册默认数据库
	// 我的mysql的root用户密码为HX141656，打算把数据表建立在名为hx_shiyan数据库里
	// 备注：此处第一个参数必须设置为“default”（因为我现在只有一个数据库），
	//否则编译报错说：必须有一个注册DB的别名为 default
	orm.RegisterDataBase("default", _MYSQL_DRIVER, _DB_NAME, 10)

}

func main() {

	o := orm.NewOrm()

	orm.RunSyncdb("default", false, true)

	u := UserName{Name: "tom", Password: 123456}

	id, err := o.Insert(&u)
	if err != nil {
		fmt.Println("errinsert:", err)
	} else {
		fmt.Println("id:", id)
	}

	u = UserName{Name: "bob", Password: 78910}
	id, err = o.Insert(&u)
	if err != nil {
		fmt.Println("errinsert2:", err)
	} else {
		fmt.Println("id2:", id)
	}

	// update
	u.Name = "astaxie"
	u.Password = 123456
	u.Id = 1
	num, err := o.Update(&u)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	// read one
	//u1 := UserName{Id: 1}
	u1 := &UserName{Name: "astaxie", Password: 78910}
	err = o.Read(u1, "Name", "Password")
	fmt.Printf("ERR: %v\n", err)

}
