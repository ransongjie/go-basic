package data_type

import (
	"fmt"
	"unsafe"
)

func RunDataTypeLen() {
	//布尔
	var abool bool = true
	len := unsafe.Sizeof(abool)
	fmt.Println(len)

	//整数
	var aint int = 10
	len = unsafe.Sizeof(aint)
	fmt.Println(len)
	var auint uint = 20
	len = unsafe.Sizeof(auint)
	fmt.Println(len)
}
