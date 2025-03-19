package storage

type Storage struct {
	PostgresDB
}

func NewStorage() (*Storage, error) {
	postgres, err := NewPostgresDB()
	if err != nil {
		return &Storage{}, err
	}
	return &Storage{PostgresDB: postgres}, nil
}