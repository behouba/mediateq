package postgres

import (
	"context"

	"github.com/behouba/mediateq"
)

func NewDatabase(cfg *mediateq.DBConfig) (mediateq.Database, error) {
	return DB{}, nil
}

type DB struct {
}

// Delete implements mediateq.Database
func (DB) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

// Get implements mediateq.Database
func (DB) Get(ctx context.Context, id string) (mediateq.Media, error) {
	panic("unimplemented")
}

// Save implements mediateq.Database
func (DB) Save(ctx context.Context, m *mediateq.Media) error {
	panic("unimplemented")
}
