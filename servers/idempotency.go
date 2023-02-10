package servers

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/emil-petras/project-idempotency-service/db"
	"github.com/emil-petras/project-idempotency-service/utils"
	idempotencyProto "github.com/emil-petras/project-proto/idempotency"
	"github.com/go-redis/redis"
)

type IdempotencyServer struct {
	idempotencyProto.UnimplementedIdempotencyServiceServer
}

func (i *IdempotencyServer) Check(ctx context.Context, in *idempotencyProto.Request) (*idempotencyProto.Response, error) {
	key, err := utils.Hash([]byte(in.Value))
	if err != nil {
		return nil, fmt.Errorf("check failed: %w", err)
	}

	_, err = db.Client.Get(key).Result()
	response := idempotencyProto.Response{}
	if err == redis.Nil {
		minutes, err := strconv.Atoi(os.Getenv("REDIS_EXPIRATION"))
		if err != nil {
			return nil, fmt.Errorf("reading REDIS_EXPIRATION failed: %w", err)
		}

		expiration := time.Minute * time.Duration(minutes)
		db.Client.Set(key, in.Value, expiration)

		response.Exists = false
		return &response, nil
	} else if err != nil {
		return nil, fmt.Errorf("redis get failed: %w", err)
	} else {
		response.Exists = true
		return &response, nil
	}
}
