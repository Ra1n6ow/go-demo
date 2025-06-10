package main

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

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
	// 创建一个 FlagSet 实例，用于管理命令行标志
	// 相较于 pflag.CommandLine 这个默认的全局 FlagSet 实例，NewFlagSet 创建的 FlagSet 实例是独立的，不会影响 pflag.CommandLine
	flagset := pflag.NewFlagSet("test", pflag.ExitOnError)

	// pflag.<Type>P、pflag.<Type>VarP 类方法名以 P 结尾的，支持简短标志。
	var ip = flagset.IntP("ip", "i", 1234, "help message for ip")

	var boolVar bool
	flagset.BoolVarP(&boolVar, "boolVar", "b", true, "help message for boolVar")

	var h host
	flagset.VarP(&h, "host", "H", "help message for host")

	// 禁止打印帮助信息时对标志进行重排序
	flagset.SortFlags = false

	flagset.Parse(os.Args[1:])

	fmt.Printf("ip: %d\n", *ip)
	// 布尔类型的标志指定参数 --boolVar=false 需要使用等号 = 而非空格
	fmt.Printf("boolVar: %t\n", boolVar)
	fmt.Printf("host: %+v\n", h)

	i, err := flagset.GetInt("ip")
	fmt.Printf("i: %d, err: %v\n", i, err)
}
