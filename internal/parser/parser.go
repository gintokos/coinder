package parser

import "context"

type Parser struct {
}

type Database interface {
}

func NewParser(db Database) (*Parser, error) {

	return &Parser{}, nil
}

func (p *Parser) Run() error {
	return nil
}

func (p *Parser) GraceFullShutDown(ctx context.Context) error {
	return nil
}
