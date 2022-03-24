package migration

import (
	"errors"
	"fmt"
)

// Add new migration
func (receiver *Instance) Add(key string, upSQLs []string, downSQLs []string) error {
	for _, migration := range receiver.Migrations {
		if migration.Key == key {
			return fmt.Errorf("duplicated key. [%v]", key)
		}
	}

	if len(upSQLs) != len(downSQLs) {
		return errors.New("len(upSQLs) != len(downSQLs)")
	}

	receiver.Migrations = append(receiver.Migrations, &Migration{
		Key:      key,
		UpSQLs:   upSQLs,
		DownSQLs: downSQLs,
	})
	return nil
}
