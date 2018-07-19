package main

//包reflect中两个结构体Type,Value
//分别对应对象的类型和值的数据,以及两个重要的函数
//reflect.TypeOf(i interface{}) Type
//reflect.TypeOf()返回值的类型就是reflect.Type。
//reflect.ValueOf(i interface{}) Value
//reflect.ValueIOf()返回值的类型就是reflect.Value
//reflect.Type几个比较常用的方法：
//获取返回值的类型  Kind()Kind 或者String()string、Name()string
//获取返回值的包的路径PkgPath() string
//获取返回值的字段 主要通过两个方法
//获取字段的数量 NumField() int
//获取字段Field(i int) StructField
//获取方法的数量 NumMethod() int
//获取方法名Method(int) Method 再通过访问Method结构体中的字段Name,就可以得到方法名了
//返回的字段为StructField类型，结构体StructField中主要的三个字段： Name string、Type Type、 Tag StructTag
//通过访问Name string来查询字段名，访问Type Type来访问字段的类型，访问Tag StructTag来访问字段的标签
//获取标签需要用到 func (tag StructTag) Get(key string) string方法

//reflect.Value的几个常用的方法
// 判断 v 值是否可以被修改。只有可寻址的 v 值可被修改。
// 结构体中的非导出字段（通过 Field() 等方法获取的）不能修改，所有方法不能修改。
//func (v Value) CanSet() bool 函数返回true表示可以修改这个v
// 获取“指针所指的对象”或“接口所包含的对象”
//func (v Value) Elem() reflect.Value 只有通过该函数返回的value才能使得CanSet（）函数返回true并且Elem()的接受者也必须为指针类型
// 获取 v 值的字符串描述
//func (v Value) String() string
// 获取 v 值的类型
//func (v Value) Type() reflect.Type
// 返回 v 值的类别，如果 v 是空值，则返回 reflect.Invalid。
//func (v Value) Kind() reflect.Kind
// 获取 v 的方法数量
//func (v Value) NumMethod() int
// 结构体
// 获取 v 值的字段数量，v 值必须是结构体。
//func (v Value) NumField() int
// 根据索引获取 v 值的字段，v 值必须是结构体。如果字段不存在则 panic。
//func (v Value) Field(i int) reflect.Value

func main() {

}
