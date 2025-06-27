package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

func initCassbin() *casbin.Enforcer {
	// 声明一个 gorm db 实例
	// db, err := gorm.Open(sqlite.Open("demo.db"), &gorm.Config{})
	// if err != nil {
	// fmt.Printf("创建数据库连接失败: %v\n", err)
	// }

	// 声明一个 adapter 实例, 传入 db 实例
	// a, err := gormadapter.NewAdapterByDB(db)
	a, err := gormadapter.NewAdapter("mysql", "root:123456@tcp(127.0.0.1:3306)/casbin_testdb", true)
	if err != nil {
		fmt.Printf("创建数据库或 adapter 失败: %v\n", err)
	}

	// 声明一个 model 用于存储 casBin model 策略
	m, err := model.NewModelFromString(`[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && keyMatch2(r.obj,p.obj) && r.act == p.act`)
	if err != nil {
		fmt.Printf("创建 model 失败: %v\n", err)
	}

	// 声明一个 enforcer 实例, 传入 model 和 adapter 实例
	e, err := casbin.NewEnforcer(m, a)
	if err != nil {
		fmt.Printf("创建 enforcer 失败: %v\n", err)
	}

	// 加载策略
	err = e.LoadPolicy()
	if err != nil {
		fmt.Printf("加载策略失败: %v\n", err)
	}

	// 添加策略, 允许 admin 组通过 GET 方法访问 /api/user 路由
	_, _ = e.AddPolicy("admin", "/api/user", "GET")

	// 添加组策略绑定用户到组
	_, _ = e.AddGroupingPolicy("layzer", "admin")

	return e
}

func checkPermission(e *casbin.Enforcer, sub, obj, act string) {
	ok, err := e.Enforce(sub, obj, act)
	if err != nil {
		fmt.Printf("验证权限出错: %v\n", err)
	}
	if ok {
		fmt.Printf("用户 %s 有权限访问 %s 路由 %s 方法\n", sub, obj, act)
	} else {
		fmt.Printf("用户 %s 没有权限访问 %s 路由 %s 方法\n", sub, obj, act)
	}
}

func main() {
	e := initCassbin()

	checkPermission(e, "layzer", "/api/user", "GET")
	checkPermission(e, "layzer", "/api/user", "POST")
}
