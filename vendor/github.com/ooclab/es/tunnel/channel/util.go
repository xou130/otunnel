package channel

import (
	"net"

	"github.com/sirupsen/logrus"
)

func closeConn(conn net.Conn) {
	defer func() {
		if r := recover(); r != nil {
			logrus.Warn("closeConn recovered: ", r)
		}
	}()
	conn.Close()
}
