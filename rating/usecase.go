package rating

import (
	"context"

	model "github.com/mochadwi/go-article/models"
)

type RatingUsecase interface {
	Create(context.Context, *model.Rating) (*model.Rating, error)
	GetByID(ctx context.Context, lessonId int64) (*model.Rating, error)
}
