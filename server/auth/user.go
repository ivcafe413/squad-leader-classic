package auth

import (
	"github.com/google/uuid"
)

var users = make(map[uuid.UUID]*User)
var usernameLookup = make(map[string]*User)
//var nextID = 0 //For Dev/Test ONLY

type User struct {
	ID       uuid.UUID	`json:"id"` //TODO: int ID for User is temp/test ONLY
	Username string 	`json:"username"`
}

func NewUser(username string) *User {
	//id := nextID

	newUser := new(User)
	newUser.ID = uuid.New()
	newUser.Username = username

	users[newUser.ID] = newUser

	//Reverse Map Users by Name
	usernameLookup[username] = newUser
	//nextID++

	return newUser
}

func GetUserByName(username string) *User {
	u := usernameLookup[username]
	return u
}
