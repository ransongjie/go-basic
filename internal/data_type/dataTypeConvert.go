package data_type

import (
	"fmt"
	"strconv"
)

/**
 *数据类型转换
 *注意每种类型的可表示范围
 */
func RunDataTypeConvert() {
	//整数转整数
	var cint int = 3
	var cint32 int32 = int32(cint)
	fmt.Println(cint32)

	//整数转浮点数
	var aint int = 10
	var afloat32 float32 = float32(aint)
	fmt.Println(afloat32)
	//浮点数转整数
	var bfloat32 float32 = 3.2
	var bint int = int(bfloat32)
	fmt.Println(bint)

	//整数转string
	var dint int = 3
	var dstring string = strconv.Itoa(dint)
	fmt.Println(dstring)
	//string转整数
	var estring string = "3"
	eint, err := strconv.Atoi(estring)
	fmt.Println(eint)
	fmt.Println(err)

	//浮点数转string
	var ffloat32 float32 = 2.3
	///保留2位小数，float32
	var fstring string = strconv.FormatFloat(float64(ffloat32), 'f', 2, 32)
	fmt.Println(fstring)
	//string转浮点数
	var hstring string = "2.3"
	hfloat32, err := strconv.ParseFloat(hstring, 32)
	fmt.Println(hfloat32)
	fmt.Println(err)
	//复数

	//字符串

	//布尔

	//指针

	//数组转切片
	var aarray [3]int = [3]int{1, 2, 3}
	var aslice []int = aarray[:]
	aslice = append(aslice, 4)
	fmt.Println(aslice)
	//切片转数组

	//map
}
