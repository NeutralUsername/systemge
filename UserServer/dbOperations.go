package UserServer

import (
	"UserUtilities"
	"Utilities"
	"database/sql"
	"time"
)

func (server *UserServer) insertUser(username string, password string, email string, secretQuestion string, secretAnswer string,
	userPower int, createdAt time.Time, registeredAt time.Time, confirmedAt time.Time,
	showCommunicator bool, communicatorContent int, showTutorial bool, uiTheme int) *UserUtilities.User {
	result, err := server.Base.DbConnection.Exec("INSERT INTO users ("+
		"username, password, email, secretQuestion, secretAnswer, userPower, createdAt, registeredAt, "+
		"confirmedAt, showCommunicator, communicatorContent, showTutorial, uiTheme"+
		") VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		username, password, Utilities.GenNullString(email), Utilities.GenNullString(secretQuestion),
		Utilities.GenNullString(secretAnswer), userPower, Utilities.TimeToSqlString(createdAt),
		Utilities.TimeToSqlString(registeredAt), Utilities.TimeToSqlString(confirmedAt),
		showCommunicator, communicatorContent, showTutorial, uiTheme)

	if err != nil {
		server.Base.Logger.Log(err.Error())
		return nil
	}
	userId, err := result.LastInsertId()
	if err != nil {
		server.Base.Logger.Log(err.Error())
		return nil
	}
	return UserUtilities.CreateUser(int(userId), username, password, email, secretQuestion, secretAnswer, userPower, createdAt, registeredAt, confirmedAt, showCommunicator, communicatorContent, showTutorial, uiTheme)
}

func (server *UserServer) userSelectHandler(row *sql.Row) *UserUtilities.User {
	id := 0
	username := ""
	password := ""
	email := sql.NullString{}
	secretQuestion := sql.NullString{}
	secretAnswer := sql.NullString{}
	userPower := 0
	createdAt := time.Time{}
	registeredAt := Utilities.GetNullTime()
	confirmedAt := Utilities.GetNullTime()
	showCommunicator := false
	communicatorContent := 0
	showTutorial := false
	uiTheme := 0

	err := row.Scan(&id, &username, &password, &email, &secretQuestion, &secretAnswer,
		&userPower, &createdAt, &registeredAt, &confirmedAt,
		&showCommunicator, &communicatorContent, &showTutorial, &uiTheme)
	if err != nil {
		if err != sql.ErrNoRows {
			server.Base.Logger.Log(err.Error())
		}
		return nil
	}
	return UserUtilities.CreateUser(id, username, password, email.String, secretQuestion.String, secretAnswer.String, userPower, createdAt, registeredAt, confirmedAt, showCommunicator, communicatorContent, showTutorial, uiTheme)
}

func (server *UserServer) selectUserByUsernameAndPassword(username string, password string) *UserUtilities.User {
	row := server.Base.DbConnection.QueryRow("SELECT * "+
		"FROM users "+
		"WHERE username = ? AND password = ?",
		username, password)

	return server.userSelectHandler(row)
}

func (server *UserServer) selectUserByUserId(userId int) *UserUtilities.User {
	row := server.Base.DbConnection.QueryRow("SELECT * "+
		"FROM users "+
		"WHERE id = ?",
		userId)

	return server.userSelectHandler(row)
}

func (server *UserServer) selectUserByUsername(username string) *UserUtilities.User {
	row := server.Base.DbConnection.QueryRow("SELECT * "+
		"FROM users "+
		"WHERE username = ?",
		username)

	return server.userSelectHandler(row)
}

func (server *UserServer) updateUserRegistration(userId int, username string, password string, email string, secretQuestion string,
	secretAnswer string, userPower int, createdAt time.Time, registeredAt time.Time, confirmedAt time.Time,
	showCommunicator bool, communicatorContent int, showTutorial bool, uiTheme int) bool {
	res, err := server.Base.DbConnection.Exec("UPDATE users SET "+
		"username = ?, password = ?, email = ?, secretQuestion = ?, secretAnswer = ?, userPower = ?, createdAt = ?, registeredAt = ?, "+
		"confirmedAt = ?, showCommunicator = ?, communicatorContent = ?, showTutorial = ?, uiTheme = ? "+
		"WHERE id = ? AND userPower = ?",
		username, password, Utilities.GenNullString(email), Utilities.GenNullString(secretQuestion),
		Utilities.GenNullString(secretAnswer), userPower, Utilities.TimeToSqlString(createdAt),
		Utilities.TimeToSqlString(registeredAt), Utilities.TimeToSqlString(confirmedAt),
		showCommunicator, communicatorContent, showTutorial, uiTheme, userId, UserUtilities.USER_POWER_GUEST)
	if err != nil {
		server.Base.Logger.Log(err.Error())
		return false
	}
	if count, err := res.RowsAffected(); err != nil || count == 0 {
		return false
	}
	return true
}

func (server *UserServer) updateUserShowCommunicator(userId int, showCommunicator bool) bool {
	res, err := server.Base.DbConnection.Exec("UPDATE users SET "+
		"showCommunicator = ? "+
		"WHERE id = ? AND showCommunicator != ?",
		showCommunicator, userId, showCommunicator)
	if err != nil {
		server.Base.Logger.Log(err.Error())
		return false
	}
	if count, err := res.RowsAffected(); err != nil || count == 0 {
		return false
	}
	return true
}

func (server *UserServer) updateUserCommunicatorContent(userId int, communicatorContent int) bool {
	res, err := server.Base.DbConnection.Exec("UPDATE users SET "+
		"communicatorContent = ? "+
		"WHERE id = ? AND communicatorContent != ?",
		communicatorContent, userId, communicatorContent)
	if err != nil {
		server.Base.Logger.Log(err.Error())
		return false
	}
	if count, err := res.RowsAffected(); err != nil || count == 0 {
		return false
	}
	return true
}

func (server *UserServer) updateUserShowTutorial(userId int, showTutorial bool) bool {
	res, err := server.Base.DbConnection.Exec("UPDATE users SET "+
		"showTutorial = ? "+
		"WHERE id = ? AND showTutorial != ?",
		showTutorial, userId, showTutorial)
	if err != nil {
		server.Base.Logger.Log(err.Error())
		return false
	}
	if count, err := res.RowsAffected(); err != nil || count == 0 {
		return false
	}
	return true
}

func (server *UserServer) updateUserUiTheme(userId int, uiTheme int) bool {
	res, err := server.Base.DbConnection.Exec("UPDATE users SET "+
		"uiTheme = ? "+
		"WHERE id = ? AND uiTheme != ?",
		uiTheme, userId, uiTheme)
	if err != nil {
		server.Base.Logger.Log(err.Error())
		return false
	}
	if count, err := res.RowsAffected(); err != nil || count == 0 {
		return false
	}
	return true
}

func (server *UserServer) deleteUser(userId int) bool {
	res, err := server.Base.DbConnection.Exec("DELETE FROM users WHERE id = ?", userId)
	if err != nil {
		server.Base.Logger.Log(err.Error())
		return false
	}
	if count, err := res.RowsAffected(); err != nil || count == 0 {
		return false
	}
	return true
}
