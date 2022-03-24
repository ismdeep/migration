package migration

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"    // load mssql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"    // load mysql driver
	_ "github.com/jinzhu/gorm/dialects/postgres" // load postgres driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"   // load sqlite driver
)

type Migration struct {
	Key      string
	UpSQLs   []string
	DownSQLs []string
}

type Instance struct {
	DB         *gorm.DB
	Migrations []*Migration
}
