package http_echo

import (
	"github.com/jackwhelpton/fasthttp-routing"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/mochadwi/go-article/models"

	"fmt"
	articleUcase "github.com/mochadwi/go-article/article"
	"gopkg.in/go-playground/validator.v9"
	"time"
)

type HttpArticleHandler struct {
	AUsecase articleUcase.ArticleUsecase
}

func (a *HttpArticleHandler) GetAll(c *routing.Context) error {

	numS := c.Query("num")
	num, _ := strconv.Atoi(numS)
	cursor := c.Query("cursor")
	ctx := c.RequestCtx
	if ctx == nil {
		ctx = c.RequestCtx
	}

	listAr, nextCursor, err := a.AUsecase.GetAll(ctx, cursor, int64(num))

	//reqId := uuid.NewV4().String()
	var response = &models.BaseResponse{
		RequestID: "",
		Now:       time.Now().Unix(),
	}

	if err != nil {
		response.Code = getStatusCode(err)
		response.Message = err.Error()
		response.Data = listAr

		c.Response.SetStatusCode(response.Code)
		return c.Write(response)
	}

	if len(*listAr) > 0 {
		response.Message = models.DATA_AVAILABLE_SUCCESS
	} else {
		response.Message = models.DATA_EMPTY_SUCCESS
	}

	response.Code = http.StatusOK
	response.Data = listAr

	c.Response.Header.Set(`X-Cursor`, nextCursor)
	c.Response.SetStatusCode(response.Code)
	return c.Write(response)
}

func (a *HttpArticleHandler) GetByTitle(c *routing.Context) error {

	title := c.Param("title")
	//title2 := c.QueryParam("title")

	ctx := c.RequestCtx
	if ctx == nil {
		ctx = c.RequestCtx
	}

	art, err := a.AUsecase.GetByTitle(ctx, title)

	//reqId := uuid.NewV4().String()
	var response = &models.BaseResponse{
		RequestID: "",
		Now:       time.Now().Unix(),
	}

	if err != nil {
		response.Code = getStatusCode(err)
		response.Message = err.Error()
		response.Data = art

		c.Response.SetStatusCode(response.Code)
		return c.Write(response)
	}

	response.Code = http.StatusOK
	response.Message = models.DATA_AVAILABLE_SUCCESS
	response.Data = art

	fmt.Print("Handler: ")
	fmt.Println(art)

	c.Response.SetStatusCode(response.Code)
	return c.Write(response)
}

func (a *HttpArticleHandler) GetByID(c *routing.Context) error {

	idP, err := strconv.Atoi(c.Param("id"))
	id := int64(idP)

	ctx := c.RequestCtx
	if ctx == nil {
		ctx = c.RequestCtx
	}

	art, err := a.AUsecase.GetByID(ctx, id)

	if err != nil {
		c.Response.SetStatusCode(getStatusCode(err))
		return c.Write(models.BaseResponse{Message: err.Error()})
	}

	return c.Write(art)
}

func isRequestValid(m *models.Article) (bool, error) {

	validate := validator.New()

	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (a *HttpArticleHandler) Create(c *routing.Context) error {
	var article models.Article
	err := c.Read(&article)

	var response = &models.BaseResponse{
		RequestID: "",
		Now:       time.Now().Unix(),
	}

	if err != nil {
		response.Code = http.StatusUnprocessableEntity
		response.Message = string(err.Error())
		response.Data = article

		c.Response.SetStatusCode(response.Code)
		return c.Write(response) // todo write custom header instead
	}

	if ok, err := isRequestValid(&article); !ok {
		response.Code = http.StatusBadRequest
		response.Message = string(err.Error())
		response.Data = article

		c.Response.SetStatusCode(response.Code)
		return c.Write(response)
	}
	ctx := c.RequestCtx
	if ctx == nil {
		ctx = c.RequestCtx
	}

	ar, err := a.AUsecase.Create(ctx, &article)

	if err != nil {
		response.Code = getStatusCode(err)
		response.Message = string(err.Error())
		response.Data = article

		c.Response.SetStatusCode(response.Code)
		return c.Write(response)
	}

	response.Code = http.StatusCreated
	response.Message = models.DATA_CREATED_SUCCESS
	response.Data = ar

	c.Response.SetStatusCode(response.Code)
	return c.Write(response)
}

func (a *HttpArticleHandler) Update(c *routing.Context) error {

	fmt.Print("[Handler] Update id: ")
	fmt.Println(c.Param("id"))

	var article models.Article
	var id, err = strconv.Atoi(c.Param("id"))

	var response = &models.BaseResponse{
		RequestID: "",
		Now:       time.Now().Unix(),
	}

	if err != nil {
		response.Code = http.StatusUnprocessableEntity
		response.Message = string(err.Error())
		response.Data = id

		c.Response.SetStatusCode(response.Code)
		return c.Write(response)
	}

	err = c.Read(&article)

	article.ID = int64(id)

	fmt.Print("[Handler] Update: ")
	fmt.Println(article)
	if err != nil {
		response.Code = http.StatusUnprocessableEntity
		response.Message = string(err.Error())
		response.Data = article

		c.Response.SetStatusCode(response.Code)
		return c.Write(response)
	}

	if ok, err := isRequestValid(&article); !ok {
		response.Code = http.StatusBadRequest
		response.Message = string(err.Error())
		response.Data = ok

		c.Response.SetStatusCode(response.Code)
		return c.Write(response)
	}

	ctx := c.RequestCtx
	if ctx == nil {
		ctx = c.RequestCtx
	}

	ar, err := a.AUsecase.Update(ctx, &article)

	if err != nil {
		response.Code = getStatusCode(err)
		response.Message = string(err.Error())
		response.Data = article

		c.Response.SetStatusCode(response.Code)
		return c.Write(response)
	}

	response.Code = http.StatusOK
	response.Message = models.DATA_UPDATED_SUCCESS
	response.Data = ar

	c.Response.SetStatusCode(response.Code)
	return c.Write(response)
}

func getStatusCode(err error) int {

	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case models.INTERNAL_SERVER_ERROR:
		return http.StatusInternalServerError
	case models.NOT_FOUND_ERROR:
		return http.StatusNotFound
	case models.CONFLIT_ERROR:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

func NewArticleHttpFastHttpHandler(e *routing.Router, us articleUcase.ArticleUsecase) {
	handler := &HttpArticleHandler{
		AUsecase: us,
	}

	e.Get("/article", handler.GetAll)
	e.Post("/article", handler.Create)
	e.Get("/article/:title", handler.GetByTitle)
	e.Put("/article", handler.Update) // Use Query
	//e.DELETE("/article/:id", handler.Delete)
}
