package rating

import (
	"context"
	"errors"

	"github.com/thernande/movie-micro/rating/internal/repository"
	"github.com/thernande/movie-micro/rating/pkg/model"
)

// ErrNotFound is returned when no ratings are found for a
// record.
var ErrNotFound = errors.New("ratings not found for a record")

type ratingRepository interface {
	Get(ctx context.Context, recordID model.RecordID, recordType model.RecordType) ([]model.Rating, error)
	Put(ctx context.Context, recordID model.RecordID, recordType model.RecordType, rating *model.Rating) error
}

// Controller defines a rating service controller.
type Controller struct {
	repo ratingRepository
}

// New creates a rating service controller.
func New(repo ratingRepository) *Controller {
	return &Controller{repo}
}

// GetAggregatedRating retrieves the aggregated rating for a given record ID and type.
//
// Parameters:
//   - ctx: the context.Context for the request.
//   - recordID: the ID of the record.
//   - recordType: the type of the record.
//
// Returns:
//   - float64: the aggregated rating.
//   - error: any error that occurred during the retrieval.
func (c *Controller) GetAggregatedRating(ctx context.Context, recordID model.RecordID, recordType model.RecordType) (float64, error) {
	ratings, err := c.repo.Get(ctx, recordID, recordType)
	if err != nil && err == repository.ErrNotFound {
		return 0, ErrNotFound
	} else if err != nil {
		return 0, err
	}
	sum := float64(0)
	for _, r := range ratings {
		sum += float64(r.Value)
	}
	return sum / float64(len(ratings)), nil
}

// PutRating writes a rating for a given record.
func (c *Controller) PutRating(ctx context.Context, recordID model.RecordID, recordType model.RecordType, rating *model.Rating) error {
	return c.repo.Put(ctx, recordID, recordType, rating)
}
