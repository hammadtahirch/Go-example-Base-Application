package services

import (
	"fmt"
	"os"
	"strconv"

	"github.com/hammadtahirch/golang_basic_app/app/models"
	"github.com/hammadtahirch/golang_basic_app/app/models/repositories"
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

		id, _ := l.ur.SaveError("Log", er)
		fmt.Printf("%+v\n", id)
		var ad, _ = strconv.ParseBool(os.Getenv("APP_DEBUG"))
		er.ID = id
		if os.Getenv("APP_ENV") == "production" && ad == false {
			er.SystemError = "Site Working in production mood. See db logs"
		}
	}
	return er
}
