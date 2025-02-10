package postgres

import (
	"github.com/gintokos/coinder/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	batchSize     = 2000
	BY_PRICE      = "BY_PRICE"
	BY_MARKET_CAP = "BY_MARKET_CAP"
)

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

func (d *Database) DefaultSearchCoins(opt models.SearchCoinOpt) ([]models.DBCoin, error) {
    db := d.db
    query := db.Model(&models.DBCoin{})

    switch opt.SortedBy {
    case BY_PRICE:
        query = query.Order("coins.price DESC")
    case BY_MARKET_CAP:
        query = query.Order("coins.market_cap DESC")
    }

    if !opt.LikedByUser {
        query = d.addOnlyUnlikedCoins(query, opt.UserID)
    }

    var coins []models.DBCoin
    result := query.Offset(opt.Limit * opt.Page).Limit(opt.Limit).Find(&coins)

    return coins, result.Error
}

func (d *Database) addOnlyUnlikedCoins(tx *gorm.DB, userID int64) *gorm.DB {
    return tx.Where("coins.id NOT IN (?)",
        d.db.Table("likes").
            Select("like_coin_id").
            Where("like_user_id = ?", userID))
}

func (d *Database) CustomSearchCoins(opt models.SearchCoinOpt) ([]models.DBCoin, error) {
	return nil, nil
}
