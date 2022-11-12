package redis_test

import (
	"context"
	"cybercoin/dal/redis"
	"fmt"
	"testing"
)

func TestNewClient(t *testing.T) {
	ctx := context.Background()
	client := redis.NewClient(ctx)
	client.Set(ctx, "a", "cd", 0)
	get := client.Get(ctx, "a")
	fmt.Println(get.Val())
	fmt.Println(get.Err())
}
