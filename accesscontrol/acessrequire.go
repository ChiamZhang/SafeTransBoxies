/**
 * @Author: Junhao Zhang
 * @Description: 使用的api,传入一个AccessRequester类型的切片，里面是基于AccessRequester接口实现的对象。
 *               每个数据对象都实现了一个函数CheckAccessPermission，当所有的切片的每一个检测都通过后才同意授权。
 * @File: acessrequir
 * @Date: 2024/4/24 上午9:36
 */

package accesscontrol

func GetPermission(requesterSlice []AccessRequester) (bool, error) {

	for _, r_e := range requesterSlice {
		ok, err := r_e.CheckAccessPermission()
		if err != nil {
			return false, err
		}
		if ok == false {
			return false, nil
		}
	}
	return true, nil
}
