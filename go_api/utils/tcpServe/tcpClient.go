package tcpServe

import (
	"fmt"
	"go_api/app/config"
	"net"
)

var TcpConn net.Conn

func TcpClient() (err error) {
	TcpConn, err = net.Dial("tcp", config.ApplicationConfig.TerminalTcp)
	if err != nil {
		fmt.Println("err :", err)
	}
	return
}

func TcpDate(data []byte) (err error) {
	if TcpConn != nil {
		_, err = TcpConn.Write(data) // 发送数据
		if err != nil {
			err = TcpClient()
			if err != nil {
				return
			}
			_, err = TcpConn.Write(data) // 发送数据
		}
		_, err = TcpConn.Write(data) // 发送数据
	} else {
		err = TcpClient()
		if err != nil {
			return
		}
		_, err = TcpConn.Write(data) // 发送数据
	}
	return
}
