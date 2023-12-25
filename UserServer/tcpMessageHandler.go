package UserServer

import (
	"MessageTypes"
	"TCPUtilities"
	"UserUtilities"
	"Utilities"
	"net"
	"time"
)

func (server *UserServer) HandleTcpMessage(incomingTcpConnection *net.Conn, message []byte) {
	messageType, messageData1, messageData2 := TCPUtilities.ParseMessage(message)
	switch messageType {
	case MessageTypes.GETUSERID:
		server.handleGetUserId(incomingTcpConnection, messageData1, messageData2)
	case MessageTypes.USERIDEXISTS:
		server.handleUserIdExists(incomingTcpConnection, messageData1, messageData2)
	case MessageTypes.GETUSERPRIVATE:
		server.handleGetUserPrivate(incomingTcpConnection, messageData1, messageData2)
	case MessageTypes.GETUSERPUBLIC:
		server.handleGetUserPublic(incomingTcpConnection, messageData1, messageData2)
	case MessageTypes.AUTHENTICATE:
		server.handleAuthenticate(incomingTcpConnection, messageData1, messageData2)
	case MessageTypes.NEWGUEST:
		server.handleNewGuest(incomingTcpConnection, messageData1, messageData2)
	case MessageTypes.REGISTER:
		server.handleRegister(incomingTcpConnection, messageData1, messageData2)
	case MessageTypes.SHOWCOMMUNICATOR:
		server.handleShowCommunicator(incomingTcpConnection, messageData1, messageData2)
	case MessageTypes.COMMUNICATORCONTENT:
		server.handleCommunicatorContent(incomingTcpConnection, messageData1, messageData2)
	case MessageTypes.SHOWTUTORIAL:
		server.handleShowTutorial(incomingTcpConnection, messageData1, messageData2)
	case MessageTypes.UITHEME:
		server.handleUiTheme(incomingTcpConnection, messageData1, messageData2)
	}
}

func (server *UserServer) handleGetUserId(incomingTcpConnection *net.Conn, messageData1, messageData2 []string) {
	username := messageData1[0]
	if user := server.selectUserByUsername(username); user != nil {
		TCPUtilities.MessageConnection(incomingTcpConnection, TCPUtilities.ConstructMessage(MessageTypes.TRUE, []string{Utilities.IntToString(user.Id)}, nil))
		return
	}
	TCPUtilities.MessageConnection(incomingTcpConnection, TCPUtilities.ConstructMessage(MessageTypes.ERROR, []string{MessageTypes.GETUSERID}, nil))
}

func (server *UserServer) handleUserIdExists(incomingTcpConnection *net.Conn, messageData1, messageData2 []string) {
	if userId := Utilities.StringToInt(messageData1[0]); server.selectUserByUserId(userId) != nil {
		TCPUtilities.MessageConnection(incomingTcpConnection, TCPUtilities.ConstructMessage(MessageTypes.TRUE, nil, nil))
		return
	}
	TCPUtilities.MessageConnection(incomingTcpConnection, TCPUtilities.ConstructMessage(MessageTypes.ERROR, []string{MessageTypes.USERIDEXISTS}, nil))
}

func (server *UserServer) handleGetUserPrivate(incomingTcpConnection *net.Conn, messageData1, messageData2 []string) {
	reponseMessageData := []string{}
	for _, userId := range messageData2 {
		if user := server.selectUserByUserId(Utilities.StringToInt(userId)); user != nil {
			reponseMessageData = append(reponseMessageData, string(user.ToStringPrivate()))
		} else {
			TCPUtilities.MessageConnection(incomingTcpConnection, TCPUtilities.ConstructMessage(MessageTypes.ERROR, []string{MessageTypes.GETUSERPRIVATE}, nil))
			return
		}
	}
	TCPUtilities.MessageConnection(incomingTcpConnection, TCPUtilities.ConstructMessage(MessageTypes.TRUE, reponseMessageData, nil))
}

