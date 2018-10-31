package repository

import (
	"context"
	"github.com/sirupsen/logrus"

	"github.com/mochadwi/go-article/article"
	"github.com/mochadwi/go-article/models"
	"github.com/jinzhu/gorm"
	"strconv"
	"fmt"
)

type gormsqlArticleRepository struct {
	Conn *gorm.DB
}

func NewGormsqlArticleRepository(Conn *gorm.DB) article.ArticleRepository {

	return &gormsqlArticleRepository{Conn}
}

func (m *gormsqlArticleRepository) GetAll(ctx context.Context, cursor string, num int64) (*[]*models.Article, error) {
	//m.Conn.Begin()

	var result []*models.Article
	var cursorInt, _ = strconv.ParseInt(cursor, 10, 64)
	errQuery := models.NewArticleQuerySet(m.Conn).IDGt(cursorInt).Limit(int(num)).All(&result)

	if errQuery != nil {
		logrus.Error(errQuery)
		return nil, errQuery
	}
	//defer m.Conn.Close()

	return &result, nil

}
func (m *gormsqlArticleRepository) GetByID(ctx context.Context, id int64) (*models.Article, error) {

	var ac models.Article
	if errQuery := models.NewArticleQuerySet(m.Conn).IDEq(id).One(&ac); errQuery != nil {
		fmt.Print("[gorm error] GetById: ")
		fmt.Println(errQuery)
		logrus.Error(errQuery)
		return nil, errQuery
	}

	fmt.Print("[gorm success] GetById: ")
	fmt.Println(ac)
	return &ac, nil
}

func (m *gormsqlArticleRepository) GetByTitle(ctx context.Context, title string) (*models.Article, error) {

	//m.Conn.Begin()
	var ac models.Article
	if errQuery := models.NewArticleQuerySet(m.Conn).TitleEq(title).One(&ac); errQuery != nil {
		if &ac == nil {
			return nil, models.NOT_FOUND_ERROR
		}

		logrus.Error(errQuery)
		return nil, errQuery
	}

	//fmt.Print("Repo: ")
	//fmt.Println(ac)
	//defer m.Conn.Close()
	return &ac, nil
}

func (m *gormsqlArticleRepository) Create(ctx context.Context, a *models.Article) (int64, error) {
	// Create
	if errQuery := a.Create(m.Conn); errQuery != nil {
		logrus.Debug("Created At: ", a.CreatedAt)
		return 0, errQuery
	} // end Create

	return a.ID, nil
}

func (m *gormsqlArticleRepository) Delete(ctx context.Context, id int64) (bool, error) {

	// This method doesn't delete the data, for backup purpose
	// it only update the deleted_at fields
	if err := models.NewArticleQuerySet(m.Conn).IDEq(id).Delete(); err != nil {
		return false, err
	} // end Delete

	return true, nil
}

func (m *gormsqlArticleRepository) Update(ctx context.Context, ar *models.Article) (*models.Article, error) {

	//fmt.Print("Gorm: ")
	//fmt.Println(ar)

	// Update
	if err :=
		models.NewArticleQuerySet(m.Conn).
			IDEq(ar.ID).
			GetUpdater().
			SetContent(ar.Content).
			SetThumbnail(ar.Thumbnail).
			SetTitle(ar.Title).
			Update();
		err != nil {
		return nil, err
	} // end Update

	//fmt.Print("Gorm: ")
	//fmt.Println(ar)

	return ar, nil
}
