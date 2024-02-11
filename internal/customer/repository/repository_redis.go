package customerrepo

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RepositoryRedis struct {
	client *redis.Client
}

func NewRedis(client *redis.Client) *RepositoryRedis {
	return &RepositoryRedis{client: client}
}

func (r RepositoryRedis) SetRefreshToken(ctx context.Context, token string, expireIn time.Duration) (err error) {
	err = r.client.Set(ctx, token, "", expireIn).Err()
	if err != nil {
		err = fmt.Errorf("customer.repository_redis.SetRefreshToken: failed to set refresh token: %w", err)
		return
	}

	return
}
