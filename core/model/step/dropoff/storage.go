package dropoff

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/openmarketplaceengine/openmarketplaceengine/redis/client"
	"github.com/openmarketplaceengine/openmarketplaceengine/redis/marshalutils"
)

type Storage struct {
	client     *redis.Client
	expiration time.Duration
}

var storage = newStorage(3 * time.Hour)

func newStorage(expiration time.Duration) *Storage {
	return &Storage{
		client:     client.NewStoreClient(),
		expiration: expiration,
	}
}

func key(driverID string) string {
	return fmt.Sprintf("dropoff-%s", driverID)
}

func (s *Storage) Store(ctx context.Context, dropOff DropOff) error {
	m, err := marshalutils.StructToMap(dropOff)

	if err != nil {
		return fmt.Errorf("store StructToMap error: %s", err)
	}

	key := key(dropOff.DriverID)

	err = s.client.HSet(ctx, key, m).Err()
	if err != nil {
		return fmt.Errorf("store HSet error: %s", err)
	}

	s.client.Expire(ctx, key, s.expiration)

	return nil
}

func (s *Storage) Retrieve(ctx context.Context, driverID string) (dropOff DropOff, err error) {
	key := key(driverID)
	m, err := s.client.HGetAll(ctx, key).Result()
	if err != nil {
		err = fmt.Errorf("retrieve HGetAll error: %s", err)
		return
	}

	if len(m) == 0 {
		err = fmt.Errorf("dropOff not found by driverID %q", driverID)
		return
	}

	err = marshalutils.MapToStruct(m, &dropOff)
	if err != nil {
		err = fmt.Errorf("retrieve MapToStruct error: %s", err)
		return
	}
	return
}
