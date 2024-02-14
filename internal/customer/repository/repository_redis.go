package customerrepo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/arfan21/synapsis_id/internal/entity"
	"github.com/arfan21/synapsis_id/pkg/constant"
	"github.com/redis/go-redis/v9"
)

type RepositoryRedis struct {
	client *redis.Client
}

func NewRedis(client *redis.Client) *RepositoryRedis {
	return &RepositoryRedis{client: client}
}

func (r RepositoryRedis) SetRefreshToken(ctx context.Context, token string, expireIn time.Duration, payload entity.CustomerRefreshToken) (err error) {
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		err = fmt.Errorf("customer.repository_redis.SetRefreshToken: failed to marshal payload: %w", err)
		return
	}

	err = r.client.Set(ctx, token, string(payloadJson), expireIn).Err()
	if err != nil {
		err = fmt.Errorf("customer.repository_redis.SetRefreshToken: failed to set refresh token: %w", err)
		return
	}

	return
}

func (r RepositoryRedis) IsRefreshTokenExist(ctx context.Context, token string) (payload entity.CustomerRefreshToken, err error) {
	resultStr, err := r.client.Get(ctx, token).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			err = constant.ErrUnauthorizedAccess
		}
		err = fmt.Errorf("customer.repository_redis.IsRefreshTokenExist: failed to get refresh token: %w", err)
		return
	}

	err = json.Unmarshal([]byte(resultStr), &payload)
	if err != nil {
		err = fmt.Errorf("customer.repository_redis.IsRefreshTokenExist: failed to unmarshal payload: %w", err)
		return
	}

	return
}

func (r RepositoryRedis) DeleteRefreshToken(ctx context.Context, token string) (err error) {
	err = r.client.Del(ctx, token).Err()
	if err != nil && !errors.Is(err, redis.Nil) {
		err = fmt.Errorf("customer.repository_redis.DeleteRefreshToken: failed to delete refresh token: %w", err)
		return
	}

	return
}
