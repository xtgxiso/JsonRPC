// json_rpc_server project main.go
package main

import (
	"net/rpc"
	"net"
	"log"
	"net/rpc/jsonrpc"
)

//自己的数据类
type MyMath struct{
	
}

//加法--只能两个参数
func (mm *MyMath) Add(num map[string]float64,reply *float64) error {
    *reply = num["num1"] + num["num2"]
    return nil
}

//减法--只能两个参数
func (mm *MyMath) Sub(num map[string]string,reply *string) error {
    *reply = num["num1"] + num["num2"]
    return nil
}

func main() {
	//注册MyMath类，以代客户端调用
    rpc.Register(new(MyMath))
    listener, err := net.Listen("tcp", ":1215")
    if err != nil {
        log.Fatal("listen error:", err)
    }
    for {
        conn, err := listener.Accept()
        if err != nil {
            continue
        }
		//新协程来处理--json
        go jsonrpc.ServeConn(conn)
    }
}
