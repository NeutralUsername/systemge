package WebsocketServer

import (
	"MessageTypes"
	"TCPUtilities"
	"net/http"
)

func (server *WebsocketServer) HandleWebsocketConnection() func(http.ResponseWriter, *http.Request) {
	return func(httpResponseWriter http.ResponseWriter, httpRequest *http.Request) {
		if wsConnection, err := upgrader.Upgrade(httpResponseWriter, httpRequest, nil); err == nil {
			username, password := processLoginCookies(httpRequest)
			userId := 0
			if userId = server.AuthenticateCredentials(username, password); userId == -1 {
				if newGuestUserId := server.NewGuest(); newGuestUserId != -1 {
					userId = newGuestUserId
				} else {
					server.Logger.Log("Error: Could not create new guest user.")
					wsConnection.Close()
					return
				}
			}
			connection := server.addConnection(userId, wsConnection)
			connection.message(TCPUtilities.ConstructMessage(MessageTypes.LOGIN, server.assembleLoginData(connection.user.Id), nil))
			for {
				if _, message, err := wsConnection.ReadMessage(); err == nil {
					connection.watchdog.Reset(WEBSOCKET_WATCHDOG_TIMEOUT)
					server.handleWebsocketMessage(message, connection)
					continue
				}
				break
			}
		} else {
			server.Logger.Log("Error: Could not upgrade connection to websocket.")
		}
	}
}

func processLoginCookies(httpRequest *http.Request) (username string, password string) {
	usernameCookie, _ := httpRequest.Cookie(COOKIE_KEY_USERNAME)
	if usernameCookie != nil {
		username = usernameCookie.Value
	}
	passwordCookie, _ := httpRequest.Cookie(COOKIE_KEY_PASSWORD)
	if passwordCookie != nil {
		password = passwordCookie.Value
	}
	return username, password
}
