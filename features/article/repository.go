package article

import (
	"context"

	"github.com/mochadwi/go-article/models"
)

type ArticleRepository interface {
	Create(ctx context.Context, a *models.Article) (int64, error)
	GetAll(ctx context.Context, cursor string, num int64) ([]*models.Article, error)
	GetByID(ctx context.Context, id int64) (*models.Article, error)
	GetByTitle(ctx context.Context, title string) (*models.Article, error)
	Update(ctx context.Context, article *models.Article) (*models.Article, error)
	Delete(ctx context.Context, id int64) (bool, error)
}
