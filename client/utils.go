package main
import (
	"fmt"
	"net"
	"encoding/binary"
	"gocode/chatroom/common/message"
	"encoding/json"
)

func readPkg(conn net.Conn)(mes message.Message,err error){

	buf := make([]byte,8096)
	fmt.Println("读取客户端发送的数据。。。")
	_,err = conn.Read(buf[:4])
	if err != nil {
		// err = errors.New("read pkg header error")
		return
	}
	// fmt.Println("读到的buf=",buf[:4])

	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[:4])

	n,err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil{
		// err = errors.New("read pkg body error")
		return
	}

	err = json.Unmarshal(buf[:pkgLen],&mes)
	if err != nil{
		fmt.Println("json.Unmarshal err=",err)
		return
	}

	return
}

func writePkg(conn net.Conn,data []byte) (err error){

	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4],pkgLen)

	n,err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("发送失败err",err)
		return
	}

	n,err = conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("发送失败err",err)
		return
	}

	err = writePkg(conn,data)
	return

}