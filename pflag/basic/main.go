package main

import (
	"fmt"

	"github.com/spf13/pflag"
)

// 实现了 pflag.Value 接口，可以作为 pflag.Var 的参数
type host struct {
	value string
}

func (h *host) String() string {
	return h.value
}

func (h *host) Set(v string) error {
	h.value = v
	return nil
}

func (h *host) Type() string {
	return "host"
}

func main() {
	// pflag.<Type> 类方法名会将标志参数值存储在指针中并返回
	var ip *int = pflag.Int("ip", 1234, "help message for ip")

	var port int
	// pflag.<Type>Var 类方法名中包含 Var 关键字的，会将标志参数值绑定到第一个指针类型的参数。
	pflag.IntVar(&port, "port", 8080, "help message for port")

	var h host
	pflag.Var(&h, "host", "help message for host")

	// 解析命令行参数
	pflag.Parse()

	fmt.Printf("ip: %d\n", *ip)
	fmt.Printf("port: %d\n", port)
	fmt.Printf("host: %+v\n", h)

	fmt.Printf("NFlag: %v\n", pflag.NFlag()) // 返回已设置的命令行标志个数
	fmt.Printf("NArg: %v\n", pflag.NArg())   // 返回处理完标志后剩余的参数个数
	fmt.Printf("Args: %v\n", pflag.Args())   // 返回处理完标志后剩余的参数列表
	fmt.Printf("Arg(1): %v\n", pflag.Arg(1)) // 返回处理完标志后剩余的参数列表中第 i 项
}

/*
go run main.go --ip 1 x y --host baidu a b
ip: 1
port: 8080
host: {value:baidu}
NFlag: 2
NArg: 4
Args: [x y a b]
Arg(1): y
*/
