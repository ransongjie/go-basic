package obj

import "fmt"
//结构体
type Teacher struct{
	//属性
	Name string
	Age int
	Salary float32
}
//方法
func (teacher Teacher) getName(){
	fmt.Println(teacher.Name)
}

//
func RunStruct1(){
	var teacher Teacher=Teacher{"xcrj",30,20.2}
	fmt.Println(teacher)
	teacher.getName()
}

//new
func RunStruct2(){
	var pteacher *Teacher
	pteacher=new(Teacher)//new 分配空间，默认初始化，返回指针类型
	pteacher.Name="xcrj"
	pteacher.Age=30
	pteacher.Salary=20.2
	fmt.Println(*pteacher)

	pteacher=&Teacher{"xcrj",30,20.2}
	fmt.Println(*pteacher)
}

type Car struct{
	Motor
	Name string
}
type Motor struct{
	Name string
}

//组合代替继承
//组合中部分与整体同生共死
//聚合中部分与整体不同生共死
func RunStruct3(){
	car:=Car{Motor{"motor 520"},"BYD"}
	fmt.Println(car)
}