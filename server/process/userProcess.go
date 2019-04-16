package process2
import (
	"fmt"
	"net"
	"gocode/chatroom/common/message"
	"gocode/chatroom/server/utils"
	"encoding/json"
)

type UserProcess struct{
	Conn net.Conn
}

func (this *UserProcess) SreverProcessLogin(mes *message.Message) (err error){

	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data),&loginMes)
	if err != nil{
		fmt.Println("json.Unmarshal fail err=",err)
		return
	}

	var resMes message.Message
	resMes.Type = message.LoginMesType

	var loginResMes message.LoginResMes

	if loginMes.UserId == 100 && loginMes.UserPwd == "abc" {

		loginResMes.Code = 200
	}else{
		
		loginResMes.Code = 500
		loginResMes.Error = "该用户不存在，请注册在使用..."
	}

	data,err := json.Marshal(loginResMes)
	if err != nil{
		fmt.Println("json.Marshal fail err=",err)
		return
	}

	resMes.Data = string(data)

	data,err = json.Marshal(resMes)
	if err != nil{
		fmt.Println("json.marshal fail err=",err)
		return
	}
	
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	// err = writePkg(conn,data)
	err = tf.WritePkg(data)
	return
}
