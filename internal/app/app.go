package app

import "context"

type App struct {
}

type DataBase interface {
}

func NewApp(db DataBase) (*App, error) {
	return &App{}, nil
}

func (a *App) Run() error {

	return nil
}

func (a *App) GraceFullShutDown(ctx context.Context) error {

	return nil
}
