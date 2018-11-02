package usecase

import (
	"context"
	"time"

	"github.com/mochadwi/go-article/models"
	"github.com/mochadwi/go-article/article"
	"fmt"
	"github.com/jinzhu/gorm"
)

type articleUsecase struct {
	articleRepos   article.ArticleRepository
	contextTimeout time.Duration
}

func NewArticleUsecase(a article.ArticleRepository, timeout time.Duration) article.ArticleUsecase {
	return &articleUsecase{
		articleRepos:   a,
		contextTimeout: timeout,
	}
}

func (a *articleUsecase) GetAll(c context.Context, cursor string, num int64) (*[]*models.Article, string, error) {
	if num == 0 {
		num = 10
	}

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	listArticle, err := a.articleRepos.GetAll(ctx, cursor, num)
	if err != nil {
		return nil, "", err
	}

	return listArticle, cursor, nil
}

func (a *articleUsecase) GetByID(c context.Context, id int64) (*models.Article, error) {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, err := a.articleRepos.GetByID(ctx, id)
	if err != nil {
		fmt.Print("[usecase error] GetById: ")
		fmt.Println(res)
		return nil, err
	}

	fmt.Print("[usecase success] GetById: ")
	fmt.Println(res)
	return res, nil
}

func (a *articleUsecase) Update(c context.Context, ar *models.Article) (*models.Article, error) {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	_, err := a.GetByID(ctx, ar.ID)
	if err != nil {
		return nil, models.NOT_FOUND_ERROR
	}

	existedArticle, _ := a.GetByTitle(ctx, ar.Title)
	if existedArticle != nil {
		return nil, models.CONFLIT_ERROR
	}

	ar.UpdatedAt = time.Now()
	return a.articleRepos.Update(ctx, ar)
}

func (a *articleUsecase) GetByTitle(c context.Context, title string) (*models.Article, error) {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	res, err := a.articleRepos.GetByTitle(ctx, title)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, models.NOT_FOUND_ERROR
		}

		return nil, err
	}

	//fmt.Print("Ucase: ")
	//fmt.Println(res)

	return res, nil
}

func (a *articleUsecase) Create(c context.Context, m *models.Article) (*models.Article, error) {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	existedArticle, _ := a.GetByTitle(ctx, m.Title)
	if existedArticle != nil {
		return nil, models.CONFLIT_ERROR
	}

	id, err := a.articleRepos.Create(ctx, m)
	if err != nil {
		return nil, err
	}

	m.ID = id
	return m, nil
}

func (a *articleUsecase) Delete(c context.Context, id int64) (bool, error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	existedArticle, _ := a.articleRepos.GetByID(ctx, id)
	if existedArticle == nil {
		return false, models.NOT_FOUND_ERROR
	}
	return a.articleRepos.Delete(ctx, id)
}
