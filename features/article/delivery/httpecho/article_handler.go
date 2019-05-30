package httpecho

import (
	"context"
	"net/http"
	"strconv"

	"github.com/mochadwi/go-article/models"

	"fmt"
	"github.com/labstack/echo"
	baseHandler "github.com/mochadwi/go-article/base"
	articleUcase "github.com/mochadwi/go-article/features/article"
	"time"
)

type HttpArticleHandler struct {
	AUsecase articleUcase.ArticleUsecase
}

func (a *HttpArticleHandler) GetAll(c echo.Context) error {

	numS := c.QueryParam("num")
	num, _ := strconv.Atoi(numS)
	cursor := c.QueryParam("cursor")
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	listAr, nextCursor, err := a.AUsecase.GetAll(ctx, cursor, int64(num))

	//reqId := uuid.NewV4().String()
	var response = &models.BaseResponse{
		RequestID: "",
		Now:       time.Now().Unix(),
	}

	if err != nil {
		response.Code = baseHandler.GetStatusCode(err)
		response.Message = err.Error()
		response.Data = listAr
		return c.JSON(baseHandler.GetStatusCode(err), response)
	}

	if len(*listAr) > 0 {
		response.Message = models.DATA_AVAILABLE_SUCCESS
	} else {
		response.Message = models.DATA_EMPTY_SUCCESS
	}

	response.Code = http.StatusOK
	response.Data = listAr

	c.Response().Header().Set(`X-Cursor`, nextCursor)

	return c.JSON(response.Code, response)
}

func (a *HttpArticleHandler) GetByTitle(c echo.Context) error {

	title := c.Param("title")
	//title2 := c.QueryParam("title")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	art, err := a.AUsecase.GetByTitle(ctx, title)

	//reqId := uuid.NewV4().String()
	var response = &models.BaseResponse{
		RequestID: "",
		Now:       time.Now().Unix(),
	}

	if err != nil {
		response.Code = baseHandler.GetStatusCode(err)
		response.Message = err.Error()
		response.Data = art
		return c.JSON(response.Code, response)
	}

	response.Code = http.StatusOK
	response.Message = models.DATA_AVAILABLE_SUCCESS
	response.Data = art

	//fmt.Print("Handler: ")
	//fmt.Println(art)
	return c.JSON(response.Code, response)
}

func (a *HttpArticleHandler) GetByID(c echo.Context) error {

	idP, err := strconv.Atoi(c.Param("id"))
	id := int64(idP)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	art, err := a.AUsecase.GetByID(ctx, id)

	if err != nil {
		return c.JSON(baseHandler.GetStatusCode(err), models.BaseResponse{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, art)
}

func (a *HttpArticleHandler) Create(c echo.Context) error {
	var article models.Article
	err := c.Bind(&article)

	var response = &models.BaseResponse{
		RequestID: "",
		Now:       time.Now().Unix(),
	}

	if err != nil {
		response.Code = http.StatusUnprocessableEntity
		response.Message = string(err.Error())
		response.Data = article
		return c.JSON(response.Code, response)
	}

	if ok, err := baseHandler.IsRequestValid(&article); !ok {
		response.Code = http.StatusBadRequest
		response.Message = string(err.Error())
		response.Data = article
		return c.JSON(response.Code, response)
	}
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	ar, err := a.AUsecase.Create(ctx, &article)

	if err != nil {
		response.Code = baseHandler.GetStatusCode(err)
		response.Message = string(err.Error())
		response.Data = article
		return c.JSON(response.Code, response)
	}

	response.Code = http.StatusCreated
	response.Message = models.DATA_CREATED_SUCCESS
	response.Data = ar
	return c.JSON(response.Code, response)
}

func (a *HttpArticleHandler) Update(c echo.Context) error {

	fmt.Print("[Handler] Update id: ")
	fmt.Println(c.Param("id"))

	var article models.Article
	var id, err = strconv.Atoi(c.QueryParam("id"))

	var response = &models.BaseResponse{
		RequestID: "",
		Now:       time.Now().Unix(),
	}

	if err != nil {
		response.Code = http.StatusUnprocessableEntity
		response.Message = string(err.Error())
		response.Data = id
		return c.JSON(response.Code, response)
	}

	err = c.Bind(&article)

	article.ID = int64(id)

	fmt.Print("[Handler] Update: ")
	fmt.Println(article)
	if err != nil {
		response.Code = http.StatusUnprocessableEntity
		response.Message = string(err.Error())
		response.Data = article
		return c.JSON(response.Code, response)
	}

	if ok, err := baseHandler.IsRequestValid(&article); !ok {
		response.Code = http.StatusBadRequest
		response.Message = string(err.Error())
		response.Data = ok
		return c.JSON(response.Code, response)
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	ar, err := a.AUsecase.Update(ctx, &article)

	if err != nil {
		response.Code = baseHandler.GetStatusCode(err)
		response.Message = string(err.Error())
		response.Data = article
		return c.JSON(response.Code, response)
	}

	response.Code = http.StatusOK
	response.Message = models.DATA_UPDATED_SUCCESS
	response.Data = ar
	return c.JSON(response.Code, response)
}

//func (a *HttpArticleHandler) Delete(c echo.Context) error {
//var response = &models.BaseResponse{
//	RequestID: "",
//	Now:       time.Now().Unix(),
//}
//
//idP, err := strconv.Atoi(c.Param("id"))
//id := int64(idP)
//ctx := c.Request().Context()
//if ctx == nil {
//	ctx = context.Background()
//}
//
//status, err := a.AUsecase.Delete(ctx, id)
//
//if err != nil {
//	response.Code = baseHandler.GetStatusCode(err)
//	response.Message = string(err.Error())
//	response.Data = status
//	return c.JSON(response.Code, response)
//}
//
//response.Code = http.StatusNoContent
//response.Message = models.DATA_DELETED_SUCCESS
//response.Data = status
//return c.JSON(response.Code, response)
//}

func NewArticleHttpEchoHandler(e *echo.Echo, us articleUcase.ArticleUsecase) {
	handler := &HttpArticleHandler{
		AUsecase: us,
	}

	e.GET("/article", handler.GetAll)
	e.POST("/article", handler.Create)
	e.GET("/article/:title", handler.GetByTitle)
	e.PUT("/article", handler.Update) // Use Query
	//e.DELETE("/article/:id", handler.Delete)
}
