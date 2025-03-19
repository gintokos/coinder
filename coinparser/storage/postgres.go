package storage

import (
	"fmt"
	"time"

	"github.com/gintokos/coinder/backend/constants"
	"github.com/gintokos/coinder/backend/models"
	"github.com/gintokos/coinder/backend/pkg/gerror"
	"github.com/gintokos/coinder/coinparser/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

type PostgresDB struct {
	db *gorm.DB
}

var gormConfig = &gorm.Config{
	NowFunc: func() time.Time {
		return time.Now().UTC()
	}}

func NewPostgresDB() (PostgresDB, error) {
	var d PostgresDB

	cfg := config.Database()

	dsnnew := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.Port,
	)

	dbnew, err := gorm.Open(postgres.Open(dsnnew), gormConfig)
	if err != nil {
		return PostgresDB{}, fmt.Errorf("error on connecting db: %v", err)
	}

	d.db = dbnew
	d.db.Config.Logger = logger.Default.LogMode(logger.Silent)

	return PostgresDB{}, nil
}

func (d *PostgresDB) UpdateCoins(coins []models.DBCoin) error {
	return d.db.Transaction(func(tx *gorm.DB) error {
		if len(coins) > 100 {
			if err := tx.Exec(`
                SET LOCAL work_mem = '128MB';              
                SET LOCAL maintenance_work_mem = '512MB';     
                SET LOCAL temp_buffers = '32MB';          
            `).Error; err != nil {
				return gerror.New(err, constants.ErrDatabase, 500)
			}
		}

		return tx.Session(&gorm.Session{
			PrepareStmt:       true,
			AllowGlobalUpdate: true,
		}).Clauses(
			clause.OnConflict{
				Columns: []clause.Column{{Name: "id"}},
				DoUpdates: clause.AssignmentColumns([]string{
					"name", "symbol", "slug", "last_updated", "date_added", "date_launched",
					"price", "volume_24h", "volume_change_24h", "percent_change_1h",
					"percent_change_24h", "percent_change_7d", "market_cap",
					"market_cap_dominance", "fully_diluted_market_cap",
					"logo", "description",
				}),
			},
		).CreateInBatches(coins, constants.BATCH_SIZE).Error
	})
}
