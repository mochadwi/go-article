package rating

import (
	"context"

	"github.com/mochadwi/go-article/models"
)

type RatingRepository interface {
	Create(ctx context.Context, a *models.Rating) (int64, error)
	GetByID(ctx context.Context, lessonId int64) (*models.Rating, error)
}
