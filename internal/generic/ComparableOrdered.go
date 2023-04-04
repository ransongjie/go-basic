package generic

//可比较 comparable golang内置接口
//== !=
type MyMapC[K comparable,V any] map[K]V

//可排序 ordered
//> < >= <=
// Ordered 代表所有可比大小排序的类型
type Ordered interface {
    Integer | Float | ~string
}
type Integer interface {
    Signed | Unsigned
}
type Signed interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64
}
type Unsigned interface {
    ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}
type Float interface {
    ~float32 | ~float64
}