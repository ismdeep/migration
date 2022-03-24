package migration

import (
	"errors"
	"github.com/jinzhu/gorm"
)

const defaultName = "default-713c05ef5cc90973"

func Init(db *gorm.DB) error {
	if db == nil {
		return errors.New("db is nil")
	}

	instances[defaultName] = &Instance{
		DB:         db,
		Migrations: nil,
	}
	return nil
}

func Add(key string, upSQLs []string, downSQLs []string) error {
	return instances[defaultName].Add(key, upSQLs, downSQLs)
}
