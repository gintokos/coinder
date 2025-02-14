package postgres

import (
	"errors"

	"github.com/gintokos/coinder/internal/constants"
	"github.com/gintokos/coinder/internal/models"
	"github.com/gintokos/coinder/pkg/gerror"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (d *Database) UpdateCoins(coins []models.DBCoin) error {
	return d.db.Transaction(func(tx *gorm.DB) error {
		if len(coins) > 100 {
			if err := tx.Exec(`
                SET LOCAL work_mem = '128MB';              
                SET LOCAL maintenance_work_mem = '512MB';     
                SET LOCAL temp_buffers = '32MB';          
            `).Error; err != nil {
				return gerror.New(err, constants.ErrDatabase , 500)
			}
		}

		return tx.Session(&gorm.Session{
			PrepareStmt:       true,
			AllowGlobalUpdate: true,
		}).Clauses(
			clause.OnConflict{
				UpdateAll: true,
			},
		).CreateInBatches(coins, constants.BATCH_SIZE).Error
	})
}

// DefaultSearchCoins retrieves a list of coins from the database based on the specified
// search options. It supports sorting by price or market cap and can filter out coins
// liked by the user. The results are paginated according to the provided limit and page
// number. Returns an error if no coins are found.
func (d *Database) DefaultSearchCoins(opt models.SearchCoinOpt) ([]models.DBCoin, error) {
	db := d.db
	query := db.Model(&models.DBCoin{})

	switch opt.SortedBy {
	case constants.BY_PRICE:
		query = query.Order("coins.price DESC")
	case constants.BY_MARKET_CAP:
		query = query.Order("coins.market_cap DESC")
	}

	if !opt.LikedByUser {
		query = d.addOnlyUnlikedCoins(query, opt.UserID)
	}

	var coins []models.DBCoin
	result := query.Offset(opt.Limit * opt.Page).Limit(opt.Limit).Find(&coins)

	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, gerror.New(result.Error, constants.ErrNotFound, 404)
	}

	return coins, gerror.New(result.Error, constants.ErrServer, 500)
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
