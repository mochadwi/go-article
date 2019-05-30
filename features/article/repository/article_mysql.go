package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/mochadwi/go-article/features/article"
	"github.com/mochadwi/go-article/models"
)

type mysqlArticleRepository struct {
	Conn *sql.DB
}

func NewMysqlArticleRepository(Conn *sql.DB) article.ArticleRepository {

	return &mysqlArticleRepository{Conn}
}

func (m *mysqlArticleRepository) fetch(ctx context.Context, query string, args ...interface{}) (*[]*models.Article, error) {

	rows, err := m.Conn.QueryContext(ctx, query, args...)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer rows.Close()
	var result *[]*models.Article
	for rows.Next() {
		t := new(models.Article)
		err = rows.Scan(
			&t.ID,
			&t.Title,
			&t.Content,
			&t.Thumbnail,
			&t.UpdatedAt,
			&t.CreatedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		*result = append(*result, t)
	}

	return result, nil
}

func (m *mysqlArticleRepository) GetAll(ctx context.Context, cursor string, num int64) (*[]*models.Article, error) {

	query := `SELECT id,title,content, author_id, updated_at, created_at
  						FROM article WHERE ID > ? LIMIT ?`

	return m.fetch(ctx, query, cursor, num)

}
func (m *mysqlArticleRepository) GetByID(ctx context.Context, id int64) (*models.Article, error) {
	query := `SELECT id,title,content, author_id, updated_at, created_at
  						FROM article WHERE ID = ?`

	list, err := m.fetch(ctx, query, id)
	if err != nil {
		return nil, err
	}

	a := models.Article{}
	if len(*list) > 0 {
		a = *(*list)[0]
	} else {
		return nil, models.NOT_FOUND_ERROR
	}

	return &a, nil
}

func (m *mysqlArticleRepository) GetByTitle(ctx context.Context, title string) (*models.Article, error) {
	query := `SELECT id,title,content, author_id, updated_at, created_at
  						FROM article WHERE title = ?`

	list, err := m.fetch(ctx, query, title)
	if err != nil {
		return nil, err
	}

	a := models.Article{}
	if len(*list) > 0 {
		a = *(*list)[0]
	} else {
		return nil, models.NOT_FOUND_ERROR
	}
	return &a, nil
}

func (m *mysqlArticleRepository) Create(ctx context.Context, a *models.Article) (int64, error) {

	query := `INSERT  article SET title=? , content=? , author_id=?, updated_at=? , created_at=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {

		return 0, err
	}

	logrus.Debug("Created At: ", a.CreatedAt)
	res, err := stmt.ExecContext(ctx, a.Title, a.UpdatedAt, a.CreatedAt)
	if err != nil {

		return 0, err
	}
	return res.LastInsertId()
}

func (m *mysqlArticleRepository) Delete(ctx context.Context, id int64) (bool, error) {
	query := "DELETE FROM article WHERE id = ?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return false, err
	}
	res, err := stmt.ExecContext(ctx, id)
	if err != nil {

		return false, err
	}
	rowsAfected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}
	if rowsAfected != 1 {
		err = fmt.Errorf("Weird  Behaviour. Total Affected: %d", rowsAfected)
		logrus.Error(err)
		return false, err
	}

	return true, nil
}
func (m *mysqlArticleRepository) Update(ctx context.Context, ar *models.Article) (*models.Article, error) {
	query := `UPDATE article set title=?, content=?, author_id=?, updated_at=? WHERE ID = ?`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, nil
	}

	res, err := stmt.ExecContext(ctx, ar.Title, ar.Content, ar.UpdatedAt, ar.ID)
	if err != nil {
		return nil, err
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if affect != 1 {
		err = fmt.Errorf("Weird  Behaviour. Total Affected: %d", affect)
		logrus.Error(err)
		return nil, err
	}

	return ar, nil
}
