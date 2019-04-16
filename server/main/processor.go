package main
import (
	"fmt"
	"net"
	"gocode/chatroom/common/message"
	"gocode/chatroom/server/process"
	"gocode/chatroom/server/utils"
	"io"
)

type Processor struct{
	Conn net.Conn
}

func (this *Processor) ServerProcessMes(mes *message.Message) (err error){

	switch mes.Type {
		case message.LoginMesType :
			//处理登录
			up := &process2.UserProcess{
				Conn: this.Conn,
			}
			err = up.SreverProcessLogin(mes)
		case message.RegisterMesType :
			//处理注册
		default :
			fmt.Println("消息类型不存在，无法处理...")
	}
	return
}

func (this *Processor) Process2() (err error){

	for {
		
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes,err := tf.ReadPkg()
		if err != nil{
			if err == io.EOF{
				fmt.Println("客户端退出，服务器端也退出。。")
				return nil
			}else{
				fmt.Println("readPkg err=",err)
				return	err
			}
			
		}

		err = this.ServerProcessMes(&mes)
		if err != nil{
			return err
		}
	}
}