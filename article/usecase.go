package article

import (
	"context"

	model "github.com/mochadwi/go-article/models"
)

type ArticleUsecase interface {
	Create(context.Context, *model.Article) (*model.Article, error)
	GetAll(ctx context.Context, cursor string, num int64) (*[]model.Article, string, error)
	GetByID(ctx context.Context, id int64) (*model.Article, error)
	GetByTitle(ctx context.Context, title string) (*model.Article, error)
	Update(ctx context.Context, ar *model.Article) (*model.Article, error)
	Delete(ctx context.Context, id int64) (bool, error)
}
