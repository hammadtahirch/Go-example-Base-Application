package repositories

import (
	"fmt"
	"time"

	"git-lab.boldapps.net/nifty-logix/mvc/app/models"
	"git-lab.boldapps.net/nifty-logix/mvc/config"
	"github.com/jinzhu/gorm"
)

// LogRepository This hepls to maintain ther dependency
type LogRepository struct {
	db *gorm.DB
}

// SaveError this function helps to save data in storage
func (lr *LogRepository) SaveError(l string, me models.Error) (int64, error) {

	s := &models.Log{
		Type:          l,
		Message:       me.Message,
		SyatemMessage: me.SystemError,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	fmt.Printf("%+v\n", s)
	db := config.DBConnection()
	er := db.Save(s).Find(s).Error
	if er != nil {
		return 0, er
	}
	return s.ID, er
}
