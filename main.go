/**
 * @Author: Junhao Zhang
 * @Description:
 * @File: accesscontrol
 * @Version: 1.0.0
 * @Date: 2024/4/19 ??10:05
 */
package main

import (
	"SafeBoxies/accesscontrol"
	"SafeBoxies/util"
	"fmt"
	"log"
)

/**
* @Author: Junhao Zhang
* @Description:main 程序入口
* @DateTime: 2024/4/19 ??4:41
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

	//--------------features chanage------------------//

	// 初始化 Casbin 权限控制模型和策略。
	err = accesscontrol.InitCasbin()
	// 如果初始化失败，打印错误信息并使用 log.Fatalf 停止程序。
	if err != nil {
		log.Fatalf("Failed to initialize Casbin: %v", err)
	}

	// 定义一个用户结构体，包含用户年龄和组。
	sub := accesscontrol.Submodels{Age: 19, Group: "admin"}

	// 检查定义的用户是否有权限对 "document" 执行 "write" 操作。
	ok, err := accesscontrol.CheckPermission(sub, "document", "read")
	// 如果检查过程中出现错误，打印错误信息并停止程序。
	if err != nil {
		log.Fatalf("Failed to enforce policy: %v", err)
	}

	// 如果用户有权限，则打印允许访问的信息。
	if ok {
		fmt.Println("允许访问")
	} else {
		// 如果用户没有权限，打印拒绝访问的信息。
		fmt.Println("ac deny")
	}
}
