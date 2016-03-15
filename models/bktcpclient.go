package models

import (
	"net"
)

type BkTcpClient struct {
	tcpAddr   *net.TCPAddr
	conn      *net.TCPConn
	keepAlive bool
	isClosed  bool
}

func NewBkTcpClient(proto string, ipport string, keepAlive bool) *BkTcpClient {
	tcpAddr, _ := net.ResolveTCPAddr(proto, ipport)
	return &BkTcpClient{tcpAddr: tcpAddr, conn: nil, keepAlive: keepAlive, isClosed: true}
}

//新建一个长连接
func (a *BkTcpClient) Connect() (int, error) {
	var err error
	a.conn, err = net.DialTCP("tcp", nil, a.tcpAddr)
	if err != nil {
		a.isClosed = true
		return 1, err
	}
	a.isClosed = false
	return 0, nil
}

//关闭一个长连接
func (a *BkTcpClient) Close() {
	if a.conn != nil {
		a.conn.Close()
		a.conn = nil
		a.isClosed = true
	}
	return
}

//发送报文
func (a *BkTcpClient) Send(d *BkNotifyPacket) (*BkNotifyPacket, error) {
	var err error
	if a.keepAlive != true {
		//短连接需要每次都连接
		//若已经打开连接则先关闭
		if a.isClosed == false {
			a.Close()
		}

		//打开连接
		_, err = a.Connect()
		if err != nil {
			return nil, err
		}

		//保证关闭
		defer a.Close()

		bkProtocol := &BkNotifyProtocol{}
		a.conn.Write(d.Serialize())

		p, err := bkProtocol.ReadPacket(a.conn)
		if err == nil {
			rsp := p.(*BkNotifyPacket)
			return rsp, nil
		}
	} else {

		//若没有调用过连接，则需要重新连接
		if a.isClosed == true {
			_, err = a.Connect()
			if err != nil {
				return nil, err
			}
		}

		bkProtocol := &BkNotifyProtocol{}
		a.conn.Write(d.Serialize())

		p, err := bkProtocol.ReadPacket(a.conn)
		if err == nil {
			rsp := p.(*BkNotifyPacket)
			return rsp, nil
		}
	}
	return nil, err
}
