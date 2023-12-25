package UserUtilities

import (
	"encoding/json"
	"sync"
	"time"
)

type User struct {
	mutex       *sync.Mutex
	connections map[string]struct{}

	Id                  int       `json:"userId"`
	Username            string    `json:"username"`
	Password            string    `json:"password"`
	Email               string    `json:"email"`
	SecretQuestion      string    `json:"secretQuestion"`
	SecretAnswer        string    `json:"secretAnswer"`
	UserPower           int       `json:"userPower"`
	CreatedAt           time.Time `json:"createdAt"`
	RegisteredAt        time.Time `json:"registeredAt"`
	ConfirmedAt         time.Time `json:"confirmedAt"`
	ShowCommunicator    bool      `json:"showCommunicator"`
	CommunicatorContent int       `json:"communicatorContent"`
	ShowTutorial        bool      `json:"showTutorial"`
	UiTheme             int       `json:"uiTheme"`
}

type Users []*User

func (user *User) ToStringPrivate() []byte {
	userMarshal, _ := json.Marshal(user)
	return userMarshal
}

func (user *User) ToStringPublic() []byte {
	userMarshal, _ := json.Marshal(struct {
		Id           int       `json:"userId"`
		Username     string    `json:"username"`
		UserPower    int       `json:"userPower"`
		CreatedAt    time.Time `json:"createdAt"`
		RegisteredAt time.Time `json:"registeredAt"`
		ConfirmedAt  time.Time `json:"confirmedAt"`
	}{
		user.Id,
		user.Username,
		user.UserPower,
		user.CreatedAt,
		user.RegisteredAt,
		user.ConfirmedAt,
	})
	return userMarshal
}

func CreateUser(userId int, name string, password string, email string, secretQuestion string, secretAnswer string,
	userPower int, createdAt time.Time, registeredAt time.Time, confirmedAt time.Time,
	showCommunicator bool, communicatorContent int, showTutorial bool, uiTheme int) *User {

	return &User{
		mutex:       &sync.Mutex{},
		connections: map[string]struct{}{},

		Id:                  userId,
		Username:            name,
		Password:            password,
		Email:               email,
		SecretQuestion:      secretQuestion,
		SecretAnswer:        secretAnswer,
		UserPower:           userPower,
		CreatedAt:           createdAt,
		RegisteredAt:        registeredAt,
		ConfirmedAt:         confirmedAt,
		ShowCommunicator:    showCommunicator,
		CommunicatorContent: communicatorContent,
		ShowTutorial:        showTutorial,
		UiTheme:             uiTheme,
	}
}
