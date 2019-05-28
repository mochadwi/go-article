package rating

import (
	"context"

	model "github.com/mochadwi/go-article/models"
)

type RatingUsecase interface {
	Create(context.Context, *model.Rating) (*model.Rating, error)
	GetAll(ctx context.Context, cursor string, num int64) (*[]*model.Rating, string, error)
	GetByID(ctx context.Context, id int64) (*model.Rating, error)
	GetByTitle(ctx context.Context, title string) (*model.Rating, error)
	Update(ctx context.Context, ar *model.Rating) (*model.Rating, error)
	Delete(ctx context.Context, id int64) (bool, error)
}
