package postgres

import (
	"fmt"
	"time"

	"github.com/gintokos/coinder/internal/config"
	"github.com/gintokos/coinder/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

var gormConfig = &gorm.Config{
	NowFunc: func() time.Time {
		return time.Now().UTC()
	}}

func NewDatabase() (*Database, error) {
	var d Database

	cfg := config.GetConfig().Database
	err := d.CreateDataBase()
	if err != nil {
		return nil, fmt.Errorf("creating database was ended with error")
	}

	dsnnew := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.Port,
	)

	dbnew, err := gorm.Open(postgres.Open(dsnnew), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("error on connecting db: %v", err)
	}

	d.db = dbnew
	err = d.CreateTables()
	if err != nil {
		return nil, fmt.Errorf("error on creating tables")
	}

	return &d, nil
}

func (d *Database) CreateTables() error {
	db := d.db

	var Allcoins models.DBCoin
	err := db.AutoMigrate(&Allcoins)
	if err != nil {
		return fmt.Errorf("error on automigrating DBcoins: %w", err)
	}

	return nil
}

func (d *Database) CreateDataBase() error {
	cfg := config.GetConfig().Database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=postgres port=%s sslmode=disable Timezone=UTC",
		cfg.Host, cfg.User, cfg.Password, cfg.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		return fmt.Errorf("error connecting to postgres: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("error getting sql.DB: %w", err)
	}
	defer sqlDB.Close()

	var exists bool
	err = db.Raw("SELECT EXISTS(SELECT datname from pg_catalog.pg_database WHERE datname = ?)", cfg.Name).Scan(&exists).Error
	if err != nil {
		return fmt.Errorf("error in checking existing database: %w", err)
	}

	if !exists {
		createDBq := fmt.Sprintf("CREATE DATABASE %s", cfg.Name)
		if err := db.Exec(createDBq).Error; err != nil {
			return fmt.Errorf("error on creating database: %w", err)
		}
	}

	return nil
}
