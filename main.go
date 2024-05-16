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
	"github.com/ChiamZhang/SafeTransBoxies/util"
	"os"
)

/**
* @Author: Junhao Zhang
* @Description:main 程序入口
* @DateTime: 2024/4/19  4:41
 */

func RootCheck() bool {
	//------------------演示root检查模块--------------------------//

	// 调用 IsRoot 函数来检查当前用户是否是 root 用户。
	isRoot, err := util.IsRoot()
	// 如果在检查过程中发生错误，打印错误信息并返回，停止程序。
	if err != nil {
		fmt.Printf("运行错误: %s\n", err)
		return false
	}
	// 如果当前用户是 root 用户，则打印禁止信息并停止程序。
	if isRoot {
		fmt.Println("禁止 root 用户运行，正在停止.")
		return false
	} else {
		// 如果当前用户不是 root 用户，则打印确认信息。
		fmt.Println("非 root 用户执行.")
		return true

	}

}
func Encrypt() {
	//------------------加解密演示--------------------------//
	file, err := os.Open("./dataandpolicy/data/model.conf")
	if err != nil {
		fmt.Println("打开文件时出错：", err)
		return
	}
	defer file.Close()

	// 获取文件信息，以确定文件大小
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("获取文件信息时出错：", err)
		return
	}

	// 根据文件大小动态分配读取数据的缓冲区大小
	fileSize := fileInfo.Size()
	data := make([]byte, fileSize)

	// 读取文件内容并显示在控制台
	count, err := file.Read(data)
	if err != nil {
		fmt.Println("读取文件时出错：", err)
		return
	}
	fmt.Println("文件内容：", string(data[:count]))

	middlewares := map[string][]byte{
		"model.conf": []byte("这是中间件1的内容"),
		"policy.csv": []byte("这是中间件2的内容"),
		"dataTable":  []byte("这是中间件2的内容"),
	}

	// AES密钥，注意密钥长度必须是16、24或32字节
	key := []byte("0123456789abcdef0123456789abcdef")

	// ZIP 文件密码
	zipPassword := "123456"

	// 输出加密的zip文件路径
	outputPath := "./dataandpolicy/data/encrypted_middlewares.zip"

	// 创建加密并带密码保护的zip文件
	err = util.CreateEncryptedZipFile(middlewares, key, zipPassword, outputPath)
	if err != nil {
		panic("错误: " + err.Error())
	}

	fmt.Println("加密ZIP文件创建成功")

}
func permit() {
	//--------------授权演示------------------//

	//requesterSlice := []accesscontrol2.AccessRequester{}
	//
	//casbinAbacRequester := accesscontrol2.CasbinAbacRequester{
	//	AbacModels: accesscontrol2.AbacModels{
	//		SubModel: accesscontrol2.SubModels{
	//			Age:   18,
	//			Group: "admin",
	//		},
	//		Obj: "document",
	//		Act: "write",
	//	},
	//}
	//err = casbinAbacRequester.InitChecker()
	//if err != nil {
	//	print(err)
	//}
	//
	//requesterSlice = append(requesterSlice, &casbinAbacRequester)
	//
	//ok, err := accesscontrol2.GetPermission(requesterSlice)
	//if err != nil {
	//	print((err))
	//} else if ok {
	//	print("授权成功")
	//} else {
	//	print("授权失败")
	//}
}

// 定义 main 函数，程序的入口点。
func main() {
	if !RootCheck() {
		return
	}
	Encrypt()

}
