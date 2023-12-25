package System

import (
	"DBUtilities"
	"HTTPUtilities"
	"UserServer"
	"WebsocketServer"
	"time"
)

type System struct {
	Status          int
	TLS             bool
	WebServer       *HTTPUtilities.HTTPServer
	RedirectServer  *HTTPUtilities.HTTPServer
	WebsocketServer *WebsocketServer.WebsocketServer
	UserServer      *UserServer.UserServer
}

func CreateSystem(tls bool) *System {
	dbConnection := DBUtilities.InitializeConnection(DB_USERNAME, DB_PASSWORD, DB_NAME)
	//DBUtilities.DropAllTables(dbConnection)
	var webServer *HTTPUtilities.HTTPServer
	var redirectServer *HTTPUtilities.HTTPServer
	var websocketServer *WebsocketServer.WebsocketServer
	if tls {
		webServer = InitializeWebServerTLS()
		redirectServer = InitializeRedirectServerTLS()
		websocketServer = InitializeWebsocketServerTLS(dbConnection)
	} else {
		webServer = InitializeWebServer()
		redirectServer = nil
		websocketServer = InitializeWebsocketServer(dbConnection)
	}
	return &System{
		TLS:             tls,
		Status:          CLOSED,
		WebServer:       webServer,
		WebsocketServer: websocketServer,
		RedirectServer:  redirectServer,
		UserServer:      InitializeUserServer(dbConnection),
	}
}

func (system *System) Start() {

	if system.Status != CLOSED {
		println("#### servers already started ####")
		return
	}
	println("#### starting servers ####")

	go system.UserServer.Base.TcpListener.Start()
	println("userserver ✓")

	if system.TLS {
		go system.WebServer.StartHTTPS(CERT_PATH, KEY_PATH)
		println("webserver ✓")

		go system.RedirectServer.StartHTTP()
		println("redirectserver ✓")

		go system.WebsocketServer.WsListener.StartHTTPS(CERT_PATH, KEY_PATH)
	} else {
		go system.WebServer.StartHTTP()
		println("webserver ✓")

		go system.WebsocketServer.WsListener.StartHTTP()
	}

	go system.WebsocketServer.Base.TcpListener.Start()
	println("websocketserver ✓")

	system.Status = LIVE
	println("#### started servers ####")
}

func (system *System) Close() {

	if system.Status != LIVE {
		println("#### servers not started ####")
		return
	}
	println("#### closing servers ####")

	system.WebServer.Close()
	println("webserver ✗")

	if system.TLS {
		system.RedirectServer.Close()
		println("redirectserver ✗")
	}

	system.WebsocketServer.WsListener.Close()
	system.WebsocketServer.DisconnectConnections()

	for system.WebsocketServer.ConnectedWebsocketsCount() > 0 ||
		system.WebsocketServer.Base.TcpListener.OngoingRequestCount() > 0 ||
		system.UserServer.Base.TcpListener.OngoingRequestCount() > 0 {

		time.Sleep(100 * time.Millisecond)
	}

	system.WebsocketServer.Base.TcpListener.Stop()
	println("websocketserver ✗")

	system.UserServer.Base.TcpListener.Stop()
	println("userserver ✗")

	system.Status = CLOSED
	println("#### closed servers ####")
}
