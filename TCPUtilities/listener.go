package TCPUtilities

import (
	"net"
	"sync/atomic"
	"time"
)

type Listener struct {
	Listener        *net.Listener
	handler         func(*net.Conn, []byte)
	address         string
	ongoingRequests atomic.Int32
}

func (listener *Listener) OngoingRequestCount() int32 {
	return listener.ongoingRequests.Load()
}

func CreateListener(address string, handler func(*net.Conn, []byte)) *Listener {
	return &Listener{
		Listener:        nil,
		handler:         handler,
		address:         address,
		ongoingRequests: atomic.Int32{},
	}
}

func (listener *Listener) Start() {
	tcpListener, err := net.Listen("tcp", listener.address)
	if err != nil {
		panic(err)
	}
	listener.Listener = &tcpListener
	for listener.Listener != nil {
		if tcpIncomingConnection, err := tcpListener.Accept(); err == nil {
			tcpIncomingConnection.SetDeadline(time.Now().Add(5 * time.Second))
			go func() {
				for {
					if message := ReadNextMessage(&tcpIncomingConnection); message != nil {
						listener.ongoingRequests.Add(1)
						listener.handler(&tcpIncomingConnection, message)
						listener.ongoingRequests.Add(-1)
						continue
					}
					tcpIncomingConnection.Close()
					break
				}
			}()
		}
	}
}
func (listener *Listener) Stop() {
	(*listener.Listener).Close()
	listener.Listener = nil
}