func (server *UserServer) handleGetUserPublic(incomingTcpConnection *net.Conn, messageData1, messageData2 []string) {
	reponseMessageData := []string{}
	for _, userId := range messageData2 {
		if user := server.selectUserByUserId(Utilities.StringToInt(userId)); user != nil {
			reponseMessageData = append(reponseMessageData, string(user.ToStringPublic()))
		} else {
			TCPUtilities.MessageConnection(incomingTcpConnection, TCPUtilities.ConstructMessage(MessageTypes.ERROR, []string{MessageTypes.GETUSERPUBLIC}, nil))
			return
		}
	}
	TCPUtilities.MessageConnection(incomingTcpConnection, TCPUtilities.ConstructMessage(MessageTypes.TRUE, reponseMessageData, nil))
}

func (server *UserServer) handleAuthenticate(incomingTcpConnection *net.Conn, messageData1, messageData2 []string) {
	username := messageData1[0]
	password := messageData1[1]
	if user := server.selectUserByUsernameAndPassword(username, password); user != nil {
		TCPUtilities.MessageConnection(incomingTcpConnection, TCPUtilities.ConstructMessage(MessageTypes.TRUE, []string{Utilities.IntToString(user.Id)}, nil))
		return
	}
	TCPUtilities.MessageConnection(incomingTcpConnection, TCPUtilities.ConstructMessage(MessageTypes.ERROR, []string{MessageTypes.AUTHENTICATE}, nil))
}

func (server *UserServer) handleNewGuest(incomingTcpConnection *net.Conn, messageData1, messageData2 []string) {
	user := server.insertUser(server.Base.Random.GenerateRandomString(UserUtilities.DEFAULT_USERNAME_LENGTH, UserUtilities.VALID_USERNAME_CHARS),
		Utilities.SHA256string(server.Base.Random.GenerateRandomString(UserUtilities.DEFAULT_USERNAME_LENGTH, UserUtilities.VALID_USERNAME_CHARS)),
		"", "", "", UserUtilities.USER_POWER_GUEST, time.Now(), Utilities.GetNullTime(), Utilities.GetNullTime(), false, UserUtilities.COMMUNICATOR_CONTENT_CONTACTS, true, UserUtilities.UI_THEME_DARK)

	for i := 0; user == nil && i < MAX_NEW_USER_ATTEMPTS; i++ {
		user = server.insertUser(server.Base.Random.GenerateRandomString(UserUtilities.DEFAULT_USERNAME_LENGTH, UserUtilities.VALID_USERNAME_CHARS),
			Utilities.SHA256string(server.Base.Random.GenerateRandomString(UserUtilities.DEFAULT_USERNAME_LENGTH, UserUtilities.VALID_USERNAME_CHARS)),
			"", "", "", UserUtilities.USER_POWER_GUEST, time.Now(), (time.Time{}), (time.Time{}), false, UserUtilities.COMMUNICATOR_CONTENT_CONTACTS, true, UserUtilities.UI_THEME_DARK)
	}
	if user != nil {
		if true { //further operations required for new users
			TCPUtilities.MessageConnection(incomingTcpConnection, TCPUtilities.ConstructMessage(MessageTypes.TRUE, []string{Utilities.IntToString(user.Id)}, nil))
			return
		}
		server.deleteUser(user.Id)
	}
	TCPUtilities.MessageConnection(incomingTcpConnection, TCPUtilities.ConstructMessage(MessageTypes.ERROR, []string{MessageTypes.NEWGUEST}, nil))
}

