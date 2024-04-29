/**
 * @Author: Junhao Zhang
 * @Description:用于ABAC的属性定义
 * @File: abacmodels.go
 * @Date: 2024/4/24 上午10:15
 */

package accesscontrol

type Submodels struct {
	Age   int
	Group string
}

type AbacModels struct {
	submodels Submodels
	obj       string
	act       string
}
