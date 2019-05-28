package usecase

import (
	"context"
	"time"

	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/mochadwi/go-article/models"
	"github.com/mochadwi/go-article/rating"
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

func (a *ratingUsecase) GetAll(c context.Context, cursor string, num int64) (*[]*models.Rating, string, error) {
	if num == 0 {
		num = 10
	}

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	listRating, err := a.ratingRepos.GetAll(ctx, cursor, num)
	if err != nil {
		return nil, "", err
	}

	return listRating, cursor, nil
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

func (a *ratingUsecase) Update(c context.Context, ar *models.Rating) (*models.Rating, error) {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	_, err := a.GetByID(ctx, ar.ID)
	if err != nil {
		return nil, models.NOT_FOUND_ERROR
	}

	existedRating, _ := a.GetByTitle(ctx, ar.Title)
	if existedRating != nil {
		return nil, models.CONFLIT_ERROR
	}

	ar.UpdatedAt = time.Now()
	return a.ratingRepos.Update(ctx, ar)
}

func (a *ratingUsecase) GetByTitle(c context.Context, title string) (*models.Rating, error) {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	res, err := a.ratingRepos.GetByTitle(ctx, title)
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

func (a *ratingUsecase) Create(c context.Context, m *models.Rating) (*models.Rating, error) {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	existedRating, _ := a.GetByTitle(ctx, m.Title)
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

func (a *ratingUsecase) Delete(c context.Context, id int64) (bool, error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	existedRating, _ := a.ratingRepos.GetByID(ctx, id)
	if existedRating == nil {
		return false, models.NOT_FOUND_ERROR
	}
	return a.ratingRepos.Delete(ctx, id)
}