func (server *UserServer) handleRegister(incomingTcpConnection *net.Conn, messageData1, messageData2 []string) {
	userId := Utilities.StringToInt(messageData2[0])
	username := messageData1[0]
	password := messageData1[1]
	email := messageData1[2]
	secretQuestion := messageData1[3]
	secretAnswer := messageData1[4]
	if UserUtilities.IsValidUsername(username) {
		if UserUtilities.IsValidPassword(password) {
			if UserUtilities.IsValidEmail(email) {
				if UserUtilities.IsValidSecrets(secretQuestion, secretAnswer) {

					if server.updateUserRegistration(userId, username, password, email, secretQuestion, secretAnswer, UserUtilities.USER_POWER_UNCONFIRMED,
						time.Now(), Utilities.GetNullTime(), Utilities.GetNullTime(), false, UserUtilities.COMMUNICATOR_CONTENT_CONTACTS, true, UserUtilities.UI_THEME_DARK) {

						TCPUtilities.MessageConnection(incomingTcpConnection, TCPUtilities.ConstructMessage(MessageTypes.TRUE, nil, nil))
						return
					}
				}
			}
		}
	}
	TCPUtilities.MessageConnection(incomingTcpConnection, TCPUtilities.ConstructMessage(MessageTypes.ERROR, []string{MessageTypes.REGISTER}, nil))
}

func (server *UserServer) handleShowCommunicator(incomingTcpConnection *net.Conn, messageData1, messageData2 []string) {
	userId := Utilities.StringToInt(messageData2[0])
	showCommunicator := messageData1[0] == MessageTypes.TRUE

	if server.updateUserShowCommunicator(userId, showCommunicator) {
		TCPUtilities.MessageConnection(incomingTcpConnection, TCPUtilities.ConstructMessage(MessageTypes.TRUE, nil, nil))
		return
	}
	TCPUtilities.MessageConnection(incomingTcpConnection, TCPUtilities.ConstructMessage(MessageTypes.ERROR, []string{MessageTypes.SHOWCOMMUNICATOR}, nil))
}

func (server *UserServer) handleShowTutorial(incomingTcpConnection *net.Conn, messageData1, messageData2 []string) {
	userId := Utilities.StringToInt(messageData2[0])
	showTutorial := messageData1[0] == MessageTypes.TRUE

	if server.updateUserShowTutorial(userId, showTutorial) {
		TCPUtilities.MessageConnection(incomingTcpConnection, TCPUtilities.ConstructMessage(MessageTypes.TRUE, nil, nil))
		return
	}
	TCPUtilities.MessageConnection(incomingTcpConnection, TCPUtilities.ConstructMessage(MessageTypes.ERROR, []string{MessageTypes.SHOWTUTORIAL}, nil))
}

func (server *UserServer) handleUiTheme(incomingTcpConnection *net.Conn, messageData1, messageData2 []string) {
	userId := Utilities.StringToInt(messageData2[0])
	uiTheme := Utilities.StringToInt(messageData1[0])
	if uiTheme == UserUtilities.UI_THEME_LIGHT || uiTheme == UserUtilities.UI_THEME_DARK {
		if server.updateUserUiTheme(userId, uiTheme) {
			TCPUtilities.MessageConnection(incomingTcpConnection, TCPUtilities.ConstructMessage(MessageTypes.TRUE, nil, nil))
			return
		}
	}
	TCPUtilities.MessageConnection(incomingTcpConnection, TCPUtilities.ConstructMessage(MessageTypes.ERROR, []string{MessageTypes.UITHEME}, nil))
}

func (server *UserServer) handleCommunicatorContent(incomingTcpConnection *net.Conn, messageData1, messageData2 []string) {
	userId := Utilities.StringToInt(messageData2[0])
	communicatorContent := Utilities.StringToInt(messageData1[0])
	if communicatorContent == UserUtilities.COMMUNICATOR_CONTENT_CONTACTS ||
		communicatorContent == UserUtilities.COMMUNICATOR_CONTENT_CHATS ||
		communicatorContent == UserUtilities.COMMUNICATOR_CONTENT_NOTIFICATONS ||
		communicatorContent == UserUtilities.COMMUNICATOR_CONTENT_BLOCKS {
		if server.updateUserCommunicatorContent(userId, communicatorContent) {
			TCPUtilities.MessageConnection(incomingTcpConnection, TCPUtilities.ConstructMessage(MessageTypes.TRUE, nil, nil))
			return
		}
	}
	TCPUtilities.MessageConnection(incomingTcpConnection, TCPUtilities.ConstructMessage(MessageTypes.ERROR, []string{MessageTypes.COMMUNICATORCONTENT}, nil))
}
