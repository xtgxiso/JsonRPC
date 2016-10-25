// json_rpc_client project main.go
package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

func main() {
	//连接服务--json
    client, err := jsonrpc.Dial("tcp", "127.0.0.1:1215")
    if err != nil {
        log.Fatal("dialing:", err)
    }
    var reply int
	var num = make(map[string]int)
	num["num1"] = 3
	num["num2"] = 2
	//调用远程MyMath的Add方法,也只能是三个参数
    err = client.Call("MyMath.Add",num,&reply)
    if err != nil {
        log.Fatal("arith error:", err)
    }
	//输出结果
    fmt.Println(reply)
	var ret string
	var str = make(map[string]string)
	str["num1"] = "hello"
	str["num2"] = "world"
	//调用远程MyMath的Sub方法,也只能是三个参数
	err = client.Call("MyMath.Sub",str,&ret)
    if err != nil {
        log.Fatal("arith error:", err)
    }
	//输出结果
    fmt.Println(ret)
	client.Close()
}
