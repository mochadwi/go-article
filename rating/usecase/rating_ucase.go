package usecase

import (
	"context"
	"time"

	"fmt"
	"github.com/mochadwi/go-article/features/rating"
	"github.com/mochadwi/go-article/models"
)

type ratingUsecase struct {
	ratingRepos    rating.RatingRepository
	contextTimeout time.Duration
}

func NewRatingUsecase(a rating.RatingRepository, timeout time.Duration) rating.RatingUsecase {
	return &ratingUsecase{
		ratingRepos:    a,
		contextTimeout: timeout,
	}
}

func (a *ratingUsecase) GetByID(c context.Context, id int64) (*models.Rating, error) {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, err := a.ratingRepos.GetByID(ctx, id)
	if err != nil {
		fmt.Print("[usecase error] GetById: ")
		fmt.Println(res)
		return nil, err
	}

	fmt.Print("[usecase success] GetById: ")
	fmt.Println(res)
	return res, nil
}

func (a *ratingUsecase) Create(c context.Context, m *models.Rating) (*models.Rating, error) {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	existedRating, _ := a.GetByID(ctx, m.LessonID)
	if existedRating != nil {
		return nil, models.CONFLIT_ERROR
	}

	id, err := a.ratingRepos.Create(ctx, m)
	if err != nil {
		return nil, err
	}

	m.ID = id
	return m, nil
}
