package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn){
	
	defer conn.Close()
	
	processor := &Processor{
		Conn: conn,
	}
	err := processor.Process2()
	
	if err != nil{
		fmt.Println("客户端和服务端通讯协程错误err=",err)
		return
	}
}

func main(){
	fmt.Println("服务器在8889端口监听...~~")

	listen,err := net.Listen("tcp","0.0.0.0:8889")
	defer listen.Close()
	if err != nil{
		fmt.Println("net.Listen err=",err)
		return
	}

	for {
		fmt.Println("等待客户端来连接服务器...")
		conn,err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err = ",err)
		}

		//一旦连接成功，则启动一个协程和客户端保持通讯
		go process(conn)
	}
}