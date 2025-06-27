package main

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	goadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

// CasBinService 是声明了一个 CasBinService 结构体,包含了 enforcer 和 adapter 两个属性
type CasBinService struct {
	enforcer *casbin.Enforcer
	adapter  *goadapter.Adapter
}

// NewCasBinService 是声明了一个 NewCasBinService 函数,用于初始化 CasBinService 结构体
func NewCasBinService(db *gorm.DB) (*CasBinService, error) {
	a, err := goadapter.NewAdapterByDB(db)
	if err != nil {
		return nil, err
	}

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
		return nil, err
	}

	e, err := casbin.NewEnforcer(m, a)
	if err != nil {
		return nil, err
	}

	return &CasBinService{enforcer: e, adapter: a}, nil
}
