package postgres

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/gintokos/coinder/internal/models"
)

func (d *Database) SaveCoins(coins []models.DBCoin) error {
	filename := fmt.Sprintf("coins_%s.json", time.Now().Format("2006-01-02"))

	jsonData, err := json.MarshalIndent(coins, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling coins: %w", err)
	}

	if err := os.WriteFile(filename, jsonData, 0644); err != nil {
		return fmt.Errorf("error writing file: %w", err)
	}

	slog.Info("coins were saved", "filename", filename)
	return nil
}
