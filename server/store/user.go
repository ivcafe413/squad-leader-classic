package store

var users = make(map[int]*User)
var usersLookup = make(map[string]int)
var nextID = 0 //For Dev/Test ONLY

type User struct {
	ID       int    `json:"id"` //TODO: int ID for User is temp/test ONLY
	Username string `json:"username"`
}

func NewUser(username string) *User {
	id := nextID

	newUser := new(User)
	newUser.ID = id
	newUser.Username = username

	users[id] = newUser
	usersLookup[username] = id
	nextID++

	return newUser
}

func LookupUser(username string) *User {
	id := usersLookup[username]
	return users[id]
}
