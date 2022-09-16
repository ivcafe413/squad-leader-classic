package store

var Users = make(map[int]*User)
var nextID = 0 //For Dev/Test ONLY

type User struct {
	ID int `json:"id"` //TODO: int ID for User is temp/test ONLY
	Username string `json:"username"`
}

func NewUser(username string) *User {
	id := nextID
	
	newUser := new(User)
	newUser.ID = id
	newUser.Username = username

	Users[id] = newUser
	nextID++

	return newUser
}