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

 /**
  泛型切片 map chan
  泛型函数
  泛型结构体 接口
  泛型接收器
  没有泛型方法
  */

//泛型slice
func RunGeneric00(){
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

//泛型map
func RunGeneric01(){
	type IntStrMap map[int]string;
	type IFSMap[K int|float32|string,V int|float32|string] map[K]V
	var iFSMap IFSMap[string,int]=map[string]int{}
	iFSMap["xcrj"]=0
	fmt.Println(iFSMap);
}

//泛型通道
func RunGeneric02(){
	type IntChan chan int
	type IFSChan[T int|float32|string] chan T
	//make(chan int,10) 最大capacity
	var ifsChan IFSChan[int]=make(chan int,1)
	ifsChan<-10
	val:=<-ifsChan
	fmt.Println(val)
}

/**
 泛型函数(Generic function)
 */
func add[T int|int64|float32|float64](a T,b T)(r T){
	r=a+b
	return r
}
func RunGeneric10(){
	var a int=10
	var b int=20
	var r=add[int](a,b)
	fmt.Println(r)
}

//泛型结构体
type Student[T int|string] struct{
	Code T
	Name string
}
func RunGeneric20(){
	var stu Student[int]=Student[int]{123,"xcrj"}
	fmt.Println(stu)
}

//泛型接口, 由方法集合 向 类型集合转变
type Apple[T int|string] interface{
	GetPrice(price T)
}
type FujiApple struct{
	Name string
}
func (fujiApple FujiApple) GetPrice(price int){
	fmt.Println(fujiApple.Name,": ",price)
}
func RunGeneric21(){
	var apple Apple[int]=FujiApple{"fuji apple"}
	apple.GetPrice(10)
}

/**
 泛型接收器(Generic receiver)
 */
type IFSlice[T int|float32] []T
// 入参，出参，函数体都可以使用泛型
func(ifSlice IFSlice[T]) sum()T{
	var sum T
	for _,num:=range ifSlice{
		sum+=num
	}
	return sum
}
func RunGeneric30(){
	var ifSlice IFSlice[int]=[]int{}
	ifSlice=append(ifSlice,1)
	ifSlice=append(ifSlice,2)
	ifSlice=append(ifSlice,3)
	var sum int=ifSlice.sum()
	fmt.Println(sum)
}

/**
 泛型接收器(Generic receiver)
 泛型队列
 interface{} 代码任意类型 可以用 any 代替
 */
type Queue[T interface{}] struct{
	elements []T
}
func(queue *Queue[T]) Offer(element T){
	queue.elements=append(queue.elements,element)
}
//出参2：队列是否为空
func(queue *Queue[T]) Poll()(T,bool){
	var r T
	if len(queue.elements)==0{
		return r,true
	}
	r=queue.elements[0];
	queue.elements=queue.elements[1:];
	return r,len(queue.elements)==0
}
func(queue *Queue[T]) Size()int{
	return len(queue.elements)
}
func RunGeneric31(){
	var iQue Queue[int]
	iQue.Offer(1)
	iQue.Offer(2)
	iQue.Offer(3)
	v,ok:=iQue.Poll()
	fmt.Println(v,", ",ok)
	fmt.Println(iQue.Size())
}
