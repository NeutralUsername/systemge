package WebsocketServer

import (
	"MessageTypes"
	"Ports"
	"TCPUtilities"
	"Utilities"
)

func (server *WebsocketServer) GetUserPrivate(userId int) string {
	messageType, messageData1, _ := TCPUtilities.ParseMessage(TCPUtilities.AskServer(TCPUtilities.ConstructMessage(MessageTypes.GETUSERPRIVATE, nil, []string{Utilities.IntToString(userId)}), Ports.USERSERVER_TCP_LISTENER))
	if messageType == MessageTypes.TRUE {
		return messageData1[0]
	}
	return ""
}
func (server *WebsocketServer) GetUserPublic(userIds []string) []string {
	_, messageData1, _ := TCPUtilities.ParseMessage(TCPUtilities.AskServer(TCPUtilities.ConstructMessage(MessageTypes.GETUSERPUBLIC, nil, userIds), Ports.USERSERVER_TCP_LISTENER))
	return messageData1
}
func (server *WebsocketServer) NewGuest() int {
	messageType, messageData1, _ := TCPUtilities.ParseMessage(TCPUtilities.AskServer(TCPUtilities.ConstructMessage(MessageTypes.NEWGUEST, nil, nil), Ports.USERSERVER_TCP_LISTENER))
	if messageType == MessageTypes.TRUE {
		return Utilities.StringToInt(messageData1[0])
	}
	return -1
}
func (server *WebsocketServer) AuthenticateCredentials(username string, password string) int {
	messageType, messageData1, _ := TCPUtilities.ParseMessage(TCPUtilities.AskServer(TCPUtilities.ConstructMessage(MessageTypes.AUTHENTICATE, []string{username, password}, nil), Ports.USERSERVER_TCP_LISTENER))
	if messageType == MessageTypes.TRUE {
		return Utilities.StringToInt(messageData1[0])
	}
	return -1
}

func (server *WebsocketServer) Register(userId int, username, password, email, secretQuestion, secretAnswer string) bool {
	messageType, _, _ := TCPUtilities.ParseMessage(TCPUtilities.AskServer(TCPUtilities.ConstructMessage(MessageTypes.REGISTER, []string{username, password, email, secretQuestion, secretAnswer}, []string{Utilities.IntToString(userId)}), Ports.USERSERVER_TCP_LISTENER))
	return messageType == MessageTypes.TRUE
}
func (server *WebsocketServer) UiTheme(userId int, uiTheme string) bool {
	messageType, _, _ := TCPUtilities.ParseMessage(TCPUtilities.AskServer(TCPUtilities.ConstructMessage(MessageTypes.UITHEME, []string{uiTheme}, []string{Utilities.IntToString(userId)}), Ports.USERSERVER_TCP_LISTENER))
	return messageType == MessageTypes.TRUE
}
func (server *WebsocketServer) ShowCommunicator(userId int, showCommunicator bool) bool {
	messageType, _, _ := TCPUtilities.ParseMessage(TCPUtilities.AskServer(TCPUtilities.ConstructMessage(MessageTypes.SHOWCOMMUNICATOR, []string{Utilities.BoolToString(showCommunicator)}, []string{Utilities.IntToString(userId)}), Ports.USERSERVER_TCP_LISTENER))
	return messageType == MessageTypes.TRUE
}
func (server *WebsocketServer) CommunicatorContent(userId int, communicatorContent string) bool {
	messageType, _, _ := TCPUtilities.ParseMessage(TCPUtilities.AskServer(TCPUtilities.ConstructMessage(MessageTypes.COMMUNICATORCONTENT, []string{communicatorContent}, []string{Utilities.IntToString(userId)}), Ports.USERSERVER_TCP_LISTENER))
	return messageType == MessageTypes.TRUE
}
func (server *WebsocketServer) ShowTutorial(userId int, showTutorial bool) bool {
	messageType, _, _ := TCPUtilities.ParseMessage(TCPUtilities.AskServer(TCPUtilities.ConstructMessage(MessageTypes.SHOWTUTORIAL, []string{Utilities.BoolToString(showTutorial)}, []string{Utilities.IntToString(userId)}), Ports.USERSERVER_TCP_LISTENER))
	return messageType == MessageTypes.TRUE
}
