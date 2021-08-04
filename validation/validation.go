package validation

import (
	"strconv"

	li "github.com/YogeshTembe/golang_project/logwrapper"
	"github.com/YogeshTembe/golang_project/model"
	uuid "github.com/satori/go.uuid"
)

var UserIds = make(map[string]struct{})

var StandardLogger = li.NewLogger()

func Validate(user *model.User, logger *li.StandardLogger) bool {
	isValid := true
	if user.Name == "" {
		logger.InvalidArg("name", user.Id.String())
		isValid = false
	}
	if user.Email == "" {
		logger.InvalidArg("Email-id", user.Id.String())
		isValid = false
	}
	if len(strconv.Itoa(user.PhoneNumber)) != 10 {
		logger.InvalidArg("phone number", user.Id.String())
		isValid = false
	}
	if user.Id.String() == "00000000-0000-0000-0000-000000000000" {
		user.Id = uuid.NewV4()
	}

	_, isFound := UserIds[user.Id.String()]
	if isFound {
		logger.InvalidArg("user-id already present for username-"+user.Name, user.Id.String())
		isValid = false
	}

	return (isValid)
}
