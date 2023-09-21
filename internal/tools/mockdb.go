package tools

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"CBA321-FED654-IHG987": {
		AuthToken: "123ABC",
		Username:  "alex",
	},
	"FED654-CBA321-IHG987": {
		AuthToken: "456DEF",
		Username:  "joe",
	},
	"IHG987-CBA321-FED654": {
		AuthToken: "789GHI",
		Username:  "allen",
	},
}

var mockTodoDetails = map[string]TodoDetails{
	"ABED-1234-XXX0": {
		Title:       "Do Laundry",
		Description: "",
		DueDate:     "",
		UserId:      "CBA321-FED654-IHG987",
	},
	"ABED-1234-XXX1": {
		Title:       "Do Laundry",
		Description: "",
		DueDate:     "",
		UserId:      "FED654-CBA321-IHG987",
	},
	"ABED-1234-XXX2": {
		Title:       "Do Laundry",
		Description: "",
		DueDate:     "",
		UserId:      "IHG987-CBA321-FED654",
	},
	"ABED-1234-XXX3": {
		Title:       "Do Dishes",
		Description: "",
		DueDate:     "",
		UserId:      "CBA321-FED654-IHG987",
	},
	"ABED-1234-XXX4": {
		Title:       "Do Dished",
		Description: "",
		DueDate:     "",
		UserId:      "FED654-CBA321-IHG987",
	},
	"ABED-1234-XXX5": {
		Title:       "Do Dishes",
		Description: "",
		DueDate:     "",
		UserId:      "IHG987-CBA321-FED654",
	},
}

var mockUserDetails = map[string]UserDetails{
	"CBA321-FED654-IHG987": {
		FirstName: "Alex",
		LastName:  "Martinez",
		Username:  "alex",
	},
	"FED654-CBA321-IHG987": {
		FirstName: "Joe",
		LastName:  "Doe",
		Username:  "joe",
	},
	"IHG987-CBA321-FED654": {
		FirstName: "Allen",
		LastName:  "Harper",
		Username:  "allen",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) (*LoginDetails, error) {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	clientData := LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		err := errors.New("an error occurred while fetching Login details")
		log.Error(err)
		return nil, err
	}

	return &clientData, nil
}

func (d *mockDB) GetTodoDetails(id string) (*TodoDetails, error) {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = TodoDetails{}
	clientData, ok := mockTodoDetails[id]
	if !ok {
		err := errors.New("an error occurred while fetching Todo details")
		log.Error(err)
		return nil, err
	}

	return &clientData, nil
}

func (d *mockDB) GetUserDetails(id string) (*UserDetails, error) {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = UserDetails{}
	clientData, ok := mockUserDetails[id]
	if !ok {
		err := errors.New("an error occurred while fetching User details")
		log.Error(err)
		return nil, err
	}

	return &clientData, nil
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
