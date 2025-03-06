package cache

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/gintokos/coinder/internal/models"
)

type Cache struct {
	// permDB PermanentDB `json:"-"`

	mu sync.RWMutex `json:"-"`
	// CachedItems map[string]cachedItem `json:"cached_items"`
	CUsers map[int64]CUser `json:"cached_users"`
}

type CUser map[int]bool

// type cachedItem struct {
// 	IsIncrement bool      `json:"IsIncrement"`
// 	ExpiringAt  time.Time `json:"expiring_at"`
// }

// type PermanentDB interface {
// 	IncrementLike(coinid int, userid int64) error
// }

func New() (*Cache, error) {
	cache := &Cache{
		// CachedItems: make(map[string]cachedItem),
		// permDB:      permDB,
		CUsers: make(map[int64]CUser),
	}

	err := cache.loadFromTempFile()
	if err != nil {
		return nil, err
	}

	go cache.UpdatingLoop()

	return cache, nil
}

func (c *Cache) LikesInfo(coins []models.DBCoin, userid int64) []models.CoinResp {
	c.mu.RLock()
	defer c.mu.RUnlock()

	var rcoins []models.CoinResp
	cuser, ok := c.CUsers[userid]
	if !ok {
		for _, c := range coins {
			rcoins = append(rcoins, models.CoinResp{
				DBCoin:  c,
				IsLiked: false,
			})
		}
		return rcoins
	}

	for _, coin := range coins {
		isIncremented := cuser[coin.ID]
		rcoins = append(rcoins, models.CoinResp{
			DBCoin:  coin,
			IsLiked: isIncremented,
		})
	}

	return rcoins
}

func (c *Cache) StoreLiked(isIncrement bool, coinid int, userid int64) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.CUsers[userid]; !ok {
		c.CUsers[userid] = make(CUser)
	}
	c.CUsers[userid][coinid] = isIncrement

	// c.CachedItems[fmt.Sprintf("%d_%d", coinid, userid)] = cachedItem{
	// 	IsIncrement: isIncrement,
	// 	ExpiringAt:  time.Now().Add(time.Hour * 10),
	// }
}

func (c *Cache) UpdatingLoop() {
	// go func() {
	ticker := time.NewTicker(time.Hour * 24)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		c.CUsers = make(map[int64]CUser)
		c.mu.Unlock()
	}

	// }()

	// ticker := time.NewTicker(time.Minute * 5)
	// defer ticker.Stop()

	// for range ticker.C {
	// 	var expiredKeys []string
	// 	var expiredItems []cachedItem

	// 	c.mu.RLock()
	// 	now := time.Now()
	// 	for key, item := range c.CachedItems {
	// 		if item.ExpiringAt.Before(now) {
	// 			expiredKeys = append(expiredKeys, key)
	// 			expiredItems = append(expiredItems, item)
	// 		}
	// 	}
	// 	c.mu.RUnlock()

	// 	for i, key := range expiredKeys {
	// 		item := expiredItems[i]

	// 		if item.IsIncrement {
	// 			keyslice := strings.Split(key, "_")
	// 			coinID, _ := strconv.Atoi(keyslice[0])
	// 			userID, _ := strconv.ParseInt(keyslice[1], 10, 64)

	// 			err := c.permDB.IncrementLike(coinID, userID)
	// 			if err != nil {
	// 				slog.Error("error on incrementing like", "error", gerror.FullError(err))
	// 			} else {
	// 				c.mu.Lock()
	// 				delete(c.CachedItems, key)
	// 				c.mu.Unlock()
	// 			}
	// 		}
	// 	}
	// }
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

	files, err := os.ReadDir(tempDir)
	if err != nil {
		slog.Error("Error reading temp directory", "error", err)
	} else {
		for _, file := range files {
			if file.IsDir() {
				continue
			}

			if !strings.HasPrefix(file.Name(), "cache_") || !strings.HasSuffix(file.Name(), ".json") {
				continue
			}

			fileDate := file.Name()[6 : len(file.Name())-5]

			if fileDate != currentDate {
				filePath := filepath.Join(tempDir, file.Name())
				err := os.Remove(filePath)
				if err != nil {
					slog.Error("Failed to remove old cache file", "file", filePath, "error", err)
				} else {
					slog.Info("Removed old cache file", "file", filePath)
				}
			}
		}
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
