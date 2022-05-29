package pg

import (
	"context"
	"time"

	"github.com/billyboar/bytely/schema"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	defaultTimeout = 3
	totalRetries   = 5
)

type Client struct {
	pool *pgxpool.Pool
}

func NewClient(connStr string) (*Client, error) {
	var pool *pgxpool.Pool
	var err error

	for i := 0; i < totalRetries; i++ {
		pool, err = pgxpool.Connect(context.Background(), connStr)
		if err == nil {
			break
		}
		time.Sleep(time.Duration(i*2) * time.Second) // backoff exponentially
	}

	if err != nil {
		return nil, err
	}

	return &Client{
		pool: pool,
	}, nil
}

func (pg *Client) AddURL(ctx context.Context, url schema.URL) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*defaultTimeout)
	defer cancel()

	query := `INSERT INTO urls (short_url_key, original_url) VALUES ($1, $2)`
	_, err := pg.pool.Exec(ctx, query, url.ShortURLKey, url.OriginalURL)
	return err
}

func (pg *Client) GetURL(ctx context.Context, shortURLKey string) (*schema.URL, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*defaultTimeout)
	defer cancel()

	query := `SELECT created_at, short_url_key, original_url, clicks FROM urls WHERE short_url_key = $1`
	row := pg.pool.QueryRow(ctx, query, shortURLKey)

	var url schema.URL
	err := row.Scan(&url.CreatedAt,
		&url.ShortURLKey, &url.OriginalURL, &url.Clicks)
	if err != nil {
		if err != pgx.ErrNoRows {
			return nil, err
		} else {
			return nil, nil
		}
	}

	return &url, nil
}

func (pg *Client) DeleteURL(ctx context.Context, shortURLKey string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*defaultTimeout)
	defer cancel()

	query := `DELETE FROM urls WHERE short_url_key = $1 RETURNING short_url_key`
	row := pg.pool.QueryRow(ctx, query, shortURLKey)

	err := row.Scan(&shortURLKey)
	if err != nil {
		if err != pgx.ErrNoRows {
			return false, err
		} else {
			return false, nil
		}
	}

	return true, nil
}

func (pg *Client) IncrementClicks(ctx context.Context, shortURLKey string) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*defaultTimeout)
	defer cancel()

	query := `UPDATE urls SET clicks = clicks + 1 WHERE short_url_key = $1`
	_, err := pg.pool.Exec(ctx, query, shortURLKey)

	return err
}

func (pg *Client) Ping() error {
	return pg.pool.Ping(context.Background())
}

func (pg *Client) FlushAll() error {
	_, err := pg.pool.Exec(context.Background(), `TRUNCATE urls`)
	return err
}

func (pg *Client) Close() {
	pg.pool.Close()
}
