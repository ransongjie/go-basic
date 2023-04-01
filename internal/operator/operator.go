package operator

import "fmt"

//赋值, 算数, 关系, 逻辑, 位, 指针
func RunOperator() {
	//赋值, =, +=, -= ...
	var aint int = 10
	var bint int = 20
	bint += 1
	bint -= 1
	fmt.Println(aint)
	fmt.Println(bint)

	//算数, +, -, *, /, %, ++, --
	fmt.Println(aint + bint)
	fmt.Println(aint - bint)
	fmt.Println(aint * bint)
	fmt.Println(aint / bint)
	fmt.Println(aint % bint)
	///必须单独写为一行
	aint++
	fmt.Println(aint)
	aint--
	fmt.Println(aint)
	///无先加，先减
	///++aint
	///fmt.Println(aint)
	///--aint
	///fmt.Println(aint)

	//关系, ==, !=, >, <, >=, <=
	fmt.Println(aint == bint)
	fmt.Println(aint != bint)
	fmt.Println(aint > bint)
	fmt.Println(aint < bint)
	fmt.Println(aint >= bint)
	fmt.Println(aint <= bint)

	//逻辑, &&, ||, !
	var abool bool = true
	var bbool bool = false
	fmt.Println(abool && bbool)
	fmt.Println(abool || bbool)
	fmt.Println(!abool)

	//位, &, |, ^ (异或), <<, >>
	var cint int = 10
	var dint int = 20
	fmt.Println(cint & dint)
	fmt.Println(cint | dint)
	fmt.Println(cint ^ dint)
	///不改变符号位
	fmt.Println(cint >> 1)
	///不改变原数值
	fmt.Println(cint)
	///不改变符号位
	fmt.Println(cint << 1)
	///不改变原数值
	fmt.Println(cint)

	//指针, &, *
	var fint int = 10
	var fintPointer *int = &fint
	///指针址
	fmt.Println(fintPointer)
	///指针值
	fmt.Println(*fintPointer)
}
