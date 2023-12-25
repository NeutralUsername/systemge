package System

import (
	"HTTPUtilities"
	"Ports"
	"UserServer"
	"Utilities"
	"WebsocketServer"
	"database/sql"
	"net/http"
)

func InitializeWebServer() *HTTPUtilities.HTTPServer {
	webserver := HTTPUtilities.CreateHttpServer(Ports.HTTP_LISTENER, http.NewServeMux())
	webserver.SetHandlerFunc("/", HTTPUtilities.SendDirectory(FRONTEND_PATH))

	return webserver
}
func InitializeWebServerTLS() *HTTPUtilities.HTTPServer {
	webserver := HTTPUtilities.CreateHttpServer(Ports.HTTPS_LISTENER, http.NewServeMux())
	webserver.SetHandlerFunc("/", HTTPUtilities.SendDirectory(FRONTEND_PATH))

	return webserver
}

func InitializeRedirectServerTLS() *HTTPUtilities.HTTPServer {
	redirectServer := HTTPUtilities.CreateHttpServer(Ports.HTTP_LISTENER, http.NewServeMux())
	redirectServer.SetHandlerFunc("/", HTTPUtilities.Redirect(HTTPS_URL))

	return redirectServer
}

func InitializeWebsocketServer(dbConnection *sql.DB) *WebsocketServer.WebsocketServer {
	websocketServer := WebsocketServer.CreateWebsocketServer()
	websocketServer.Base.SetTcpListener(Ports.WEBSOCKETSERVER_TCP_LISTENER, websocketServer.HandleTcpMessage)
	websocketServer.Base.Random = Utilities.CreateRandom()
	websocketServer.Base.Logger = Utilities.CreateLogger(WEBSOCKET_SERVER_ERROR_LOG_PATH)
	websocketServer.Base.DbConnection = dbConnection
	websocketServer.CreateConnectionsTable()
	websocketServer.CreateDisconnectsTable()
	websocketServer.WsListener = HTTPUtilities.CreateHttpServer(Ports.WEBSOCKET_LISTENER, http.NewServeMux())
	websocketServer.WsListener.SetHandlerFunc("/ws", websocketServer.HandleWebsocketConnection())

	return websocketServer
}

func InitializeWebsocketServerTLS(dbConnection *sql.DB) *WebsocketServer.WebsocketServer {
	websocketServer := WebsocketServer.CreateWebsocketServer()
	websocketServer.Base.SetTcpListener(Ports.WEBSOCKETSERVER_TCP_LISTENER, websocketServer.HandleTcpMessage)
	websocketServer.Base.Random = Utilities.CreateRandom()
	websocketServer.Base.Logger = Utilities.CreateLogger(WEBSOCKET_SERVER_ERROR_LOG_PATH)
	websocketServer.Base.DbConnection = dbConnection
	websocketServer.CreateConnectionsTable()
	websocketServer.CreateDisconnectsTable()
	websocketServer.WsListener = HTTPUtilities.CreateHttpServer(Ports.WEBSOCKET_LISTENER, http.NewServeMux())
	websocketServer.WsListener.SetHandlerFunc("/ws", websocketServer.HandleWebsocketConnection())

	return websocketServer
}

func InitializeUserServer(dbConnection *sql.DB) *UserServer.UserServer {
	userServer := UserServer.CreateUserServer()
	userServer.Base.SetTcpListener(Ports.USERSERVER_TCP_LISTENER, userServer.HandleTcpMessage)
	userServer.Base.Random = Utilities.CreateRandom()
	userServer.Base.Logger = Utilities.CreateLogger(USER_SERVER_ERROR_LOG_PATH)
	userServer.Base.DbConnection = dbConnection
	userServer.CreateUsersTable()

	return userServer
}
