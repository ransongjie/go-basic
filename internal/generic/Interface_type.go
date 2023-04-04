package generic

/**
 接口 由 方法集合体 到 类型集合体
 */

//简化写法
type ASlice[T int|int8|int16|int32|int64|uint|uint8|uint16] []T
type AIFace interface{
	int|int8|int16|int32|int64|uint|uint8|uint16
}
type ASlice2[T AIFace] []T

//组合写法
type IntIFace interface{
	int|int8|int16|int32|int64
}
type UintIFace interface{
	uint|uint8|uint16|uint32|uint64
}
type FloatIFace interface{
	float32|float64
}
type BSlice[T IntIFace|UintIFace|FloatIFace|string] []T

//类型并集 |
type MyUnion interface{
	int|int8|int16
}

//类型交集 换行
type MyInter interface{
	int|int8|int16
	int
}
type MySliceInter[T MyInter] []T
var mySliceInter MySliceInter[int] 
// var mySliceInter MySliceInter[int8]//报错

//所有类型的集合 空接口
type AllType[T interface{}] []T
//等价值写法
type AllType2[T any] []T

//~底层类型
/**
 不允许的情况
 1. ~非基本类型
 2. ~接口
 允许的情况
 1. ~[]byte
 */
//不允许底层类型
var bSlice BSlice[int]
type BInt int
// var bSlice2 BSlice[BInt]//将报错

//允许底层类型
type IntIFace2 interface{
	~int|~int8|~int16|~int32|~int64
}
type UintIFace2 interface{
	~uint|~uint8|~uint16|~uint32|~uint64
}
type FloatIFace2 interface{
	~float32|~float64
}
type BSlice2[T IntIFace2|UintIFace2|FloatIFace2|~string] []T
var cSlice BSlice2[int]
type BInt2 int
var cSlice2 BSlice2[BInt2]