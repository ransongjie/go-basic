package generic

import(
	"fmt"
)     
/**
 泛型
 类型形参(Type parameter)
 类型实参(Type argument)
 类型形参列表(Type parameter list)
 类型约束(Type constraint)
 泛型类型(Generic type)
 实例化(Instantiation)
 */
func RunGeneric(){
	type IntSlice []int//整型切片
	type Float32Slice []float32
	type StrSlice []int
	/**
	 T 类型形参
	 T int|float32|string 类型形参列表
	 int|float32|string 类型约束
	 IFSSlice 泛型类型
	 */
	type IFSSlice[T int|float32|string] []T
	/**
	 IFSSlice[int] 类型实参
	 IFSSlice[T int|float32|string] 类型形参
	 ifss 实例
	 */
	var ifss IFSSlice[int]=[]int{}
	ifss=append(ifss,1)
	fmt.Println(ifss)
}

/**
 泛型接收器(Generic receiver)
 */
func RunGeneric2(){

}

/**
 泛型函数(Generic function)
 */
 func RunGeneric3(){

}