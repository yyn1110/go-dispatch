package kserver

import (
	"bufio"
	//"encoding/binary"
	"io"
	"kserver/pack"
	"logger"
	"net"
)

const (
	dataBufferSize = 1024
)

var dataBuffer = make([]byte, 1)

//var dataBuffer = bytes.NewBuffer(buf)

func StartServe(hostAndPort string) {
	listener := initServer(hostAndPort)
	for {
		conn, err := listener.Accept()
		logger.LogF(err, "StartServe accept err")
		go connectionHandler(conn)
	}
}

//连接
func connectionHandler(conn net.Conn) {
	connFrom := conn.RemoteAddr().String()
	logger.Log("Connect from :" + connFrom)
	//connectSuccess(conn)

	// for {
	// 	b := make([]byte, dataBufferSize)
	// 	length, err := conn.Read(b)
	// 	dataBuffer = append(dataBuffer, b)

	// 	//dataBuffer[dataBufferSize] = 0 //防止溢出
	// 	logger.LogF(err, "conn read")
	// 	switch err {

	// 	case nil:
	// 		handleData(conn, length, err, dataBuffer)
	// 	default:
	// 		logger.LogF(err, "close conn")
	// 		disconnect(conn)
	// 	}
	// }
	read := bufio.NewReader(conn)
	parsePackage(read)

}
func parsePackage(reader io.Reader) (p *pack.PackageData, err error) {
	head := make([]byte, pack.PACKAGEHEADLEN)
	_, err = io.ReadAtLeast(reader, head, int(pack.PACKAGEHEADLEN))
	logger.LogF(err, "parsePackage ")
}
func disconnect(conn net.Conn) {
	err := conn.Close()

	logger.Log("Closed  connection:" + conn.RemoteAddr().String() + "  " + err.Error())

}
func initServer(hostAndPort string) *net.TCPListener {
	serverAddr, err := net.ResolveTCPAddr("tcp", hostAndPort)
	if err != nil {

		logger.LogF(err, "ResolveTCPAddr")

	}

	listener, err := net.ListenTCP("tcp", serverAddr)
	if err != nil {
		logger.LogF(err, "ResolveTCPAddr")
	}

	//logger.LogDebug(listener.)
	logger.Log("listener host " + listener.Addr().String())
	return listener
}
