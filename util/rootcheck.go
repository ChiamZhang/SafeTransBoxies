/**
 * @Author: Junhao Zhang
 * @Description:
 * @File: rootcheck.go
 * @Version: 1.0.0
 * @Date: 2024/4/19 上午9:30
 */

package util

import (
	"errors"
	"os/user"
)

/**
* @Author: Junhao Zhang
* @Description:
* @Params:
* @Return: bool,error
* @DateTime: 2024/4/19 下午4:37
 */

func IsRoot() (bool, error) {

	currentUser, err := user.Current()
	if err != nil {
		return false, errors.New("无法获取当前用户信息")
	}
	return currentUser.Uid == "0", nil
}
