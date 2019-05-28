package repository

import (
	"context"
	"github.com/sirupsen/logrus"

	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/mochadwi/go-article/models"
	"github.com/mochadwi/go-article/rating"
	"strconv"
)

type gormsqlRatingRepository struct {
	Conn *gorm.DB
}

func NewGormsqlRatingRepository(Conn *gorm.DB) rating.RatingRepository {

	return &gormsqlRatingRepository{Conn}
}

func (m *gormsqlRatingRepository) GetAll(ctx context.Context, cursor string, num int64) (*[]*models.Rating, error) {
	//m.Conn.Begin()

	var result []*models.Rating
	var cursorInt, _ = strconv.ParseInt(cursor, 10, 64)
	errQuery := models.NewRatingQuerySet(m.Conn).IDGt(cursorInt).Limit(int(num)).All(&result)

	if errQuery != nil {
		logrus.Error(errQuery)
		return nil, errQuery
	}
	//defer m.Conn.Close()

	return &result, nil

}
func (m *gormsqlRatingRepository) GetByID(ctx context.Context, id int64) (*models.Rating, error) {

	var ac models.Rating
	if errQuery := models.NewRatingQuerySet(m.Conn).IDEq(id).One(&ac); errQuery != nil {
		fmt.Print("[repo error] GetById: ")
		fmt.Println(errQuery)
		logrus.Error(errQuery)
		return nil, errQuery
	}

	fmt.Print("[repo success] GetById: ")
	fmt.Println(ac)
	return &ac, nil
}

func (m *gormsqlRatingRepository) GetByTitle(ctx context.Context, title string) (*models.Rating, error) {

	//m.Conn.Begin()
	var ac models.Rating
	if errQuery := models.
		NewRatingQuerySet(m.Conn).
		TitleEq(title).
		One(&ac); errQuery != nil {
		//fmt.Print("[repo error]: ")
		//fmt.Println(ac)

		logrus.Error(errQuery)
		//fmt.Print("[repo error]: ")
		//fmt.Println(ac)
		return nil, errQuery
	}

	//defer m.Conn.Close()
	fmt.Print("[repo error]: ")
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

func (m *gormsqlRatingRepository) Delete(ctx context.Context, id int64) (bool, error) {

	// This method doesn't delete the data, for backup purpose
	// it only update the deleted_at fields
	if err := models.NewRatingQuerySet(m.Conn).IDEq(id).Delete(); err != nil {
		return false, err
	} // end Delete

	return true, nil
}

func (m *gormsqlRatingRepository) Update(ctx context.Context, ar *models.Rating) (*models.Rating, error) {

	//fmt.Print("Gorm: ")
	//fmt.Println(ar)

	// Update
	if err :=
		models.NewRatingQuerySet(m.Conn).
			IDEq(ar.ID).
			GetUpdater().
			SetContent(ar.Content).
			SetThumbnail(ar.Thumbnail).
			SetTitle(ar.Title).
			Update(); err != nil {
		return nil, err
	} // end Update

	//fmt.Print("Gorm: ")
	//fmt.Println(ar)

	return ar, nil
}
