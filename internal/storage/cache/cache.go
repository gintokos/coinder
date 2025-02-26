package cache

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gintokos/coinder/pkg/gerror"
)

type Cache struct {
	permDB PermanentDB `json:"-"`

	mu          sync.RWMutex          `json:"-"`
	CashedUsers map[string]cashedItem `json:"cashed_items"`
}

type cashedItem struct {
	IsIncrement bool      `json:"key"`
	ExpiringAt  time.Time `json:"expiring_at"`
}

type PermanentDB interface {
	IncrementLike(coinid int, userid int64) error
}

func New(permDB PermanentDB) (*Cache, error) {
	cache := &Cache{
		CashedUsers: make(map[string]cashedItem),
		permDB:      permDB,
	}

	err := cache.loadFromTempFile()
	if err != nil {
		return nil, err
	}

	go cache.UpdatingLoop()

	return cache, nil
}

func (c *Cache) ChangeLike(isIncrement bool, coinid int, userid int64) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.CashedUsers[fmt.Sprintf("%d_%d", coinid, userid)] = cashedItem{
		IsIncrement: isIncrement,
		ExpiringAt:  time.Now().Add(time.Hour * 10),
	}
}

func (c *Cache) UpdatingLoop() {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		var expiredKeys []string
		var expiredItems []cashedItem

		c.mu.RLock()
		now := time.Now()
		for key, item := range c.CashedUsers {
			if item.ExpiringAt.Before(now) {
				expiredKeys = append(expiredKeys, key)
				expiredItems = append(expiredItems, item)
			}
		}
		c.mu.RUnlock()

		for i, key := range expiredKeys {
			item := expiredItems[i]

			c.mu.Lock()
			delete(c.CashedUsers, key)
			c.mu.Unlock()

			if item.IsIncrement {
				ketslice := strings.Split(key, "_")
				coinID, _ := strconv.Atoi(ketslice[0])
				userID, _ := strconv.ParseInt(ketslice[1], 10, 64)

				err := c.permDB.IncrementLike(coinID, userID)
				if err != nil {
					slog.Error("error on incrementing like", "error", gerror.FullError(err))
				}
			}
		}
	}
}

func (c *Cache) GraceFullShutDown() error {

	tempDir := "temp"
	if _, err := os.Stat(tempDir); os.IsNotExist(err) {
		err = os.Mkdir(tempDir, 0755)
		if err != nil {
			return fmt.Errorf("failed to create temp directory: %v", err)
		}
	}

	currentDate := time.Now().Format("2006-01-02")
	filename := filepath.Join(tempDir, fmt.Sprintf("cache_%s.json", currentDate))

	c.mu.RLock()
	data, err := json.MarshalIndent(c, "", "  ")
	c.mu.RUnlock()
	if err != nil {
		return fmt.Errorf("failed to marshal cache data: %v", err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write cache to file: %v", err)
	}

	return nil
}

func (c *Cache) loadFromTempFile() error {
	tempDir := "temp"

	if _, err := os.Stat(tempDir); os.IsNotExist(err) {
		return nil
	}

	currentDate := time.Now().Format("2006-01-02")
	filename := filepath.Join(tempDir, fmt.Sprintf("cache_%s.json", currentDate))

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read cache file: %v", err)
	}

	err = json.Unmarshal(data, c)
	if err != nil {
		return fmt.Errorf("failed to unmarshal cache data: %v", err)
	}

	slog.Info("Cache successfully loaded from " + filename)
	return nil
}
