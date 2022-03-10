package location

import (
	"context"
	"encoding/json"
	"fmt"

	locationV1beta1 "github.com/openmarketplaceengine/openmarketplaceengine/internal/omeapi/location/v1beta1"
	"github.com/openmarketplaceengine/openmarketplaceengine/internal/service/location/storage"

	"github.com/go-redis/redis/v8"
	"github.com/openmarketplaceengine/openmarketplaceengine/redis/publisher"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Controller struct {
	locationV1beta1.UnimplementedLocationServiceServer
	store   *storage.Storage
	pub     publisher.Publisher
	areaKey string
}

func New(storeClient *redis.Client, pubClient *redis.Client, areaKey string) *Controller {
	return &Controller{
		store:   storage.New(storeClient),
		pub:     publisher.New(pubClient),
		areaKey: areaKey,
	}
}

func (c *Controller) UpdateLocation(ctx context.Context, request *locationV1beta1.UpdateLocationRequest) (*locationV1beta1.UpdateLocationResponse, error) {
	err := c.store.Update(ctx, c.areaKey, &storage.Location{
		WorkerID:  request.WorkerId,
		Longitude: request.Longitude,
		Latitude:  request.Latitude,
	})
	if err != nil {
		return nil, err
	}

	channel := fmt.Sprintf("location-%s", request.WorkerId)

	bytes, err := json.Marshal(storage.Location{
		WorkerID:  request.WorkerId,
		Longitude: request.Longitude,
		Latitude:  request.Latitude,
	})
	if err != nil {
		return nil, err
	}
	payload := string(bytes)
	err = c.pub.Publish(ctx, channel, payload)

	if err != nil {
		return nil, err
	}

	return &locationV1beta1.UpdateLocationResponse{
		WorkerId: request.WorkerId,
	}, nil
}
func (c *Controller) QueryLocation(ctx context.Context, request *locationV1beta1.QueryLocationRequest) (*locationV1beta1.QueryLocationResponse, error) {
	l := c.store.LastLocation(ctx, c.areaKey, request.WorkerId)
	if l != nil {
		return &locationV1beta1.QueryLocationResponse{
			WorkerId:  l.WorkerID,
			Longitude: l.Longitude,
			Latitude:  l.Latitude,
			LastSeenTime: &timestamppb.Timestamp{
				Seconds: l.LastSeenTime.Unix(),
				Nanos:   0,
			},
		}, nil
	}
	return nil, fmt.Errorf("location not found for WorkerId=%s", request.WorkerId)
}
