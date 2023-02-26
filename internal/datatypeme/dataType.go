package datatypeme

import (
	"fmt"
	"time"
)

/**
 * golang数据类型以及字面量
 */
func RunDataType() {
	//整数，字面量类型todo
	/// int32或int64
	var aint int = 10
	fmt.Println(aint)
	/// uint32或uint64
	var auint uint = 10
	fmt.Println(auint)
	var aint8 int8 = 10
	fmt.Println(aint8)
	var auint8 uint8 = 10
	fmt.Println(auint8)
	///类似uint8
	var abyte byte = 10
	fmt.Println(abyte)
	var aint16 int16 = 10
	fmt.Println(aint16)
	var auint16 uint16 = 10
	fmt.Println(auint16)
	var aint32 int32 = 10
	fmt.Println(aint32)
	var auint32 uint32 = 10
	fmt.Println(auint32)
	var aint64 int64 = 10
	fmt.Println(aint64)
	var auint64 uint64 = 10
	fmt.Println(auint64)
	//浮点，字面量类型todo
	var afloat32 float32 = 10.0
	fmt.Println(afloat32)
	var afloat64 float64 = 10.0
	fmt.Println(afloat64)
	//布尔
	var abool bool = false
	fmt.Println(abool)
	//复数，实部，虚部
	var acomplex64 complex64 = 1 + 2i
	fmt.Println(acomplex64)
	fmt.Println(real(acomplex64))
	fmt.Println(imag(acomplex64))
	var acomplex128 complex128 = 1 + 2i
	fmt.Println(acomplex128)
	fmt.Println(real(acomplex128))
	fmt.Println(imag(acomplex128))
	//字符，无此类型
	//字符串
	var astring string = "xcrj"
	fmt.Println(astring)
	//时间
	var atime time.Time = time.Now()
	fmt.Println(atime)
	///时间戳
	fmt.Println(atime.Unix())
	///时间格式化
	fmt.Println(atime.Format("2006/1/02 15:04"))
	//指针
	var bint int = 10
	var bintPointer *int = &bint
	fmt.Println(bintPointer)
	fmt.Println(*bintPointer)
	//数组
	///创建
	var aintArray [3]int = [3]int{}
	fmt.Println(aintArray)
	var bintArray [3]int = [3]int{1, 2, 3}
	fmt.Println(bintArray)
	var cintArray [10]int = [10]int{0: 2, 5: 6}
	fmt.Println(cintArray)
	var dintArray [3][4]int = [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(dintArray)
	///增
	///删
	///改
	///查
	var eintArray [3]int = [3]int{1, 2, 3}
	for i, v := range eintArray {
		fmt.Println(i, ":", v)
	}
	//切片，动态数组
	///创建
	var aslice []int = []int{}
	///len cap
	var bslice []int = make([]int, 1, 4)
	///slice len= 1
	fmt.Println("slice len=", len(bslice))
	///slice cap= 4
	fmt.Println("slice cap=", cap(bslice))
	///增
	aslice = append(aslice, 1)
	///删
	var aaslice []int = aslice[1:]
	fmt.Println(aaslice)
	///改
	///查
	var cslice []int = []int{}
	cslice = append(cslice, 1)
	cslice = append(cslice, 2)
	cslice = append(cslice, 3)
	///index value
	for i, c := range cslice {
		fmt.Println(i, ":", c)
	}
	for _, c := range cslice {
		fmt.Println(c)
	}
	//map
	///创建
	var amap map[string]int = map[string]int{"xcrj1": 1, "xcrj2": 2}
	///len
	var bmap map[string]int = make(map[string]int, 10)
	///map len= 0
	fmt.Println("map len=", len(bmap))
	///增
	amap["xcrj3"] = 3
	///删
	delete(amap, "xcrj3")
	///改
	amap["xcrj3"] = 30
	///查
	for k, v := range amap {
		fmt.Println(k, ":", v)
	}
}
