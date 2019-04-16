package main
import (
	"fmt"
	"net"
	"gocode/chatroom/common/message"
	"encoding/json"
	"encoding/binary"
)

func login(userId int,userPwd string) (err error){


	// fmt.Printf("userId = %d userPwd = %s\n",userId,userPwd)

	// return nil

	conn,err := net.Dial("tcp","127.0.0.1:8889")
	if err != nil {
		fmt.Println("net.Dial err=",err)
		return
	}

	defer conn.Close()

	var mes message.Message
	mes.Type = message.LoginMesType

	var loginMes message.LoginMes

	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	data,err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err=",err)
		return
	}
	mes.Data = string(data)

	data,err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=",err)
		return
	}

	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4],pkgLen)

	n,err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("发送失败err",err)
		return
	}

	// fmt.Printf("客户端，发送消息长度=%d 内容是=%s",len(data),string(data))

	_,err = conn.Write(data)
	if err != nil{
		fmt.Println("conn.Write(bytes) fail",err)
		return
	}

	// time.Sleep(20 * time.Second)
	// fmt.Println("休眠了20s...")
	mes,err = readPkg(conn)

	if err != nil {
		fmt.Println("readPkg(conn) err=",err)
		return
	}

	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data),&loginResMes)
	if loginResMes.Code == 200 {
		fmt.Println("登录成功")
	}else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}

	return
}