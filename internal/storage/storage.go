package storage

import (
	"context"

	"github.com/billyboar/bytely/schema"
)

type Storage interface {
	AddURL(ctx context.Context, url schema.URL) error
	GetURL(ctx context.Context, shortURLKey string) (*schema.URL, error)
	DeleteURL(ctx context.Context, shortURLKey string) (bool, error)
	IncrementClicks(ctx context.Context, shortURLKey string) error
	Ping() error
	FlushAll() error
	Close()
}
