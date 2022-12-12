package schema

import (
	"context"

	"github.com/behouba/mediateq"
)

type MediaTable interface {
	Insert(ctx context.Context, media *mediateq.Media) (int, error)
	SelectByUID(ctx context.Context, id string) (*mediateq.Media, error)
	Delete(ctx context.Context, id string) error
}
