package bbox

import (
	"context"
	"testing"

	"github.com/google/uuid"

	"github.com/openmarketplaceengine/openmarketplaceengine/internal/service/tollgate"

	"github.com/openmarketplaceengine/openmarketplaceengine/cfg"
	redisClient "github.com/openmarketplaceengine/openmarketplaceengine/redis/client"
	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

func TestTollgate(t *testing.T) {
	err := cfg.Load()
	require.NoError(t, err)

	client := redisClient.NewStoreClient()
	require.NotNil(t, client)

	storage := NewStorage(client)
	ctx := context.Background()

	t.Run("testCrossing", func(t *testing.T) {
		testCrossing(ctx, t, storage)
	})
}

func testCrossing(ctx context.Context, t *testing.T, s Storage) {
	required := 4
	bBoxes := []*tollgate.BBox{{
		LonMin: 1,
		LatMin: 1,
		LonMax: 5,
		LatMax: 5,
	}, {
		LonMin: 5,
		LatMin: 5,
		LonMax: 10,
		LatMax: 10,
	}, {
		LonMin: 10,
		LatMin: 10,
		LonMax: 15,
		LatMax: 15,
	}, {
		LonMin: 15,
		LatMin: 15,
		LonMax: 20,
		LatMax: 20,
	}, {
		LonMin: 20,
		LatMin: 20,
		LonMax: 35,
		LatMax: 35,
	}}

	loc0 := tollgate.Location{}

	loc1 := tollgate.Location{
		Lon: 2,
		Lat: 2,
	}
	loc2 := tollgate.Location{
		Lon: 7,
		Lat: 7,
	}
	loc3 := tollgate.Location{
		Lon: 13,
		Lat: 13,
	}
	//noisy GPS at 4, skipping
	loc5 := tollgate.Location{
		Lon: 33,
		Lat: 33,
	}
	subjectID := uuid.NewString()
	tollgateID := "toll-1"
	c1, _ := DetectCrossing(ctx, s, tollgateID, bBoxes, required, &tollgate.Movement{SubjectID: subjectID, From: &loc0, To: &loc1})
	require.Nil(t, c1)

	c2, _ := DetectCrossing(ctx, s, tollgateID, bBoxes, required, &tollgate.Movement{SubjectID: subjectID, From: &loc0, To: &loc2})
	require.Nil(t, c2)

	c3, _ := DetectCrossing(ctx, s, tollgateID, bBoxes, required, &tollgate.Movement{SubjectID: subjectID, From: &loc0, To: &loc3})
	require.Nil(t, c3)

	c4, _ := DetectCrossing(ctx, s, tollgateID, bBoxes, required, &tollgate.Movement{SubjectID: subjectID, From: &loc0, To: &loc0})
	require.Nil(t, c4)

	visits, err := s.Visits(ctx, tollgateID, subjectID, 5)
	require.NoError(t, err)

	assert.Equal(t, []int64{1, 1, 1, 0, 0}, visits)

	c5, _ := DetectCrossing(ctx, s, tollgateID, bBoxes, required, &tollgate.Movement{SubjectID: subjectID, From: &loc0, To: &loc5})
	require.NotNil(t, c5)

	visits, err = s.Visits(ctx, tollgateID, subjectID, 5)
	require.NoError(t, err)

	assert.Equal(t, []int64{0, 0, 0, 0, 0}, visits)
}
