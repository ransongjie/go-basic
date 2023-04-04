package generic
import "fmt"

//一般接口，含有类型限制的接口
type PPP interface{
	int|int8|int32
}

//一般接口，含有类型限制的接口
//限制为自定义基本类型
type IFI interface{
	~int|~float32
	RInt() int
	RFloat() float32
}
type Intme int
func(im Intme) RInt() int{
	return 100
}
func(im Intme) RFloat() float32{
	return 100.2
}
func RunTF1(){
	var im Intme=10
	var rint int =im.RInt()
	var rfloat float32=im.RFloat()
	fmt.Println(rint,", ",rfloat)
}

//一般接口，含有类型限制的接口
//限制为结构体
type Phone interface{
	Iphone|Huawei
	Name() string
	Price() float32
}
type Iphone struct{}
type Huawei struct{}
func(iphone Iphone) Name() string{
	return "Iphone"
}
func(iphone Iphone) Price() float32{
	return 200.2
}
func(huawei Huawei) Name() string{
	return "Huawei"
}
func(huawei Huawei) Price() float32{
	return 200.1
}
func RunTF2(){
	var iphone Iphone=Iphone{}
	var name string =iphone.Name()
	var price float32=iphone.Price()
	fmt.Println(name,", ",price)

	var huawei Huawei=Huawei{}
	var name2 string =huawei.Name()
	var price2 float32=huawei.Price()
	fmt.Println(name2,", ",price2)
}

//基本接口，只含有方法的接口
type Fruit interface{
	Name() string
}
type Cabbage struct{}
func (cabbage Cabbage) Name()string{
	return "cabbage"
}
func RunTF3(){
	var cabbage Fruit=Cabbage{}
	var name string=cabbage.Name()
	fmt.Println(name)
}