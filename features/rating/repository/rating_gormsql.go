package repository

import (
	"context"
	"github.com/sirupsen/logrus"

	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/mochadwi/go-article/features/rating"
	"github.com/mochadwi/go-article/models"
)

type gormsqlRatingRepository struct {
	Conn *gorm.DB
}

func NewGormsqlRatingRepository(Conn *gorm.DB) rating.RatingRepository {

	return &gormsqlRatingRepository{Conn}
}

func (m *gormsqlRatingRepository) GetByID(ctx context.Context, lessonId int64) (*models.Rating, error) {

	var ac models.Rating
	if errQuery := models.NewRatingQuerySet(m.Conn).LessonIDEq(lessonId).One(&ac); errQuery != nil {
		fmt.Print("[repo error] GetById: ")
		fmt.Println(errQuery)
		logrus.Error(errQuery)
		return nil, errQuery
	}

	fmt.Print("[repo success] GetById: ")
	fmt.Println(ac)
	return &ac, nil
}

func (m *gormsqlRatingRepository) Create(ctx context.Context, a *models.Rating) (int64, error) {
	// Create
	if errQuery := a.Create(m.Conn); errQuery != nil {
		logrus.Debug("Created At: ", a.CreatedAt)
		return 0, errQuery
	} // end Create

	return a.ID, nil
}
