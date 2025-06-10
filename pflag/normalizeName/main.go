package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/pflag"
)

func normalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	// alias
	switch name {
	case "old-flag-name":
		name = "new-flag-name"
	}

	// --my-flag == --my_flag == --my.flag
	from := []string{"-", "_"}
	to := "."
	for _, sep := range from {
		name = strings.Replace(name, sep, to, -1)
	}
	return pflag.NormalizedName(name)
}

func main() {
	flagset := pflag.NewFlagSet("test", pflag.ExitOnError)

	var ip = flagset.IntP("new-flag-name", "i", 1234, "help message for new-flag-name")
	var myFlag = flagset.IntP("my-flag", "m", 1234, "help message for my-flag")

	flagset.SetNormalizeFunc(normalizeFunc)
	flagset.Parse(os.Args[1:])

	fmt.Printf("ip: %d\n", *ip)
	fmt.Printf("myFlag: %d\n", *myFlag)
}

/*
$ go run main.go --old-flag-name 2 --my-flag 200
ip: 2
myFlag: 200

$ go run main.go --new-flag-name 3 --my_flag 300
ip: 3
myFlag: 300
*/
