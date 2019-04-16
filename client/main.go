package main
import (
	"fmt"
	"os"
)

var userId int
var userPwd string

func main(){

var key int
var loop = true

	for {
		fmt.Println("---------------欢迎登陆多人聊天系统------------")
		fmt.Println("\t\t 1 登陆聊天室")
		fmt.Println("\t\t 2 注册用户")
		fmt.Println("\t\t 3 退出系统")
		fmt.Println("\t\t 请选择（1-3）：")

		fmt.Scanf("%d\n",&key)
		switch key {
		case 1 :
			fmt.Println("登陆聊天室")
			loop = false
		case 2 :
			fmt.Println("注册用户")
			loop = false
		case 3 :
			fmt.Println("退出系统")
			// loop = false
			os.Exit(0)
		default :
			fmt.Println("你的输入有误，请重新输入")
		}

		if !loop{
			break
		} 
	}

	if key == 1{

		fmt.Println("请输入用户的id:")
		fmt.Scanf("%d\n",&userId)
		fmt.Println("请输入用户的密码:")
		fmt.Scanf("%s\n",&userPwd)
		
		login(userId,userPwd)
		// if err != nil {
		// 	fmt.Println("登录失败!")
		// }else{
		// 	fmt.Println("登录成功!")
		// }
	}else{
		fmt.Println("进行用户注册的逻辑。。。")
	}
}