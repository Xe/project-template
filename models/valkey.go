package models

import (
	valkey "github.com/redis/go-redis/v9"
)

func ConnectValkey(valkeyURL string) (*valkey.Client, error) {
	opts, err := valkey.ParseURL(valkeyURL)
	if err != nil {
		return nil, err
	}

	return valkey.NewClient(opts), nil
}
