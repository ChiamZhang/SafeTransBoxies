/**
 * @Author: Junhao Zhang
 * @Description: 用于实现所有的访问控制的接口
 * @File: accessrequester
 * @Date: 2024/4/24 下午3:49
 */

package accesscontrol

import "github.com/casbin/casbin/v2"

// AccessRequester
// @Description: 所有访问控制的实现的接口，实现方式按照CasbinAbac的方式实现即可完成功能的添加
type AccessRequester interface {
	CheckAccessPermission() (bool, error)
}

// CasbinAbacRequester
// @Description: 用于实现AccessRequester接口的CasbinAbacRequester类型
type CasbinAbacRequester struct {

	// CasbinModelPath 是模型配置文件的路径
	CasbinModelPath string

	// CasbinPolicyPath 是策略配置文件的路径
	CasbinPolicyPath string

	// Enforcer 是一个全局的 Casbin 执行者实例
	Enforcer *casbin.Enforcer

	AbacModels AbacModels
}

// InitCasbin 初始化 Casbin 执行者
func (c *CasbinAbacRequester) InitCasbin() error {
	var err error
	// CasbinModelPath 是模型配置文件的路径
	c.CasbinModelPath = "./accesscontrol/config/model.conf"
	// CasbinPolicyPath 是策略配置文件的路径
	c.CasbinPolicyPath = "./accesscontrol/config/policy.csv"
	c.Enforcer, err = casbin.NewEnforcer(c.CasbinModelPath, c.CasbinPolicyPath)
	return err
}

// CheckPermission 用于检查是否允许访问
func (c *CasbinAbacRequester) CheckAccessPermission() (bool, error) {
	AbacModels := c.AbacModels
	ok, err := c.Enforcer.Enforce(AbacModels.submodels, AbacModels.obj, AbacModels.act)
	return ok, err

}

// CasbinAbacRequester End
