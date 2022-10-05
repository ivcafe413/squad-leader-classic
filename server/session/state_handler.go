package session

import "github.com/vagrant-technology/squad-leader/auth"

type StateHandler interface {
	UserInput(user *auth.User, msg []byte) ([]byte, error)
}