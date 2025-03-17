package postgres

import (
	"github.com/gintokos/coinder/backend/internal/models"
	"gorm.io/gorm/clause"
)

func (d *Database) UpdateUser(user models.User) error {
	return d.db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&user).Error
}
