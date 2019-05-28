package httpecho

import (
	"context"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/mochadwi/go-article/models"

	"github.com/labstack/echo"
	ratingUcase "github.com/mochadwi/go-article/rating"

	"bytes"
	"fmt"
	"github.com/mochadwi/go-article/rating/template/gofiles"
	"gopkg.in/go-playground/validator.v9"
	"time"
)

type HttpRatingHandler struct {
	AUsecase ratingUcase.RatingUsecase
}

func (a *HttpRatingHandler) GetAll(c echo.Context) error {

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
		response.Code = getStatusCode(err)
		response.Message = err.Error()
		response.Data = listAr
		return c.JSON(getStatusCode(err), response)
	}

	if len(*listAr) > 0 {
		response.Message = models.DATA_AVAILABLE_SUCCESS
	} else {
		response.Message = models.DATA_EMPTY_SUCCESS
	}

	response.Code = http.StatusOK
	response.Data = listAr

	c.Response().Header().Set(`X-Cursor`, nextCursor)

	buffer := new(bytes.Buffer)

	var ratings []string
	for _, rating := range *listAr {
		ratings = append(ratings, rating.Title)
	}

	gofiles.RatingList(ratings, buffer)
	return c.HTMLBlob(response.Code, buffer.Bytes())

	//return c.JSON(response.Code, response)
}

func (a *HttpRatingHandler) GetByTitle(c echo.Context) error {

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
		response.Code = getStatusCode(err)
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

func (a *HttpRatingHandler) GetByID(c echo.Context) error {

	idP, err := strconv.Atoi(c.Param("id"))
	id := int64(idP)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	art, err := a.AUsecase.GetByID(ctx, id)

	if err != nil {
		return c.JSON(getStatusCode(err), models.BaseResponse{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, art)
}

func isRequestValid(m *models.Rating) (bool, error) {

	validate := validator.New()

	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (a *HttpRatingHandler) Create(c echo.Context) error {
	var rating models.Rating
	err := c.Bind(&rating)

	var response = &models.BaseResponse{
		RequestID: "",
		Now:       time.Now().Unix(),
	}

	if err != nil {
		response.Code = http.StatusUnprocessableEntity
		response.Message = string(err.Error())
		response.Data = rating
		return c.JSON(response.Code, response)
	}

	if ok, err := isRequestValid(&rating); !ok {
		response.Code = http.StatusBadRequest
		response.Message = string(err.Error())
		response.Data = rating
		return c.JSON(response.Code, response)
	}
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	ar, err := a.AUsecase.Create(ctx, &rating)

	if err != nil {
		response.Code = getStatusCode(err)
		response.Message = string(err.Error())
		response.Data = rating
		return c.JSON(response.Code, response)
	}

	response.Code = http.StatusCreated
	response.Message = models.DATA_CREATED_SUCCESS
	response.Data = ar
	return c.JSON(response.Code, response)
}

func (a *HttpRatingHandler) Update(c echo.Context) error {

	fmt.Print("[Handler] Update id: ")
	fmt.Println(c.Param("id"))

	var rating models.Rating
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

	err = c.Bind(&rating)

	rating.ID = int64(id)

	fmt.Print("[Handler] Update: ")
	fmt.Println(rating)
	if err != nil {
		response.Code = http.StatusUnprocessableEntity
		response.Message = string(err.Error())
		response.Data = rating
		return c.JSON(response.Code, response)
	}

	if ok, err := isRequestValid(&rating); !ok {
		response.Code = http.StatusBadRequest
		response.Message = string(err.Error())
		response.Data = ok
		return c.JSON(response.Code, response)
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	ar, err := a.AUsecase.Update(ctx, &rating)

	if err != nil {
		response.Code = getStatusCode(err)
		response.Message = string(err.Error())
		response.Data = rating
		return c.JSON(response.Code, response)
	}

	response.Code = http.StatusOK
	response.Message = models.DATA_UPDATED_SUCCESS
	response.Data = ar
	return c.JSON(response.Code, response)
}

//func (a *HttpRatingHandler) Delete(c echo.Context) error {
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
//	response.Code = getStatusCode(err)
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

func NewRatingHttpEchoHandler(e *echo.Echo, us ratingUcase.RatingUsecase) {
	handler := &HttpRatingHandler{
		AUsecase: us,
	}

	e.GET("/rating", handler.GetAll)
	e.POST("/rating", handler.Create)
	e.GET("/rating/:title", handler.GetByTitle)
	e.PUT("/rating", handler.Update) // Use Query
	//e.DELETE("/rating/:id", handler.Delete)
}
