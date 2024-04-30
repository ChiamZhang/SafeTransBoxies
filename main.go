/**
 * @Author: Junhao Zhang
 * @Description:
 * @File: accesscontrol
 * @Version: 1.0.0
 * @Date: 2024/4/19 10:05
 */
package main

import (
	"fmt"
	"github.com/ChiamZhang/SafeTransBoxies/accesscontrol"
	"github.com/ChiamZhang/SafeTransBoxies/util"
)

/**
* @Author: Junhao Zhang
* @Description:main 程序入口
* @DateTime: 2024/4/19  4:41
 */

// 定义 main 函数，程序的入口点。
func main() {

	// 调用 IsRoot 函数来检查当前用户是否是 root 用户。
	isRoot, err := util.IsRoot()
	// 如果在检查过程中发生错误，打印错误信息并返回，停止程序。
	if err != nil {
		fmt.Printf("运行错误: %s\n", err)
		return
	}
	// 如果当前用户是 root 用户，则打印禁止信息并停止程序。
	if isRoot {
		fmt.Println("禁止 root 用户运行，正在停止.")
		return
	} else {
		// 如果当前用户不是 root 用户，则打印确认信息。
		fmt.Println("非 root 用户执行.")
	}

	//--------------授权演示------------------//

	requesterSlice := []accesscontrol.AccessRequester{}

	casbinAbacRequester := accesscontrol.CasbinAbacRequester{
		AbacModels: accesscontrol.AbacModels{
			SubModel: accesscontrol.SubModels{
				Age:   19,
				Group: "admin",
			},
			Obj: "document",
			Act: "write",
		},
	}
	err = casbinAbacRequester.InitChecker()
	if err != nil {
		print(err)
	}
	requesterSlice = append(requesterSlice, &casbinAbacRequester)

	ok, err := accesscontrol.GetPermission(requesterSlice)
	if err != nil {
		print((err))
	} else if ok {
		print("授权成功")
	} else {
		print("授权失败")
	}

}
