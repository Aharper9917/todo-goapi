package tools

import (
	log "github.com/sirupsen/logrus"
)

// Database collections
type LoginDetails struct {
	AuthToken string
	Username  string
}
type TodoDetails struct {
	Title       string
	Description string
	DueDate     string
	UserId      string
}
type UserDetails struct {
	FirstName string
	LastName  string
	Username  string
}

type DatabaseInterface interface {
	GetUserLoginDetails(id string) (*LoginDetails, error)
	GetTodoDetails(id string) (*TodoDetails, error)
	GetUserDetails(id string) (*UserDetails, error)
	SetupDatabase() error
}

func NewDatabase() (*DatabaseInterface, error) {

	var database DatabaseInterface = &mockDB{}

	var err error = database.SetupDatabase()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &database, nil
}
