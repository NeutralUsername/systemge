package WebsocketServer

import (
	"MessageTypes"
	"TCPUtilities"
	"Utilities"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  TCPUtilities.BUFFER_SIZE,
	WriteBufferSize: TCPUtilities.BUFFER_SIZE,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (server *WebsocketServer) handleWebsocketMessage(message []byte, connection *Connection) {
	if messageType, messageData1, messageData2 := TCPUtilities.ParseMessage(message); messageType != "" && len(messageData2) == 0 {
		switch messageType {
		case MessageTypes.LOGIN:
			server.handleLogin(connection, messageData1, messageData2)
		case MessageTypes.REGISTER:
			server.handleRegister(connection, messageData1, messageData2)

		case MessageTypes.UITHEME:
			server.handleUiTheme(connection, messageData1, messageData2)
		case MessageTypes.SHOWCOMMUNICATOR:
			server.handleShowCommunicator(connection, messageData1, messageData2)
		case MessageTypes.COMMUNICATORCONTENT:
			server.handleCommunicatorContent(connection, messageData1, messageData2)
		case MessageTypes.SHOWTUTORIAL:
			server.handleShowTutorial(connection, messageData1, messageData2)

		default:
			server.handleUnknownMessage(connection, messageType)
		}
	}
}

func (server *WebsocketServer) handleLogin(connection *Connection, messageData1, messageData2 []string) {
	if len(messageData1) == 2 { //username, password
		if userId := server.AuthenticateCredentials(messageData1[0], messageData1[1]); userId != -1 {
			server.changeUser(connection, userId)
			connection.message(TCPUtilities.ConstructMessage(MessageTypes.LOGIN, server.assembleLoginData(connection.user.Id), nil))
			return
		}
	}
	connection.message(TCPUtilities.ConstructMessage(MessageTypes.ERROR, []string{MessageTypes.LOGIN}, nil))
}

func (server *WebsocketServer) handleRegister(connection *Connection, messageData1, messageData2 []string) {
	if len(messageData1) == 5 { //username, password, email, secretQuestion, secretAnswer
		if server.Register(connection.user.Id, messageData1[0], messageData1[1], messageData1[2], messageData1[3], messageData1[4]) {
			server.messageUserByConnectionId(connection.id, TCPUtilities.ConstructMessage(MessageTypes.LOGIN, server.assembleLoginData(connection.user.Id), nil))
			return
		}
	}
	connection.message(TCPUtilities.ConstructMessage(MessageTypes.ERROR, []string{MessageTypes.REGISTER}, nil))
}

func (server *WebsocketServer) handleUiTheme(connection *Connection, messageData1, messageData2 []string) {
	if len(messageData1) == 1 { //uiTheme
		if server.UiTheme(connection.user.Id, messageData1[0]) {
			server.messageUserByConnectionId(connection.id, TCPUtilities.ConstructMessage(MessageTypes.UITHEME, messageData1, nil))
			return
		}
	}
	connection.message(TCPUtilities.ConstructMessage(MessageTypes.ERROR, []string{MessageTypes.UITHEME}, nil))
}

func (server *WebsocketServer) handleShowCommunicator(connection *Connection, messageData1, messageData2 []string) {
	if len(messageData1) == 1 { //bool
		if server.ShowCommunicator(connection.user.Id, Utilities.StringToBool(messageData1[0])) {
			server.messageUserByConnectionId(connection.id, TCPUtilities.ConstructMessage(MessageTypes.SHOWCOMMUNICATOR, messageData1, nil))
			return
		}
	}
	connection.message(TCPUtilities.ConstructMessage(MessageTypes.ERROR, []string{MessageTypes.SHOWCOMMUNICATOR}, nil))
}

func (server *WebsocketServer) handleCommunicatorContent(connection *Connection, messageData1, messageData2 []string) {
	if len(messageData1) == 1 { //bool
		if server.CommunicatorContent(connection.user.Id, messageData1[0]) {
			server.messageUserByConnectionId(connection.id, TCPUtilities.ConstructMessage(MessageTypes.COMMUNICATORCONTENT, messageData1, nil))
			return
		}
	}
	connection.message(TCPUtilities.ConstructMessage(MessageTypes.ERROR, []string{MessageTypes.COMMUNICATORCONTENT}, nil))
}

func (server *WebsocketServer) handleShowTutorial(connection *Connection, messageData1, messageData2 []string) {
	if len(messageData1) == 1 { //bool
		if server.ShowTutorial(connection.user.Id, Utilities.StringToBool(messageData1[0])) {
			server.messageUserByConnectionId(connection.id, TCPUtilities.ConstructMessage(MessageTypes.SHOWTUTORIAL, messageData1, nil))
			return
		}
	}
	connection.message(TCPUtilities.ConstructMessage(MessageTypes.ERROR, []string{MessageTypes.SHOWTUTORIAL}, nil))
}

func (server *WebsocketServer) handleUnknownMessage(connection *Connection, messageType string) {
	server.Base.Logger.Log("Invalid message type: " + messageType)
	connection.message(TCPUtilities.ConstructMessage("invalidMessageType", []string{messageType}, nil))
	/* connection.watchdog.Reset(1 * time.Millisecond) */
}
