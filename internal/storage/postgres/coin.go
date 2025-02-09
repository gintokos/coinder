package postgres

import (
	"github.com/gintokos/coinder/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const batchSize = 2000

func (d *Database) UpdateCoins(coins []models.DBCoin) error {
	return d.db.Transaction(func(tx *gorm.DB) error {
		if len(coins) > 100 {
			if err := tx.Exec(`
                SET LOCAL work_mem = '128MB';              
                SET LOCAL maintenance_work_mem = '512MB';     
                SET LOCAL temp_buffers = '32MB';          
            `).Error; err != nil {
				return err
			}
		}

		return tx.Session(&gorm.Session{	
			PrepareStmt:       true,
			AllowGlobalUpdate: true,
		}).Clauses(
			clause.OnConflict{
				UpdateAll: true,
			},
		).CreateInBatches(coins, batchSize).Error
	})
}

func (d *Database) Coins() ([]models.DBCoin, error) {
	var coins []models.DBCoin
	return coins, d.db.Find(&coins).Error
}