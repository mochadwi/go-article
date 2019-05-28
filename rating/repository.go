package rating

import (
	"context"

	"github.com/mochadwi/go-article/models"
)

type RatingRepository interface {
	Create(ctx context.Context, a *models.Rating) (int64, error)
	GetAll(ctx context.Context, cursor string, num int64) (*[]*models.Rating, error)
	GetByID(ctx context.Context, id int64) (*models.Rating, error)
	GetByTitle(ctx context.Context, title string) (*models.Rating, error)
	Update(ctx context.Context, rating *models.Rating) (*models.Rating, error)
	Delete(ctx context.Context, id int64) (bool, error)
}
