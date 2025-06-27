package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"
)

// CheckPermission 用于检查鉴权结果
func CheckPermission(e *casbin.Enforcer, sub, obj, act string) {
	ok, err := e.Enforce(sub, obj, act)
	if err != nil {
		fmt.Printf("检查失败: %s\n", err)
	}
	if ok {
		fmt.Printf("用户: %s 访问资源: %s 使用方法: %s 检查通过\n", sub, obj, act)
	} else {
		fmt.Printf("用户: %s 访问资源: %s 使用方法: %s 检查拒绝\n", sub, obj, act)
	}
}

func main() {
	// 初始化一个 casBin 实例
	e, err := casbin.NewEnforcer("./model.conf", "./policy.csv")
	if err != nil {
		fmt.Printf("初始化 casbin 实例失败: %s\n", err)
	}

	// 检查权限（预计输出正确）
	CheckPermission(e, "zhangsan", "/api/user", "GET")

	// 检查权限（预计输出错误）
	CheckPermission(e, "lisi", "/api/user", "GET")
}
