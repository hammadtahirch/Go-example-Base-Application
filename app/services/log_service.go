package services

import (
	"fmt"

	"git-lab.boldapps.net/nifty-logix/mvc/app/models"
	"git-lab.boldapps.net/nifty-logix/mvc/app/repositories"
)

// LogService ... This struct helps to inject the dependency
type LogService struct {
	ur *repositories.LogRepository
}

// LogError This function helps to return and log errors
func (l *LogService) LogError(Message string, SystemError string, Code int, dblog bool) models.Error {
	er := models.Error{
		SystemError: SystemError,
		Code:        Code,
		Message:     Message,
	}
	if dblog {
		er, _ := l.ur.SaveError("Log", er)
		fmt.Printf("%+v\n", er)
	}
	return er
}
