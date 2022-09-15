package store

var Users = make(map[int]User)

type User struct {
	ID int `json:"id"` //TODO: int ID for User is temp/test ONLY
	Username string `json:"username"`
}