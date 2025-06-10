package main

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

func main() {
	flags := pflag.NewFlagSet("test", pflag.ExitOnError)

	var ip = flags.IntP("ip", "i", 1234, "help message for ip")

	var boolVar bool
	flags.BoolVarP(&boolVar, "boolVar", "b", true, "help message for boolVar")

	var h string
	flags.StringVarP(&h, "host", "H", "127.0.0.1", "help message for host")

	// 弃用标志
	flags.MarkDeprecated("ip", "deprecated")
	flags.MarkShorthandDeprecated("boolVar", "please use --boolVar only")

	// 隐藏标志
	flags.MarkHidden("host")

	flags.Parse(os.Args[1:])

	fmt.Printf("ip: %d\n", *ip)
	fmt.Printf("boolVar: %t\n", boolVar)
	fmt.Printf("host: %+v\n", h)
}
