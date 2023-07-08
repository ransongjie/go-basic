# golang
|Directory|Description|
|---|---|
|data_type|Data type|
|operator|Operator|
|control|Control structure|
|func|file function|
|oop|struct and interface|
|err|error handling|
|concurrent|concurrent|
|generic|generic paradigm|
# data_type
[数据类型](https://blog.csdn.net/baidu_35805755/article/details/128961290)
[数据类型转换](https://blog.csdn.net/baidu_35805755/article/details/128966050)
# operator
[运算符](https://blog.csdn.net/baidu_35805755/article/details/128993879)
# control
[代码逻辑](https://blog.csdn.net/baidu_35805755/article/details/129196215)
# generic
泛型
- 使用时机：经常要分别为不同的类型写完全相同逻辑的代码，使用泛型将是最合适的选择
- 注意：任何泛型类型都必须传入类型实参 实例化后才可以使用

可用
- 泛型切片, map, chan
- 泛型函数
- 泛型结构体, 接口, 接收器, 没有泛型方法
# concurrent
- 任务, goroutine, MPG模型, chan, 锁, sync
- 主线程执行完毕协程退出

MPG: M(main 主线程) P(协程执行所需上下文) G(协程)
- M0上G0被阻塞, 新开M1或从线程池中拿M1执行M0上G1...

# err
异常
- 程序退出，正常退出 return，异常退出 panic异常 未recover
- 异常处理，return异常 err!=nil。defer recover异常。
- 退出处理，defer 释放资源 异常处理