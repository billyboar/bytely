package schema

import (
	"errors"
	"net/url"
	"time"
)

type URL struct {
	CreatedAt time.Time `json:"created_at"`

	// unique key for the URL
	// end result will be like this:
	// http://localhost/<short_url_key>
	ShortURLKey string `json:"short_url_key"`

	// original URL user tried to shorten
	OriginalURL string `json:"original_url"`

	// number of times the URL has been clicked
	Clicks int `json:"clicks"`
}

var (
	ErrOriginalURLMissing  = errors.New("original_url is missing")
	ErrOriginalURLTooLarge = errors.New("original_url is too large")
)

// Validate checks if the URL is valid.
// Avoided using 3rd party library for this scenario
// as it's not a requirement for this project.
func (u URL) Validate() error {
	if len(u.OriginalURL) == 0 {
		return ErrOriginalURLMissing
	}

	if len(u.OriginalURL) > 4112 {
		return ErrOriginalURLTooLarge
	}

	_, err := url.ParseRequestURI(u.OriginalURL)
	if err != nil {
		return err
	}

	return nil
}
