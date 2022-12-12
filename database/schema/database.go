package schema

import (
	"context"

	"github.com/behouba/mediateq"
)

type Database struct {
	MediaTable MediaTable
}

// Delete implements mediateq.Database
func (*Database) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

// Get implements mediateq.Database
func (*Database) Get(ctx context.Context, id string) (mediateq.Media, error) {
	panic("unimplemented")
}

// Save implements mediateq.Database
func (*Database) Save(ctx context.Context, m *mediateq.Media) error {
	panic("unimplemented")
}
