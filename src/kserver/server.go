package kserver

import (
	"logger"
	"net"
)

func initServer(hostAndPort string) *net.TCPListener {
	serverAddr, err := net.ResolveTCPAddr("tcp", hostAndPort)
	if err != nil {
		logger.LogError(err, "ResolveTCPAddr")

	}

	listener, err := net.ListenTCP("tcp", serverAddr)
	if err != nil {
		logger.LogError(err, "ResolveTCPAddr")
	}

	//logger.LogDebug(listener.)
	logger.LogNormal("listener host " + listener.Addr().String())
	return listener
}
