package main

import (
	proto2 "OldPackageTest/helloworld/proto"
	"fmt"
	"github.com/golang/protobuf/proto"
)

type Hello struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Courses []string `json:"courses"`
}

func main() {
	req := proto2.HelloRequest{
		Name:    "bobby",
		Age:     18,
		Courses: []string{"go", "gin", "微服务"},
	}
	//jsonStruct := Hello{Name: "bobby", Age: 18, Courses: []string{"go", "gin", "微服务"}}
	//jsonRsp, _ := json.Marshal(jsonStruct)
	//fmt.Println(string(jsonRsp))
	//fmt.Println(len(jsonRsp))
	// 具体的编码是如何做到的，可以搜索一下
	// protobuf原理，变长的int
	rsp, _ := proto.Marshal(&req)
	newReq := proto2.HelloRequest{}
	//fmt.Println(string(rsp))
	//fmt.Println(len(rsp))
	_ = proto.Unmarshal(rsp, &newReq)
	fmt.Println(newReq.Name, newReq.Age, newReq.Courses)
}
