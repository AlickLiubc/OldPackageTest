package main

import "fmt"

func Add(a, b int) int {
	total := a + b
	return total
}

type Company struct {
	Name    string
	Address string
}

type Employee struct {
	Name    string
	company Company
}

type PrintResult struct {
	Info string
	Err  error
}

//func RpcPrintln(employee Employee) PrintResult {
// rpc中的第2个点 传输协议，数据编码协议
// http1.x, http2.0版本协议
// http协议底层使用的也是TCP协议 http目前主流的是1.X，这种协议有性能问题 1次性 一旦结果返回则连接断开
// 1. 直接自己基于TCP/UDP协议去封装一层协议 myhttp，没有通用性，http2.0
/*
	客户端：
		1. 建立连接 tcp/http
		2. 将employee序列化成JSON字符串 - 序列化
		3. 发送JSON字符串 - 调用成功后实际上你接收到的是一个二进制的数据
		4. 等待服务器发送结果
		5. 将服务器返回的数据解析成PrintResult对象 - 反序列化
	 服务端：
		1. 监听网络端口 80
		2. 读取数据 - 二进制的JSON数据
		3. 对数据进行反序列化Employee对象
		4. 开始处理业务逻辑
		5. 将处理的结果PrintResult序列化成JSON二进制数据 - 序列化
		6. 将数据返回

	序列化和反序列化是可以选择的，不一定采用JSON、XML、protobuf、msgpack

*/
//}

func main() {
	// 现在我们想把Add函数变成一个远程的函数调用，也就意味着需要把
	// Add函数放在远程服务器上去运行
	/*
			我们原本的电商系统，这里地方有个逻辑，这个逻辑是扣减库存，
			但是库存服务是一个独立的系统，reduce
			一定会牵扯到网络，做成一个WEB服务(gin、beego、net/httpserver)

			1. 这个函数的调用参数如何传递-json(一种数据格式的协议)/xml/msgpack/protobuf
			   现在网络调用有2个短 - 客户端、应该干嘛？将数据传输到gin
		       gin-服务端负责解析数据
	*/
	fmt.Println(Add(1, 2))
	// 将这个打印的工作放到另一台服务器上
	// 我需要将本地内存对象struct
	// 这样不写，可行方式将struct序列化成JSON
	fmt.Println(Employee{
		Name: "bobby",
		company: Company{
			Name:    "慕课网",
			Address: "北京市",
		},
	})
	// 远程的服务器需要将二进制对象反解成struct对象
	// 搞这么麻烦，直接全部使用JSON去格式化不香么？
	// 这种做法在浏览器和gin服务之间是可行的，但是不适用于分布式系统
}
