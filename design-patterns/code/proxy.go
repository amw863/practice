/*
 * @Author : wb
 * @Date : 2022/2/3 4:14 下午
 */

package code

type User interface {
	login()
}

type UserController struct {
}

func (u UserController) login() {

}

// 基于接口而非基于实现的思想或者继承
type UserControllerProxy struct {
	u UserController
}

func (u UserControllerProxy) login() {
	// todo 额外事情
	u.u.login()
	// todo 额外事情
}
