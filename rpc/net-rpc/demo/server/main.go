package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Args struct {
	X, Y int
}

// ServiceA 自定义一个结构体类型
type ServiceA struct{}

// Add 为ServiceA类型增加一个可导出的Add方法
func (s *ServiceA) Add(args *Args, reply *int) error {
	*reply = args.X + args.Y
	return nil
}

func main() {
	/*
		// 基于http协议
		service := new(ServiceA)
		rpc.Register(service) // 注册RPC服务
		rpc.HandleHTTP()      // 基于HTTP协议
		l, e := net.Listen("tcp", ":9091")
		if e != nil {
			log.Fatal("listen error:", e)
		}
		http.Serve(l, nil)
	*/

	// 基于tcp协议
	service := new(ServiceA)
	rpc.Register(service) // 注册RPC服务
	l, e := net.Listen("tcp", ":9091")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	for {
		conn, _ := l.Accept()
		// 基于 gob 编码
		// rpc.ServeConn(conn)
		// 基于 json 编码
		rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
